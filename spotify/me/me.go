package me

import (
	"github.com/Oppodelldog/spotify-sleep-timer/spotify"
	"net/http"
)

const urlMe = "https://api.spotify.com/v1/me"

func Profile(token string) (Me, error) {
	var (
		me Me
	)

	err := spotify.Request(token, http.MethodGet, urlMe, nil, nil, &me)
	if err != nil {
		return Me{}, err
	}

	return me, nil
}
