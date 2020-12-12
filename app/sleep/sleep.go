package sleep

import (
	"fmt"
	"github.com/Oppodelldog/spotify-sleep-timer/spotify/player"
)

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
		return "", err
	}

	device, found := devices.GetActive()
	if !found {
		return "", fmt.Errorf("no active device found")
	}

	return device.ID, nil
}
