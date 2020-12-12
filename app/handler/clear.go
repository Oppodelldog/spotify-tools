package handler

import (
	"net/http"

	"github.com/Oppodelldog/spotify-sleep-timer/app/navigate"
	session2 "github.com/Oppodelldog/spotify-sleep-timer/app/session"
	"github.com/Oppodelldog/spotify-sleep-timer/app/timer"
)

func clearTimer(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		writeInternalServerErrorStatusCode(writer, err)

		return
	}

	session, err := session2.GetSession(request)
	if err != nil {
		navigate.FlushCookieRedirectToIndex(writer, request)

		return
	}

	session.Timer = timer.Timer{}

	session2.SetSession(session)

	navigate.RedirectToIndex(writer, request)
}
