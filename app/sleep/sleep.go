package sleep

import (
	"errors"
	"fmt"

	"github.com/Oppodelldog/spotify-sleep-timer/spotify/player"
)

var ErrNoActiveDeviceFound = errors.New("no active device found")

func pause(token string) error {
	deviceID, err := getActiveDeviceID(token)
	if err != nil {
		return fmt.Errorf("cannot pause: %w", err)
	}

	err = player.Pause(token, deviceID)
	if err != nil {
		return fmt.Errorf("cannot pause: %w", err)
	}

	return nil
}

func getActiveDeviceID(token string) (string, error) {
	devices, err := player.GetDevices(token)
	if err != nil {
		return "", fmt.Errorf("error getting active device: %w", err)
	}

	device, found := devices.GetActive()
	if !found {
		return "", ErrNoActiveDeviceFound
	}

	return device.ID, nil
}
