package core

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/librespot-org/librespot-golang/Spotify"
	"github.com/librespot-org/librespot-golang/librespot/connection"
	"github.com/librespot-org/librespot-golang/librespot/discovery"
	"github.com/librespot-org/librespot-golang/librespot/utils"
	"log"
)

var Version = "master"
var BuildID = "dev"

// Login to Spotify using username and password
func Login(username string, password string, deviceName string) (*Session, error) {
	s, err := setupSession()
	if err != nil {
		return s, err
	}

	return s, s.loginSession(username, password, deviceName)
}

func (s *Session) loginSession(username string, password string, deviceName string) error {
	s.deviceId = utils.GenerateDeviceId(deviceName)
	s.deviceName = deviceName

	err := s.startConnection()
	if err != nil {
		return err
	}
	loginPacket := makeLoginPasswordPacket(username, password, s.deviceId)
	return s.doLogin(loginPacket, username)
}

// Login to Spotify using an existing authData blob, as previously saved via Session.AuthDataBlob(). If the blob
// carries an OAuth refresh token and clientId/clientSecret are non-empty, it's redeemed for a fresh Web API access
// token (see Session.AccessToken), so features like Search work without requiring a fresh interactive login.
func LoginSaved(username string, authData []byte, deviceName string, clientId string, clientSecret string) (*Session, error) {
	ad := unmarshalAuthData(authData)
	if username == "" {
		username = ad.Username
	}
	if username == "" {
		return nil, fmt.Errorf("no username supplied and none stored in the auth data")
	}

	s, err := setupSession()
	if err != nil {
		return s, err
	}
	s.deviceId = utils.GenerateDeviceId(deviceName)
	s.deviceName = deviceName

	err = s.startConnection()
	if err != nil {
		return s, err
	}

	packet := makeLoginBlobPacket(username, ad.Blob,
		Spotify.AuthenticationType_AUTHENTICATION_STORED_SPOTIFY_CREDENTIALS.Enum(), s.deviceId)
	if err := s.doLogin(packet, username); err != nil {
		return s, err
	}

	pkce := ad.ClientID == PKCEClientId
	s.refreshToken = ad.RefreshToken
	s.clientId, s.clientSecret = ad.ClientID, ad.ClientSecret
	s.searchClientId, s.searchClientSecret = ad.SearchClientID, ad.SearchClientSecret

	if clientId != "" && clientSecret != "" {
		if pkce {
			// With a PKCE login the refresh token belongs to the shared official client id, whose Web API quota is
			// usually exhausted (search fails with 429): explicit credentials are the application used for search.
			s.searchClientId, s.searchClientSecret = clientId, clientSecret
		} else {
			// explicit credentials override the ones stored in the blob file
			s.clientId, s.clientSecret = clientId, clientSecret
		}
	}

	if pkce && s.searchClientId != "" && s.searchClientSecret != "" {
		// app token (client credentials flow) for catalog calls such as Search
		token, err := clientCredentialsToken(s.searchClientId, s.searchClientSecret)
		if err != nil && (s.searchClientId != ad.SearchClientID || s.searchClientSecret != ad.SearchClientSecret) {
			// the credentials just supplied are invalid: don't replace the stored ones
			log.Println("WARNING: the client id/secret supplied are not valid (", err, ") - ignored")
			s.searchClientId, s.searchClientSecret = ad.SearchClientID, ad.SearchClientSecret
			if s.searchClientId != "" && s.searchClientSecret != "" {
				token, err = clientCredentialsToken(s.searchClientId, s.searchClientSecret)
			} else {
				token, err = nil, nil
			}
		}

		if err != nil {
			log.Println("WARNING: the saved client id/secret for search are not valid (", err, ") - search will use the "+
				"shared client id and is likely to fail with 429. Run gospy --login --client-id ID --client-secret SECRET to fix it.")
		} else if token != nil {
			s.accessToken = token.AccessToken
			s.searchAppToken = true
		}
	}

	if s.accessToken == "" && ad.RefreshToken != "" && s.clientId != "" && (s.clientSecret != "" || pkce) {
		// a PKCE login (official client id) doesn't need a client secret
		token, err := RefreshOAuthToken(ad.RefreshToken, s.clientId, s.clientSecret)
		if err != nil {
			log.Println("could not refresh OAuth token, Web API calls (e.g. Search) will be unavailable:", err)
		} else {
			s.accessToken = token.AccessToken
			if token.RefreshToken != "" {
				s.refreshToken = token.RefreshToken
			}
		}
	}

	return s, nil
}

