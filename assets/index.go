package assets

import (
	"github.com/Oppodelldog/spotify-sleep-timer/config"
	"github.com/go-playground/statics/static"
)

var Files *static.Files
var CSS *static.Files

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
}
