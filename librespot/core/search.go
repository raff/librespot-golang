package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/librespot-org/librespot-golang/librespot/metadata"
)

// webApiImage is a single entry of the "images" arrays returned by Spotify's Web API.
type webApiImage struct {
	Url string `json:"url"`
}

func firstImageUrl(images []webApiImage) string {
	if len(images) == 0 {
		return ""
	}
	return images[0].Url
}

type webApiArtist struct {
	Name   string        `json:"name"`
	Uri    string        `json:"uri"`
	Images []webApiImage `json:"images"`
}

func (a webApiArtist) toMetadata() metadata.Artist {
	return metadata.Artist{Name: a.Name, Uri: a.Uri, Image: firstImageUrl(a.Images)}
}

type webApiAlbum struct {
	Name    string         `json:"name"`
	Uri     string         `json:"uri"`
	Images  []webApiImage  `json:"images"`
	Artists []webApiArtist `json:"artists"`
}

func (a webApiAlbum) toMetadata() metadata.Album {
	artists := make([]metadata.Artist, len(a.Artists))
	for i, ar := range a.Artists {
		artists[i] = ar.toMetadata()
	}
	return metadata.Album{Name: a.Name, Uri: a.Uri, Image: firstImageUrl(a.Images), Artists: artists}
}

type webApiTrack struct {
	Name       string         `json:"name"`
	Uri        string         `json:"uri"`
	DurationMs int            `json:"duration_ms"`
	Popularity float32        `json:"popularity"`
	Album      webApiAlbum    `json:"album"`
	Artists    []webApiArtist `json:"artists"`
}

func (t webApiTrack) toMetadata() metadata.Track {
	artists := make([]metadata.Artist, len(t.Artists))
	for i, ar := range t.Artists {
		artists[i] = ar.toMetadata()
	}
	return metadata.Track{
		Name:       t.Name,
		Uri:        t.Uri,
		Duration:   t.DurationMs,
		Popularity: t.Popularity,
		Album:      t.Album.toMetadata(),
		Artists:    artists,
		Image:      firstImageUrl(t.Album.Images),
	}
}

type webApiPlaylist struct {
	Name   string        `json:"name"`
	Uri    string        `json:"uri"`
	Images []webApiImage `json:"images"`
	Owner  struct {
		DisplayName string `json:"display_name"`
	} `json:"owner"`
}

func (p webApiPlaylist) toMetadata() metadata.Playlist {
	return metadata.Playlist{Name: p.Name, Uri: p.Uri, Image: firstImageUrl(p.Images), Author: p.Owner.DisplayName}
}

// webApiSearchResponse mirrors the relevant subset of the response shape of
// https://api.spotify.com/v1/search
type webApiSearchResponse struct {
	Tracks struct {
		Items []webApiTrack `json:"items"`
		Total int           `json:"total"`
	} `json:"tracks"`
	Albums struct {
		Items []webApiAlbum `json:"items"`
		Total int           `json:"total"`
	} `json:"albums"`
	Artists struct {
		Items []webApiArtist `json:"items"`
		Total int            `json:"total"`
	} `json:"artists"`
	Playlists struct {
		Items []webApiPlaylist `json:"items"`
		Total int              `json:"total"`
	} `json:"playlists"`
	Error struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"error"`
}

// Search performs a catalog search using Spotify's official Web API (the old Mercury-based
// "hm://searchview/km/..." endpoint this library used to use has been retired by Spotify). It
// requires the session to have a Web API access token, which is only available when the session
// was logged in via OAuth (LoginOAuth), or via LoginSaved with a blob carrying a refresh token
// and non-empty clientId/clientSecret.
func (s *Session) Search(query string, limit int) (*metadata.SearchResponse, error) {
	if s.accessToken == "" {
		return nil, fmt.Errorf("no Web API access token available: login via OAuth (or LoginSaved with " +
			"clientId/clientSecret and a saved refresh token) to use Search")
	}

	v := url.Values{}
	v.Set("q", query)
	v.Set("type", "track,album,artist,playlist")
	v.Set("limit", fmt.Sprintf("%d", limit))

	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/search?"+v.Encode(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+s.accessToken)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	webResp := webApiSearchResponse{}
	if err := json.Unmarshal(body, &webResp); err != nil {
		return nil, fmt.Errorf("could not parse search response (status %d): %v", resp.StatusCode, err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search failed: status %d: %s", resp.StatusCode, webResp.Error.Message)
	}

	result := &metadata.SearchResponse{}
	result.Results.Tracks.Total = webResp.Tracks.Total
	result.Results.Albums.Total = webResp.Albums.Total
	result.Results.Artists.Total = webResp.Artists.Total
	result.Results.Playlists.Total = webResp.Playlists.Total

	for _, t := range webResp.Tracks.Items {
		result.Results.Tracks.Hits = append(result.Results.Tracks.Hits, t.toMetadata())
	}
	for _, a := range webResp.Albums.Items {
		result.Results.Albums.Hits = append(result.Results.Albums.Hits, a.toMetadata())
	}
	for _, a := range webResp.Artists.Items {
		result.Results.Artists.Hits = append(result.Results.Artists.Hits, a.toMetadata())
	}
	for _, p := range webResp.Playlists.Items {
		result.Results.Playlists.Hits = append(result.Results.Playlists.Hits, p.toMetadata())
	}

	return result, nil
}