// Registers librespot as a Spotify Connect device via mdns. When user connects, logs on to Spotify and saves
// credentials in file at cacheBlobPath. Once saved, the blob credentials allow the program to connect to other
// Spotify Connect devices and control them.
func LoginDiscovery(cacheBlobPath string, deviceName string) (*Session, error) {
	deviceId := utils.GenerateDeviceId(deviceName)
	disc := discovery.LoginFromConnect(cacheBlobPath, deviceId, deviceName)
	return sessionFromDiscovery(disc)
}

// Login using an authentication blob through Spotify Connect discovery system, reading an existing blob data. To read
// from a file, see LoginDiscoveryBlobFile.
func LoginDiscoveryBlob(username string, blob string, deviceName string) (*Session, error) {
	deviceId := utils.GenerateDeviceId(deviceName)
	disc := discovery.CreateFromBlob(utils.BlobInfo{
		Username:    username,
		DecodedBlob: blob,
	}, "", deviceId, deviceName)
	return sessionFromDiscovery(disc)
}

// Login from credentials at cacheBlobPath previously saved by LoginDiscovery. Similar to LoginDiscoveryBlob, except
// it reads it directly from a file.
func LoginDiscoveryBlobFile(cacheBlobPath, deviceName string) (*Session, error) {
	deviceId := utils.GenerateDeviceId(deviceName)
	disc := discovery.CreateFromFile(cacheBlobPath, deviceId, deviceName)
	return sessionFromDiscovery(disc)
}

// Login to Spotify using the OAuth method
func LoginOAuth(deviceName string, clientId string, clientSecret string) (*Session, error) {
	token := getOAuthToken(clientId, clientSecret)
	s, err := loginOAuthToken(token.AccessToken, deviceName)
	if err != nil {
		return s, err
	}

	s.accessToken = token.AccessToken
	s.refreshToken = token.RefreshToken
	s.clientId, s.clientSecret = clientId, clientSecret
	return s, nil
}

// LoginOAuthPKCE logs in through the browser using the official Spotify client id (OAuth authorization code flow
// with PKCE, no client secret needed). The credentials obtained this way can be used with login5 (see
// Session.SpclientToken). Save Session.AuthDataBlob() and use LoginSaved to log in again without the browser.
//
// searchClientId/searchClientSecret are optional: the credentials of your own Spotify application, used (and saved
// in AuthDataBlob) to search through the Web API, since the shared official client id is heavily rate limited.
func LoginOAuthPKCE(deviceName string, searchClientId string, searchClientSecret string) (*Session, error) {
	// check the search credentials first, so a typo doesn't go unnoticed (or cost a browser login)
	var searchToken *OAuth
	if searchClientId != "" && searchClientSecret != "" {
		var err error
		searchToken, err = clientCredentialsToken(searchClientId, searchClientSecret)
		if err != nil {
			return nil, fmt.Errorf("invalid client id/secret (%v)", err)
		}
	}

	token, err := getPKCEOAuthToken()
	if err != nil {
		return nil, err
	}

	s, err := loginOAuthToken(token.AccessToken, deviceName)
	if err != nil {
		return s, err
	}

	s.accessToken = token.AccessToken
	s.refreshToken = token.RefreshToken
	s.clientId = PKCEClientId

	if searchToken != nil {
		s.searchClientId, s.searchClientSecret = searchClientId, searchClientSecret
		s.accessToken = searchToken.AccessToken
		s.searchAppToken = true
	}

	return s, nil
}

