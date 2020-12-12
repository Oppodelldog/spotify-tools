package handler

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/Oppodelldog/spotify-sleep-timer/config"
)

func wrapAuthentication(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Basic ") {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)

			return
		}

		up, err := base64.StdEncoding.DecodeString(auth[6:])
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)

			return
		}

		if config.HTTPAuth != "" && string(up) != config.HTTPAuth {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)

			return
		}

		h.ServeHTTP(w, r)
	}
}
