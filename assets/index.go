package assets

import (
	"github.com/Oppodelldog/spotify-sleep-timer/config"
	"github.com/go-playground/statics/static"
)

var (
	Files  *static.Files
	CSS    *static.Files
	Images *static.Files
	Fonts  *static.Files
)

func Init() {
	var err error

	cfg := &static.Config{
		UseStaticFiles: config.UseStaticFiles,
		FallbackToDisk: false,
		AbsPkgPath:     config.AbsoluteAssetsPath,
	}

	Files, err = newStaticAssets(cfg)
	if err != nil {
		panic(err)
	}

	CSS, err = newStaticCss(cfg)
	if err != nil {
		panic(err)
	}

	Images, err = newStaticIages(cfg)
	if err != nil {
		panic(err)
	}

	Fonts, err = newStaticFonts(cfg)
	if err != nil {
		panic(err)
	}
}
