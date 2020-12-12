package authorization

import (
	"encoding/json"
	"github.com/Oppodelldog/spotify-sleep-timer/config"
	"net/http"
	"net/url"
	"strings"
)

const endpointUri = "https://accounts.spotify.com/"

func AuthUrl() (string, error) {
	u, err := url.Parse(getPath("authorize"))
	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Add("client_id", config.SpotifyClientID)
	q.Add("response_type", "code")
	q.Add("redirect_uri", config.SpotifyAuthRedirectUri)
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
	data.Set("redirect_uri", config.SpotifyAuthRedirectUri)
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
	r, err := http.NewRequest(http.MethodPost, getPath("api/token"), strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return err
	}

	return nil
}

func getPath(action string) string {
	return endpointUri + action
}
