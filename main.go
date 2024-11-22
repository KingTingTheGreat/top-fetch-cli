package main

import (
	"log"

	"github.com/kingtingthegreat/top-fetch-cli/config"
	"github.com/kingtingthegreat/top-fetch-cli/convert"
	"github.com/kingtingthegreat/top-fetch-cli/env"
	"github.com/kingtingthegreat/top-fetch-cli/fetch"
	"github.com/kingtingthegreat/top-fetch-cli/output"
)

type Result struct {
	Image string
	Err   error
}

func main() {
	env.LoadEnv()
	cfg, err := config.ParseArgs()
	if err != nil {
		log.Fatal(err.Error())
	}

	var imageUrl, trackText string

	if cfg.Web {
		imageUrl, trackText = fetch.WebFetch(cfg)
	} else {
		imageUrl, trackText = fetch.LocalFetch(cfg)
	}
	// log.Println("converting")
	ansiImage, err := convert.UrlToAnsi(cfg, imageUrl)
	if err != nil {
		log.Fatal(err.Error())
	}

	output.Output(cfg, ansiImage, trackText)
}
