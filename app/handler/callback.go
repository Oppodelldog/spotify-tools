package handler

import (
	"html/template"
	"net/http"
	"time"

	"github.com/Oppodelldog/spotify-sleep-timer/app/navigate"
	"github.com/Oppodelldog/spotify-sleep-timer/app/session"
	"github.com/Oppodelldog/spotify-sleep-timer/app/storage"
	"github.com/Oppodelldog/spotify-sleep-timer/app/timer"
	"github.com/Oppodelldog/spotify-sleep-timer/spotify/authorization"
	"github.com/Oppodelldog/spotify-sleep-timer/spotify/me"
)

type callbackPage struct {
	Title    string
	SubTitle string
	Code     string
	Error    string
	State    string
}

func redirectToSpotifyAuthPage(writer http.ResponseWriter, request *http.Request) {
	authURL, err := authorization.AuthURL()
	if err != nil {
		writeInternalServerErrorStatusCode(writer, err)

		return
	}

	navigate.Redirect(writer, request, authURL)
}

func spotifyAuthCallback(t *template.Template, writer http.ResponseWriter, request *http.Request) {
	var (
		code   = request.URL.Query().Get("code")
		errMsg = request.URL.Query().Get("error")
	)

	if code != "" {
		authResponse, err := authorization.Auth(code)
		if err != nil {
			writeInternalServerErrorStatusCode(writer, err)

			return
		}

		meData, err := me.Profile(authResponse.AccessToken)
		if err != nil {
			writeInternalServerErrorStatusCode(writer, err)

			return
		}

		user := storage.Session{
			ID:   meData.ID,
			Name: meData.DisplayName,
			Spotify: storage.Spotify{
				AccessToken:  authResponse.AccessToken,
				RefreshToken: authResponse.RefreshToken,
				ExpiresIn:    authResponse.ExpiresIn,
				RefreshedAt:  time.Now(),
			},
			Timer: timer.Timer{},
		}

		sessionID := storage.Set(user)

		session.SetSessionCookie(writer, sessionID)
		navigate.RedirectToIndex(writer, request)

		return
	}

	err := t.Execute(writer, callbackPage{
		Title:    "Sleep Timer - Authorization",
		SubTitle: "Something went wrong",
		Error:    errMsg,
	})
	if err != nil {
		writeInternalServerErrorStatusCode(writer, err)

		return
	}
}
