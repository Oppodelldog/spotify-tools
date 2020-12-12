package session

import (
	"errors"
	"net/http"

	"github.com/Oppodelldog/spotify-sleep-timer/app/storage"
)

var (
	ErrNoCookieFound = errors.New("no cookie set")
	ErrUserNotFound  = errors.New("user not found by session cookie")
)

func GetSession(request *http.Request) (storage.Session, error) {
	c, err := getSessionCookie(request)
	if err != nil {
		return storage.Session{}, ErrNoCookieFound
	}

	session, b := storage.Get(c.Value)
	if !b {
		return storage.Session{}, ErrUserNotFound
	}

	return *session, nil
}

func SetSession(session storage.Session) {
	storage.Set(session)
}
