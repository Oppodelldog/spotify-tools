package handler

import (
	"html/template"
	"net/http"
	"time"

	"github.com/Oppodelldog/spotify-sleep-timer/app/storage"
	"github.com/Oppodelldog/spotify-sleep-timer/app/timer"
)

type (
	adminPage struct {
		Title    string
		Sessions []adminPageSession
	}
	adminPageSession struct {
		ID       string
		Name     string
		Timer    timer.Timer
		Spotify  storage.Spotify
		TimerDue time.Duration
		TokenDue time.Duration
	}
)

func adminView(t *template.Template, writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		writeInternalServerErrorStatusCode(writer, err)

		return
	}

	data := adminPage{
		Title:    "Sleep Timer - Admin",
		Sessions: getSessions(storage.All()),
	}

	err = t.Execute(writer, data)
	if err != nil {
		writeInternalServerErrorStatusCode(writer, err)

		return
	}
}

func getSessions(all []storage.Session) []adminPageSession {
	var sessions = make([]adminPageSession, 0, len(all))

	for _, s := range all {
		sessions = append(
			sessions,
			adminPageSession{
				ID:       s.ID[:4] + "...",
				Name:     s.Name[:3] + "...",
				TimerDue: s.Timer.AsDue().Due(),
				TokenDue: s.Spotify.TokenDue().Due(),
				Timer:    s.Timer,
				Spotify: storage.Spotify{
					AccessToken:  s.Spotify.AccessToken[:10] + "...",
					RefreshToken: s.Spotify.RefreshToken[:10] + "...",
					ExpiresIn:    s.Spotify.ExpiresIn,
					RefreshedAt:  s.Spotify.RefreshedAt,
					RefreshErr:   s.Spotify.RefreshErr,
				},
			},
		)
	}

	return sessions
}
