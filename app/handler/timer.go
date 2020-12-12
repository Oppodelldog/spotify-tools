package handler

import (
	"fmt"
	"github.com/Oppodelldog/spotify-sleep-timer/app/navigate"
	"github.com/Oppodelldog/spotify-sleep-timer/app/session"
	"github.com/Oppodelldog/spotify-sleep-timer/app/storage"
	"github.com/Oppodelldog/spotify-sleep-timer/app/timer"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func setTimer(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		writeBadRequestStatusCode(writer, err)

		return
	}

	s, err := session.GetSession(request)
	if err != nil {
		navigate.FlushCookieRedirectToIndex(writer, request)

		return
	}

	s, err = makeTimer(s, request.Form)
	if err != nil {
		writeBadRequestStatusCode(writer, err)

		return
	}

	session.SetSession(s)
	navigate.RedirectToIndex(writer, request)
}

func makeTimer(s storage.Session, form url.Values) (storage.Session, error) {
	hr, min, err := parseInts(form["h"][0], form["m"][0])
	if err != nil {
		return storage.Session{}, err
	}
	s.Timer = timer.Timer{
		Hours:   hr,
		Minutes: min,
		IsSet:   true,
		SetAt:   time.Now(),
	}

	return s, nil
}

func parseInts(a, b string) (int, int, error) {
	var err error
	res := make([]int, 3)
	for i, v := range []string{a, b} {
		res[i], err = strconv.Atoi(v)
		if err != nil {
			return 0, 0, err
		}
		if res[i] < 0 || res[i] > 59 {
			return 0, 0, fmt.Errorf("invalid number input: %v", res[i])
		}
	}

	return res[0], res[1], nil
}
