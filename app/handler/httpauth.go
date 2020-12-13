package handler

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/Oppodelldog/spotify-sleep-timer/config"
)

func httpAuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, config.HttpAuth.Realm))

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

		if config.HttpAuth.Credentials != "" && string(up) != config.HttpAuth.Credentials {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)

			return
		}

		h.ServeHTTP(w, r)
	})
}
