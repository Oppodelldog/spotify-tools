package session

import (
	"net/http"
	"time"
)

const (
	sessionCookie  = "session"
	cookiePath     = "/"
	cookieLifetime = 2 * time.Hour
)

func getSessionCookie(request *http.Request) (*http.Cookie, error) {
	return request.Cookie(sessionCookie)
}

func SetSessionCookie(writer http.ResponseWriter, sessionID string) {
	http.SetCookie(writer, &http.Cookie{
		Name:       sessionCookie,
		Value:      sessionID,
		Path:       cookiePath,
		Domain:     "",
		Expires:    time.Now().Add(cookieLifetime),
		RawExpires: "",
		MaxAge:     0,
		Secure:     false,
		HttpOnly:   false,
		SameSite:   0,
		Raw:        "",
		Unparsed:   nil,
	})
}

func ClearSessionCookie(writer http.ResponseWriter) {
	http.SetCookie(writer, &http.Cookie{
		Name:   sessionCookie,
		Value:  "",
		Path:   cookiePath,
		MaxAge: -1,
	})
}
