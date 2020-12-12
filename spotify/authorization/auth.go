package authorization

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Oppodelldog/spotify-sleep-timer/config"
)

const endpointURI = "https://accounts.spotify.com/"

func AuthURL() (string, error) {
	u, err := url.Parse(getPath("authorize"))
	if err != nil {
		return "", fmt.Errorf("error parsing spotify auth URL: %w", err)
	}

	q := u.Query()
	q.Add("client_id", config.SpotifyClientID)
	q.Add("response_type", "code")
	q.Add("redirect_uri", config.SpotifyAuthRedirectURI)
	q.Add("scope", "user-modify-playback-state  user-read-playback-state user-read-email")
	q.Add("state", "10292383492")
	u.RawQuery = q.Encode()

	return u.String(), nil
}

func Auth(code string) (AuthResponse, error) {
	var (
		authResponse AuthResponse
		data         = url.Values{}
	)

	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", config.SpotifyAuthRedirectURI)
	data.Set("client_id", config.SpotifyClientID)
	data.Set("client_secret", config.SpotifyClientSecret)

	err := request(data, &authResponse)

	return authResponse, err
}

func Token(refreshToken string) (TokenResponse, error) {
	var (
		tokenResponse TokenResponse
		data          = url.Values{}
	)

	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)
	data.Set("client_id", config.SpotifyClientID)
	data.Set("client_secret", config.SpotifyClientSecret)

	err := request(data, &tokenResponse)

	return tokenResponse, err
}

func request(data url.Values, response interface{}) error {
	// nolint: noctx
	r, err := http.NewRequest(http.MethodPost, getPath("api/token"), strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("error creating spotify auth request: %w", err)
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return fmt.Errorf("error during spotify auth request: %w", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return fmt.Errorf("error decoding spotify auth response: %w", err)
	}

	return nil
}

func getPath(action string) string {
	return endpointURI + action
}
