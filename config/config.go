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
	PaddingTop          int
	PaddingRight        int
	PaddingBottom       int
	PaddingLeft         int
}

const (
	WEB            = "web"
	DIM            = "dim"
	CHAR           = "char"
	RATIO          = "ratio"
	PATH           = "path"
	FILE           = "file"
	TIMEOUT        = "timeout"
	SILENT         = "silent"
	PADDING        = "p"
	PADDING_TOP    = "pT"
	PADDING_RIGHT  = "pR"
	PADDING_BOTTOM = "pB"
	PADDING_LEFT   = "pL"
)

const WRAP = "wrap"

var cfg config = config{
	Web:           false,
	Dim:           defaults.DEFAULT_DIM,
	Char:          defaults.DEFAULT_CHAR,
	FontRatio:     defaults.DEFAULT_RATIO,
	Path:          "source",
	File:          "",
	Wrap:          false,
	Timeout:       -1,
	PaddingTop:    0,
	PaddingRight:  0,
	PaddingBottom: 0,
	PaddingLeft:   0,
}

func ParseArgs() error {
	// configure args
	for _, argValStr := range os.Args[1:] {
		argVal := strings.SplitN(argValStr, "=", 2)
		arg := argVal[0]
		val := ""
		if len(argVal) > 1 {
			val = argVal[1]
		}
		switch arg {
		case WEB:
			cfg.Web = true
			if val != "" {
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
		case PADDING:
			newPadding, err := strconv.Atoi(val)
			if err != nil {
				return fmt.Errorf("invalid padding")
			}
			cfg.PaddingTop = newPadding
			cfg.PaddingRight = newPadding
			cfg.PaddingBottom = newPadding
			cfg.PaddingLeft = newPadding
		case PADDING_TOP:
			newPaddingTop, err := strconv.Atoi(val)
			if err != nil {
				return fmt.Errorf("invalid padding top")
			}
			cfg.PaddingTop = newPaddingTop
		case PADDING_RIGHT:
			newPaddingRight, err := strconv.Atoi(val)
			if err != nil {
				return fmt.Errorf("invalid padding right")
			}
			cfg.PaddingRight = newPaddingRight
		case PADDING_BOTTOM:
			newPaddingBottom, err := strconv.Atoi(val)
			if err != nil {
				return fmt.Errorf("invalid padding bottom")
			}
			cfg.PaddingBottom = newPaddingBottom
		case PADDING_LEFT:
			newPaddingLeft, err := strconv.Atoi(val)
			if err != nil {
				return fmt.Errorf("invalid padding left")
			}
			cfg.PaddingLeft = newPaddingLeft
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
