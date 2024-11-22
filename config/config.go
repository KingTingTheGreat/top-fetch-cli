package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kingtingthegreat/ansi-converter/defaults"
)

type config struct {
	Web                 bool
	TopFetchId          string
	SpotifyClientId     string
	SpotifyClientSecret string
	SpotifyAccessToken  string
	SpotifyRefreshToken string
	Dim                 float64
	Char                string
	FontRatio           float64
	Path                string
	File                string
	Wrap                bool
	Timeout             int
	Silent              bool
}

const (
	WEB     = "web"
	DIM     = "dim="
	CHAR    = "char="
	RATIO   = "ratio="
	PATH    = "path="
	FILE    = "file="
	TIMEOUT = "timeout="
	SILENT  = "silent"
)

const PADDING = "padding="
const WRAP = "wrap"

var OPTS = []string{WEB, DIM, CHAR, RATIO, PATH, FILE, TIMEOUT}

var cfg config = config{
	Web:       false,
	Dim:       defaults.DEFAULT_DIM,
	Char:      defaults.DEFAULT_CHAR,
	FontRatio: defaults.DEFAULT_RATIO,
	Path:      "source",
	File:      "",
	Wrap:      false,
	Timeout:   -1,
}

func ParseArgs() error {
	// configure args
	for _, arg := range os.Args[1:] {
		for _, opt := range OPTS {
			if strings.HasPrefix(arg, opt) {
				val := strings.TrimLeft(arg, opt)
				switch opt {
				case WEB:
					cfg.Web = true
					if strings.HasPrefix(val, "=") {
						val = strings.TrimLeft(val, "=")
						cfg.TopFetchId = val
					}
				case DIM:
					newDim, err := strconv.ParseFloat(val, 64)
					if err != nil {
						return fmt.Errorf("invalid dim")
					}
					cfg.Dim = newDim
				case CHAR:
					cfg.Char = val
				case RATIO:
					newRatio, err := strconv.ParseFloat(val, 64)
					if err != nil {
						return fmt.Errorf("invalid ratio")
					}
					cfg.FontRatio = newRatio
				case FILE:
					cfg.File = val
				case TIMEOUT:
					newTimeout, err := strconv.Atoi(val)
					if err != nil {
						return fmt.Errorf("invalid timeout")
					}
					cfg.Timeout = newTimeout
				case SILENT:
					cfg.Silent = true
				}

			}
		}
	}

	cfg.TopFetchId = os.Getenv("TOP_FETCH_ID")
	cfg.SpotifyClientId = os.Getenv("SPOTIFY_CLIENT_ID")
	cfg.SpotifyClientSecret = os.Getenv("SPOTIFY_CLIENT_SECRET")
	cfg.SpotifyAccessToken = os.Getenv("SPOTIFY_ACCESS_TOKEN")
	cfg.SpotifyRefreshToken = os.Getenv("SPOTIFY_REFRESH_TOKEN")

	return nil
}

func Config() *config {
	return &cfg
}
