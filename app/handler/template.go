package handler

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	"github.com/Oppodelldog/spotify-sleep-timer/config"

	"github.com/Oppodelldog/spotify-sleep-timer/assets"
	"github.com/Oppodelldog/spotify-sleep-timer/logger"
)

type templateHandler func(t *template.Template, writer http.ResponseWriter, request *http.Request)

func withTemplate(filename string, h templateHandler) http.HandlerFunc {
	const assetPath = "files/"

	return func(w http.ResponseWriter, r *http.Request) {
		t, err := loadTemplate(path.Join(assetPath, filename))
		if err != nil {
			writeInternalServerErrorStatusCode(w, err)

			return
		}

		h(t, w, r)
	}
}

func writeInternalServerErrorStatusCode(writer http.ResponseWriter, err error) {
	writeError(writer, err, http.StatusInternalServerError)
}
func writeBadRequestStatusCode(writer http.ResponseWriter, err error) {
	writeError(writer, err, http.StatusInternalServerError)
}

func writeError(writer http.ResponseWriter, err error, statusCode int) {
	writer.WriteHeader(statusCode)

	_, writeErr := writer.Write([]byte(err.Error()))
	if writeErr != nil {
		logger.Std.Errorf("error writing error (%s) to client: %v\n", err.Error(), writeErr)
	}
}

func loadTemplate(filename string) (*template.Template, error) {
	f, err := assets.Files.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("cannot open template file '%s': %w", filename, err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("cannot read template file '%s': %w", filename, err)
	}

	b = substituteBasePath(b)

	t := template.New("index")
	t.Funcs(template.FuncMap{
		"attrNot": func(v bool, attr string) template.HTMLAttr {
			if !v {
				return template.HTMLAttr(attr) // nolint: gosec
			}

			return ""
		},
	})

	t, err = t.Parse(string(b))
	if err != nil {
		return nil, fmt.Errorf("cannot parse template file '%s': %w", filename, err)
	}

	return t, nil
}

func substituteBasePath(b []byte) []byte {
	b = []byte(strings.ReplaceAll(string(b), "{BASE_PATH}", config.BasePath))

	return b
}
