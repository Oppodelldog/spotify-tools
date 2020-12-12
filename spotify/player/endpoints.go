package player

import (
	"fmt"
	"net/http"

	"github.com/Oppodelldog/spotify-sleep-timer/spotify"
)

const endpointURI = "https://api.spotify.com/v1/me/player/"

func Pause(token, deviceID string) error {
	var (
		response ErrorResponse
		query    = map[string]string{
			"device_id": deviceID,
		}
	)

	err := spotify.Request(token, http.MethodPut, getPath("pause"), nil, query, &response)
	if err != nil {
		return fmt.Errorf("error pausing playback: %w", err)
	}

	return nil
}

func GetDevices(token string) (Devices, error) {
	var response DevicesResponse

	err := spotify.Request(token, http.MethodGet, getPath("devices"), nil, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("error loading devices: %w", err)
	}

	return response.Devices, nil
}

func getPath(action string) string {
	return endpointURI + action
}
