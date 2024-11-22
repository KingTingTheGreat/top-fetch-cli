package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kingtingthegreat/ansi-converter/defaults"
)

type Config struct {
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
}

const WEB = "web"
const DIM = "dim="
const CHAR = "char="
const RATIO = "ratio="
const PATH = "path="
const FILE = "file="

const PADDING = "padding="
const WRAP = "wrap"

var OPTS = []string{WEB, DIM, CHAR, RATIO, PATH, FILE}

func ParseArgs() (Config, error) {
	config := Config{
		Web:                 false,
		TopFetchId:          os.Getenv("TOP_FETCH_ID"),
		SpotifyClientId:     os.Getenv("SPOTIFY_CLIENT_ID"),
		SpotifyClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		SpotifyAccessToken:  os.Getenv("SPOTIFY_ACCESS_TOKEN"),
		SpotifyRefreshToken: os.Getenv("SPOTIFY_REFRESH_TOKEN"),
		Dim:                 defaults.DEFAULT_DIM,
		Char:                defaults.DEFAULT_CHAR,
		FontRatio:           defaults.DEFAULT_RATIO,
		Path:                "source",
		File:                "",
		Wrap:                false,
	}

	for _, arg := range os.Args[1:] {
		arg = strings.ToLower(arg)
		for _, opt := range OPTS {
			if strings.HasPrefix(arg, opt) {
				val := strings.TrimLeft(arg, opt)
				switch opt {
				case WEB:
					config.Web = true
					if strings.HasPrefix(val, "=") {
						val = strings.TrimLeft(val, "=")
						config.TopFetchId = val
					}
				case DIM:
					newDim, err := strconv.ParseFloat(val, 64)
					if err != nil {
						return Config{}, fmt.Errorf("invalid dim")
					}
					config.Dim = newDim
				case CHAR:
					config.Char = val
				case RATIO:
					newRatio, err := strconv.ParseFloat(val, 64)
					if err != nil {
						return Config{}, fmt.Errorf("invalid ratio")
					}
					config.FontRatio = newRatio
				case FILE:
					config.File = val
				}

			}
		}
	}

	return config, nil
}
