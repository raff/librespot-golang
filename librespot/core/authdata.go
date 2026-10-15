package core

import "encoding/json"

// AuthData is the persisted, reusable authentication state for a session: the
// low-level Spotify Connect reconnection blob (used to reauthenticate the
// Mercury/audio connection), plus, when available, an OAuth refresh token
// that can be redeemed for a Web API access token (needed for e.g. search,
// since Spotify has retired the old Mercury search endpoints).
type AuthData struct {
	Username     string `json:"username,omitempty"`
	Blob         []byte `json:"blob"`
	RefreshToken string `json:"refresh_token,omitempty"`

	// OAuth application credentials needed to redeem RefreshToken, so that
	// callers don't have to supply them again on every launch.
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`

	// Credentials of the application used for Web API search (client credentials flow), when the
	// login was done with the official client id (see LoginOAuthPKCE).
	SearchClientID     string `json:"search_client_id,omitempty"`
	SearchClientSecret string `json:"search_client_secret,omitempty"`
}

func marshalAuthData(ad AuthData) []byte {
	data, _ := json.Marshal(ad)
	return data
}

// unmarshalAuthData parses the combined blob+refresh-token file format. For
// backward compatibility with files written before this format existed, if
// the data isn't valid JSON it's treated as a raw legacy blob with no
// refresh token.
func unmarshalAuthData(data []byte) AuthData {
	var ad AuthData
	if err := json.Unmarshal(data, &ad); err != nil {
		return AuthData{Blob: data}
	}
	return ad
}

// SearchCredentials returns the search client id/secret stored in an auth data blob (see Session.AuthDataBlob),
// or empty strings if there are none.
func SearchCredentials(authData []byte) (clientId string, clientSecret string) {
	ad := unmarshalAuthData(authData)
	return ad.SearchClientID, ad.SearchClientSecret
}
