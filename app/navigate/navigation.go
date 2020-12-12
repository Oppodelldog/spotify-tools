package navigate

import (
	"net/http"

	"github.com/Oppodelldog/spotify-sleep-timer/app/session"
)

func FlushCookieRedirectToIndex(writer http.ResponseWriter, request *http.Request) {
	session.ClearSessionCookie(writer)
	RedirectToIndex(writer, request)
}

func RedirectToIndex(writer http.ResponseWriter, request *http.Request) {
	Redirect(writer, request, "/")
}

func Redirect(writer http.ResponseWriter, request *http.Request, url string) {
	http.Redirect(writer, request, url, http.StatusFound)
}
