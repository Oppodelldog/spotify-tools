package handler

import (
	"github.com/Oppodelldog/spotify-sleep-timer/assets"
	"github.com/Oppodelldog/spotify-sleep-timer/config"
	"github.com/gorilla/mux"
	"net/http"
	"path"
)

func Router() http.Handler {
	r := mux.NewRouter()

	r.Path(getPath("/")).
		Handler(withTemplate("index.html", showIndexPage)).
		Methods(http.MethodGet)

	r.Path(getPath("/")).
		Handler(http.HandlerFunc(setTimer)).
		Methods(http.MethodPost)

	r.Path(getPath("/clear")).
		Handler(http.HandlerFunc(clearTimer)).
		Methods(http.MethodGet)

	r.Path(getPath("/auth")).
		Handler(http.HandlerFunc(redirectToSpotifyAuthPage)).
		Methods(http.MethodGet)

	r.Path(getPath("/callback")).
		Handler(withTemplate("callback.html", spotifyAuthCallback)).
		Methods(http.MethodGet)

	r.Path(getPath("/admin")).
		Handler(withTemplate("admin.html", adminView)).
		Methods(http.MethodGet)

	r.PathPrefix(getPath("/assets")).
		Handler(http.FileServer(assets.Files.FS())).
		Methods(http.MethodGet)

	return r
}

func getPath(s string) string {
	return path.Join(config.BasePath, s)
}