func loginOAuthToken(accessToken string, deviceName string) (*Session, error) {
	s, err := setupSession()
	if err != nil {
		return s, err
	}

	s.deviceId = utils.GenerateDeviceId(deviceName)
	s.deviceName = deviceName

	err = s.startConnection()
	if err != nil {
		return s, err
	}

	packet := makeLoginBlobPacket("", []byte(accessToken),
		Spotify.AuthenticationType_AUTHENTICATION_SPOTIFY_TOKEN.Enum(), s.deviceId)
	return s, s.doLogin(packet, "")
}

func (s *Session) doLogin(packet []byte, username string) error {
	err := s.stream.SendPacket(connection.PacketLogin, packet)
	if err != nil {
		log.Fatal("bad shannon write", err)
	}

	// Pll once for authentication response
	welcome, err := s.handleLogin()
	if err != nil {
		return err
	}

	// Store the few interesting values
	s.username = welcome.GetCanonicalUsername()
	if s.username == "" {
		// Spotify might not return a canonical username, so reuse the blob's one instead
		s.username = s.discovery.LoginBlob().Username
	}
	s.reusableAuthBlob = welcome.GetReusableAuthCredentials()

	// Poll for acknowledge before loading - needed for gopherjs
	// s.poll()
	go s.runPollLoop()

	return nil
}

func (s *Session) handleLogin() (*Spotify.APWelcome, error) {
	cmd, data, err := s.stream.RecvPacket()
	if err != nil {
		return nil, fmt.Errorf("authentication failed: %v", err)
	}

	if cmd == connection.PacketAuthFailure {
		return nil, fmt.Errorf("authentication failed")
	} else if cmd == connection.PacketAPWelcome {
		welcome := &Spotify.APWelcome{}
		err := proto.Unmarshal(data, welcome)
		if err != nil {
			return nil, fmt.Errorf("authentication failed: %v", err)
		}
		fmt.Println("Authentication succeeded: Welcome,", welcome.GetCanonicalUsername())
		fmt.Println("Blob type:", welcome.GetReusableAuthCredentialsType())
		return welcome, nil
	} else {
		return nil, fmt.Errorf("authentication failed: unexpected cmd %v", cmd)
	}
}

func (s *Session) getLoginBlobPacket(blob utils.BlobInfo) []byte {
	data, _ := base64.StdEncoding.DecodeString(blob.DecodedBlob)

	buffer := bytes.NewBuffer(data)
	buffer.ReadByte()
	readBytes(buffer)
	buffer.ReadByte()
	authNum := readInt(buffer)
	authType := Spotify.AuthenticationType(authNum)
	buffer.ReadByte()
	authData := readBytes(buffer)

	return makeLoginBlobPacket(blob.Username, authData, &authType, s.deviceId)
}

func makeLoginPasswordPacket(username string, password string, deviceId string) []byte {
	return makeLoginBlobPacket(username, []byte(password),
		Spotify.AuthenticationType_AUTHENTICATION_USER_PASS.Enum(), deviceId)
}

func makeLoginBlobPacket(username string, authData []byte,
	authType *Spotify.AuthenticationType, deviceId string) []byte {
	versionString := "librespot-golang_" + Version + "_" + BuildID

	packet := &Spotify.ClientResponseEncrypted{
		LoginCredentials: &Spotify.LoginCredentials{
			Username: proto.String(username),
			Typ:      authType,
			AuthData: authData,
		},
		SystemInfo: &Spotify.SystemInfo{
			CpuFamily:               Spotify.CpuFamily_CPU_UNKNOWN.Enum(),
			Os:                      Spotify.Os_OS_UNKNOWN.Enum(),
			SystemInformationString: proto.String("librespot-golang"),
			DeviceId:                proto.String(deviceId),
		},
		VersionString: proto.String(versionString),
	}

	packetData, err := proto.Marshal(packet)
	if err != nil {
		log.Fatal("login marshaling error: ", err)
	}
	return packetData
}
