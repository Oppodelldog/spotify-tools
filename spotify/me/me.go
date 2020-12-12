package me

import (
	"fmt"
	"net/http"

	"github.com/Oppodelldog/spotify-sleep-timer/spotify"
)

const urlMe = "https://api.spotify.com/v1/me"

func Profile(token string) (Me, error) {
	var me Me

	err := spotify.Request(token, http.MethodGet, urlMe, nil, nil, &me)
	if err != nil {
		return Me{}, fmt.Errorf("error getting profile: %w", err)
	}

	return me, nil
}
