package core

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/toqueteos/webbrowser"
)

type OAuth struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	Error        string
}

func GetOauthAccessToken(code string, redirectUri string, clientId string, clientSecret string) (*OAuth, error) {
	val := url.Values{}
	val.Set("grant_type", "authorization_code")
	val.Set("code", code)
	val.Set("redirect_uri", redirectUri)
	val.Set("client_id", clientId)
	val.Set("client_secret", clientSecret)

	return postOauthToken(val)
}

// RefreshOAuthToken redeems a previously obtained refresh token for a fresh
// access token, without requiring an interactive browser login. Spotify may
// or may not return a new refresh token in the response; if it doesn't, the
// caller should keep using the refresh token it already has.
func RefreshOAuthToken(refreshToken string, clientId string, clientSecret string) (*OAuth, error) {
	val := url.Values{}
	val.Set("grant_type", "refresh_token")
	val.Set("refresh_token", refreshToken)
	val.Set("client_id", clientId)
	if clientSecret != "" {
		val.Set("client_secret", clientSecret)
	}

	return postOauthToken(val)
}

// clientCredentialsToken gets an application (no user) token, enough for catalog calls such as search.
func clientCredentialsToken(clientId string, clientSecret string) (*OAuth, error) {
	val := url.Values{}
	val.Set("grant_type", "client_credentials")
	val.Set("client_id", clientId)
	val.Set("client_secret", clientSecret)

	return postOauthToken(val)
}

func postOauthToken(val url.Values) (*OAuth, error) {
	resp, err := http.PostForm("https://accounts.spotify.com/api/token", val)
	if err != nil {
		// Retry since there is an nginx bug that causes http2 streams to get
		// an initial REFUSED_STREAM response
		// https://github.com/curl/curl/issues/804
		resp, err = http.PostForm("https://accounts.spotify.com/api/token", val)
		if err != nil {
			return nil, err
		}
	}
	defer resp.Body.Close()
	auth := OAuth{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &auth)
	if err != nil {
		return nil, err
	}
	if auth.Error != "" {
		return nil, fmt.Errorf("error getting token %v", auth.Error)
	}
	return &auth, nil
}

func getOAuthToken(clientId string, clientSecret string) OAuth {
	ch := make(chan OAuth)

	urlPath := "https://accounts.spotify.com/authorize?" +
		"client_id=" + clientId +
		"&response_type=code" +
		"&redirect_uri=http://127.0.0.1:8888/callback" +
		"&scope=streaming"

		//fmt.Println("go to this url")
		//fmt.Println(urlPath)
	webbrowser.Open(urlPath)

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		auth, err := GetOauthAccessToken(params.Get("code"), "http://127.0.0.1:8888/callback", clientId, clientSecret)
		if err != nil {
			fmt.Fprintf(w, "Error getting token %q", err)
			return
		}
		fmt.Fprintf(w, "Got token, loggin in")
		ch <- *auth
	})

	go func() {
		log.Fatal(http.ListenAndServe(":8888", nil))
	}()

	return <-ch
}

const (
	// PKCEClientId is the client id of the official Spotify desktop client. Credentials obtained through it are
	// the ones accepted by login5/spclient (the Web API client id of a third party application is not).
	PKCEClientId    = "65b708073fc0480ea92a077233ca87bd"
	pkceRedirectUri = "http://127.0.0.1:8898/login"
	pkceScopes      = "app-remote-control,playlist-modify,playlist-modify-private,playlist-modify-public," +
		"playlist-read,playlist-read-collaborative,playlist-read-private,streaming,ugc-image-upload," +
		"user-follow-modify,user-follow-read,user-library-modify,user-library-read,user-modify," +
		"user-modify-playback-state,user-modify-private,user-personalized,user-read-birthdate," +
		"user-read-currently-playing,user-read-email,user-read-play-history,user-read-playback-position," +
		"user-read-playback-state,user-read-private,user-read-recently-played,user-top-read"
)

// getPKCEOAuthToken runs the OAuth authorization code flow with PKCE (no client secret) for the official
// Spotify client id, opening the browser and waiting for the redirect on 127.0.0.1:8898.
func getPKCEOAuthToken() (*OAuth, error) {
	rnd := make([]byte, 64)
	if _, err := rand.Read(rnd); err != nil {
		return nil, err
	}
	verifier := base64.RawURLEncoding.EncodeToString(rnd)
	sum := sha256.Sum256([]byte(verifier))
	challenge := base64.RawURLEncoding.EncodeToString(sum[:])

	state := make([]byte, 8)
	rand.Read(state)
	stateStr := base64.RawURLEncoding.EncodeToString(state)

	authUrl := "https://accounts.spotify.com/authorize?" + url.Values{
		"client_id":             {PKCEClientId},
		"response_type":         {"code"},
		"redirect_uri":          {pkceRedirectUri},
		"scope":                 {strings.ReplaceAll(pkceScopes, ",", " ")},
		"code_challenge":        {challenge},
		"code_challenge_method": {"S256"},
		"state":                 {stateStr},
	}.Encode()

	type result struct {
		auth *OAuth
		err  error
	}
	ch := make(chan result, 1)

	mux := http.NewServeMux()
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if q.Get("state") != stateStr || q.Get("code") == "" {
			fmt.Fprintf(w, "Login failed: %s", q.Get("error"))
			ch <- result{nil, fmt.Errorf("authorization failed: %q", q.Get("error"))}
			return
		}

		val := url.Values{}
		val.Set("grant_type", "authorization_code")
		val.Set("code", q.Get("code"))
		val.Set("redirect_uri", pkceRedirectUri)
		val.Set("client_id", PKCEClientId)
		val.Set("code_verifier", verifier)

		auth, err := postOauthToken(val)
		if err != nil {
			fmt.Fprintf(w, "Error getting token: %v", err)
		} else {
			fmt.Fprintf(w, "Got token, you can close this window")
		}
		ch <- result{auth, err}
	})

	ln, err := net.Listen("tcp", "127.0.0.1:8898")
	if err != nil {
		return nil, err
	}
	srv := &http.Server{Handler: mux}
	go srv.Serve(ln)
	defer srv.Close()

	fmt.Println("Open this URL in your browser to log in:")
	fmt.Println(authUrl)
	webbrowser.Open(authUrl)

	select {
	case r := <-ch:
		return r.auth, r.err
	case <-time.After(5 * time.Minute):
		return nil, fmt.Errorf("timed out waiting for browser login")
	}
}
