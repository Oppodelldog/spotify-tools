package navigate

import (
	"net/http"

	"github.com/Oppodelldog/spotify-sleep-timer/app/session"
)

func FlushCookieRedirect(writer http.ResponseWriter, request *http.Request, url string) {
	session.ClearSessionCookie(writer)
	Redirect(writer, request, url)
}

func Redirect(writer http.ResponseWriter, request *http.Request, url string) {
	http.Redirect(writer, request, url, http.StatusFound)
}
