package core

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"google.golang.org/protobuf/encoding/protowire"
)

const (
	login5Url = "https://login5.spotify.com/v3/login"

	// client id of the official Spotify desktop client, which is the one spclient accepts tokens for
	login5ClientId = "65b708073fc0480ea92a077233ca87bd"
)

var login5Errors = map[int32]string{
	0: "UNKNOWN_ERROR", 1: "INVALID_CREDENTIALS", 2: "BAD_REQUEST", 3: "UNSUPPORTED_LOGIN_PROTOCOL",
	4: "TIMEOUT", 5: "UNKNOWN_IDENTIFIER", 6: "TOO_MANY_ATTEMPTS", 7: "INVALID_PHONENUMBER", 8: "TRY_AGAIN_LATER",
}

type login5Token struct {
	token   string
	expires time.Time
}

var login5Cache struct {
	sync.Mutex
	tok login5Token
}

// buildLogin5Request encodes a login5 LoginRequest using a stored credential:
//
//	LoginRequest { client_info(1) { client_id(1), device_id(2) }, stored_credential(100) { username(1), data(2) } }
func buildLogin5Request(clientId, deviceId, username string, data []byte) []byte {
	var ci []byte
	ci = protowire.AppendTag(ci, 1, protowire.BytesType)
	ci = protowire.AppendString(ci, clientId)
	ci = protowire.AppendTag(ci, 2, protowire.BytesType)
	ci = protowire.AppendString(ci, deviceId)

	var sc []byte
	sc = protowire.AppendTag(sc, 1, protowire.BytesType)
	sc = protowire.AppendString(sc, username)
	sc = protowire.AppendTag(sc, 2, protowire.BytesType)
	sc = protowire.AppendBytes(sc, data)

	var req []byte
	req = protowire.AppendTag(req, 1, protowire.BytesType)
	req = protowire.AppendBytes(req, ci)
	req = protowire.AppendTag(req, 100, protowire.BytesType)
	req = protowire.AppendBytes(req, sc)
	return req
}

// parseLogin5Response decodes a login5 LoginResponse:
//
//	LoginResponse { ok(1) { username(1), access_token(2), stored_credential(3), access_token_expires_in(4) },
//	                error(2), challenges(3) }
func parseLogin5Response(b []byte) (token string, expiresIn int, err error) {
	for len(b) > 0 {
		num, typ, n := protowire.ConsumeTag(b)
		if n < 0 {
			return "", 0, protowire.ParseError(n)
		}
		b = b[n:]

		switch {
		case num == 1 && typ == protowire.BytesType:
			ok, n := protowire.ConsumeBytes(b)
			if n < 0 {
				return "", 0, protowire.ParseError(n)
			}
			b = b[n:]

			for len(ok) > 0 {
				fnum, ftyp, n := protowire.ConsumeTag(ok)
				if n < 0 {
					return "", 0, protowire.ParseError(n)
				}
				ok = ok[n:]

				switch {
				case fnum == 2 && ftyp == protowire.BytesType:
					s, n := protowire.ConsumeString(ok)
					if n < 0 {
						return "", 0, protowire.ParseError(n)
					}
					token, ok = s, ok[n:]
				case fnum == 4 && ftyp == protowire.VarintType:
					v, n := protowire.ConsumeVarint(ok)
					if n < 0 {
						return "", 0, protowire.ParseError(n)
					}
					expiresIn, ok = int(v), ok[n:]
				default:
					n := protowire.ConsumeFieldValue(fnum, ftyp, ok)
					if n < 0 {
						return "", 0, protowire.ParseError(n)
					}
					ok = ok[n:]
				}
			}

			if token != "" {
				return token, expiresIn, nil
			}

		case num == 3 && typ == protowire.BytesType:
			return "", 0, fmt.Errorf("login5 requires a challenge, which is not supported")

		case num == 2 && typ == protowire.VarintType:
			v, _ := protowire.ConsumeVarint(b)
			name, found := login5Errors[int32(v)]
			if !found {
				name = fmt.Sprint(v)
			}
			return "", 0, fmt.Errorf("login5 error: %s", name)

		default:
			n := protowire.ConsumeFieldValue(num, typ, b)
			if n < 0 {
				return "", 0, protowire.ParseError(n)
			}
			b = b[n:]
		}
	}

	return "", 0, fmt.Errorf("login5 response has no access token")
}

