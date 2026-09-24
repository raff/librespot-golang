package core

import "encoding/json"

// AuthData is the persisted, reusable authentication state for a session: the
// low-level Spotify Connect reconnection blob (used to reauthenticate the
// Mercury/audio connection), plus, when available, an OAuth refresh token
// that can be redeemed for a Web API access token (needed for e.g. search,
// since Spotify has retired the old Mercury search endpoints).
type AuthData struct {
	Blob         []byte `json:"blob"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func marshalAuthData(blob []byte, refreshToken string) []byte {
	data, _ := json.Marshal(AuthData{Blob: blob, RefreshToken: refreshToken})
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
