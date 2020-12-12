package player

type (
	ErrorResponse struct {
		Err Error `json:"error"`
	}
	Error struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	DevicesResponse struct {
		ErrorResponse
		Devices Devices `json:"devices"`
	}
	Devices []Device
	Device  struct {
		ID               string `json:"id"`
		IsActive         bool   `json:"is_active"`
		IsPrivateSession bool   `json:"is_private_session"`
		IsRestricted     bool   `json:"is_restricted"`
		Name             string `json:"name"`
		Type             string `json:"type"`
		VolumePercent    int    `json:"volume_percent"`
	}
)

func (devices Devices) GetActive() (Device, bool) {
	for _, device := range devices {
		if device.IsActive {
			return device, true
		}
	}

	return Device{}, false
}

func (e ErrorResponse) Error() string {
	return e.Err.Error()
}

func (e Error) Error() string {
	return e.Message
}