// SpclientToken returns an access token valid for Spotify's internal "spclient" HTTP endpoints (metadata,
// extended-metadata, storage-resolve...). It is obtained through the login5 service using the session's stored
// credentials (the replacement for the retired hm://keymaster/token endpoint), and cached until it expires.
func (s *Session) SpclientToken() (string, error) {
	login5Cache.Lock()
	defer login5Cache.Unlock()

	if login5Cache.tok.token != "" && time.Now().Before(login5Cache.tok.expires) {
		return login5Cache.tok.token, nil
	}

	if len(s.reusableAuthBlob) == 0 {
		return "", fmt.Errorf("no stored credentials available for login5")
	}

	body := buildLogin5Request(login5ClientId, s.deviceId, s.username, s.reusableAuthBlob)
	req, err := http.NewRequest("POST", login5Url, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-protobuf")
	req.Header.Set("Accept", "application/x-protobuf")

	clientToken, err := getClientToken(login5ClientId, s.deviceId)
	if err != nil {
		return "", err
	}
	req.Header.Set("Client-Token", clientToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("login5 failed: status %d: %s", resp.StatusCode, data)
	}

	token, expiresIn, err := parseLogin5Response(data)
	if err != nil {
		return "", err
	}

	if expiresIn <= 60 {
		expiresIn = 60
	}
	login5Cache.tok = login5Token{token, time.Now().Add(time.Duration(expiresIn-30) * time.Second)}
	return token, nil
}

const clientTokenUrl = "https://clienttoken.spotify.com/v1/clienttoken"

// getClientToken requests a "client token" from Spotify, which is required by login5 (and spclient) to
// identify the client application:
//
//	ClientTokenRequest { request_type(1)=1, client_data(2) { client_version(1), client_id(2),
//	    connectivity_sdk_data(3) { platform_specific_data(1) { desktop(4) { ... } }, device_id(2) } } }
//	ClientTokenResponse { response_type(1), granted_token(2) { token(1) } }
func getClientToken(clientId, deviceId string) (string, error) {
	str := func(num protowire.Number, v string) []byte {
		return protowire.AppendString(protowire.AppendTag(nil, num, protowire.BytesType), v)
	}
	msg := func(num protowire.Number, v []byte) []byte {
		return protowire.AppendBytes(protowire.AppendTag(nil, num, protowire.BytesType), v)
	}

	desktop := append(str(1, "0"), append(str(2, "0"), str(3, "0")...)...)
	sdk := append(msg(1, msg(4, desktop)), str(2, deviceId)...)
	clientData := append(str(1, "1.2.52.442.g3c4d1d3b"), append(str(2, clientId), msg(3, sdk)...)...)

	body := protowire.AppendVarint(protowire.AppendTag(nil, 1, protowire.VarintType), 1)
	body = append(body, msg(2, clientData)...)

	req, err := http.NewRequest("POST", clientTokenUrl, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-protobuf")
	req.Header.Set("Accept", "application/x-protobuf")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("client token request failed: status %d", resp.StatusCode)
	}

	for len(data) > 0 {
		num, typ, n := protowire.ConsumeTag(data)
		if n < 0 {
			return "", protowire.ParseError(n)
		}
		data = data[n:]

		if num == 2 && typ == protowire.BytesType {
			granted, n := protowire.ConsumeBytes(data)
			if n < 0 {
				return "", protowire.ParseError(n)
			}
			if fnum, ftyp, n := protowire.ConsumeTag(granted); n > 0 && fnum == 1 && ftyp == protowire.BytesType {
				tok, m := protowire.ConsumeString(granted[n:])
				if m < 0 {
					return "", protowire.ParseError(m)
				}
				return tok, nil
			}
			return "", fmt.Errorf("client token response has no token")
		}

		n = protowire.ConsumeFieldValue(num, typ, data)
		if n < 0 {
			return "", protowire.ParseError(n)
		}
		data = data[n:]
	}

	return "", fmt.Errorf("client token response has no granted token")
}
