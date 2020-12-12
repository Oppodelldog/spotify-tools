package player

import (
	"github.com/Oppodelldog/spotify-sleep-timer/spotify"
	"net/http"
)

const endpointUri = "https://api.spotify.com/v1/me/player/"

func Pause(token, deviceID string) error {
	var (
		response ErrorResponse
		query    = map[string]string{
			"device_id": deviceID,
		}
	)

	err := spotify.Request(token, http.MethodPut, getPath("pause"), nil, query, &response)
	if err != nil {
		return err
	}

	return nil
}

func GetDevices(token string) (Devices, error) {
	var response DevicesResponse

	err := spotify.Request(token, http.MethodGet, getPath("devices"), nil, nil, &response)
	if err != nil {
		return nil, err
	}

	return response.Devices, nil
}

func getPath(action string) string {
	return endpointUri + action
}
