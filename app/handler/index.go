package handler

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/Oppodelldog/spotify-sleep-timer/app/navigate"
	"github.com/Oppodelldog/spotify-sleep-timer/app/session"
	"github.com/Oppodelldog/spotify-sleep-timer/app/timer"
)

type indexPage struct {
	Username     string
	Timer        timer.Timer
	Due          int
	IsAuthorized bool
	AuthURL      string
	ClearURL     string
	Texts        indexPageTexts
	Controls     indexPageControls
}

type indexPageTexts struct {
	Title                  string
	Unauthorized           string
	ClearTimer             string
	SetTimer               string
	ToSpotifyAuthorization string
}

type indexPageControls []indexPageControl

type indexPageControl struct {
	Name  string
	Value int
}

// nolint: gomnd
func showIndexPage(t *template.Template, writer http.ResponseWriter, request *http.Request) {
	page := indexPage{
		AuthURL:  getPath("/auth"),
		ClearURL: getPath("/clear"),
		Controls: indexPageControls{
			{
				Name:  "+ 5 min",
				Value: 5,
			},
			{
				Name:  "+ 10 min",
				Value: 10,
			},
			{
				Name:  "+ 30 min",
				Value: 30,
			},
			{
				Name:  "+ 1 h",
				Value: 60,
			},
		},
		Texts: indexPageTexts{
			Title:                  "Sleep Timer",
			Unauthorized:           "You need to authorize this app",
			ClearTimer:             "Clear Timer",
			SetTimer:               "Set Timer",
			ToSpotifyAuthorization: "To Spotify Authorization",
		},
	}

	if user, err := session.GetSession(request); err == nil {
		page.Timer = user.Timer
		page.Due = int(user.Timer.AsDue().Due().Seconds())
		page.Username = user.Name
		page.IsAuthorized = true
	} else {
		switch {
		case errors.Is(session.ErrNoCookieFound, err):
		case errors.Is(session.ErrUserNotFound, err):
			navigate.FlushCookieRedirect(writer, request, indexPagePath())

			return
		}
	}

	err := t.Execute(writer, page)
	if err != nil {
		writeInternalServerErrorStatusCode(writer, err)

		return
	}
}
