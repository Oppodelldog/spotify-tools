package config

import (
	"github.com/caarlos0/env/v6"
)

var BasePath string
var BindAddress string
var HttpAuth BasicHttpAuth
var UseStaticFiles bool
var AbsoluteAssetsPath string
var SpotifyClientID string
var SpotifyClientSecret string
var SpotifyAuthRedirectURI string

type BasicHttpAuth struct {
	Credentials string
	Realm       string
}

func Init() {
	var cfg struct {
		UseStaticFiles         bool   `env:"SLEEPTIMER_ASSETS_USE_STATIC_FILES" envDefault:"true"`
		AbsoluteAssetsPath     string `env:"SLEEPTIMER_ASSETS_ABSOLUTE_ASSETS_PATH" envDefault:""`
		BasePath               string `env:"SLEEPTIMER_BASE_PATH" envDefault:""`
		HTTPAuthCredentials    string `env:"SLEEPTIMER_HTTP_AUTH_CREDENTIALS" envDefault:""`
		HTTPAuthRealm          string `env:"SLEEPTIMER_HTTP_AUTH_REALM" envDefault:"Restricted"`
		BindAddress            string `env:"SLEEPTIMER_BIND_ADDRESS" envDefault:":8080"`
		SpotifyClientID        string `env:"SLEEPTIMER_SPOTIFY_CLIENT_ID" envDefault:""`
		SpotifyClientSecret    string `env:"SLEEPTIMER_SPOTIFY_CLIENT_SECRET" envDefault:""`
		SpotifyAuthRedirectURI string `env:"SLEEPTIMER_SPOTIFY_AUTH_REDIRECT_URI" envDefault:""`
	}

	err := env.Parse(&cfg)
	if err != nil {
		panic(err)
	}

	UseStaticFiles = cfg.UseStaticFiles
	AbsoluteAssetsPath = cfg.AbsoluteAssetsPath
	BasePath = cfg.BasePath
	HttpAuth = BasicHttpAuth{
		Credentials: cfg.HTTPAuthCredentials,
		Realm:       cfg.HTTPAuthRealm,
	}
	BindAddress = cfg.BindAddress
	SpotifyClientID = cfg.SpotifyClientID
	SpotifyClientSecret = cfg.SpotifyClientSecret
	SpotifyAuthRedirectURI = cfg.SpotifyAuthRedirectURI
}
