package assets

import (
	"embed"
)

var (
	//go:embed files
	Files embed.FS

	//go:embed css
	CSS embed.FS

	//go:embed img
	Images embed.FS

	//go:embed fonts
	Fonts embed.FS
)
