package main

import (
	"log"
	"time"

	"github.com/kingtingthegreat/top-fetch-cli/config"
	"github.com/kingtingthegreat/top-fetch-cli/convert"
	"github.com/kingtingthegreat/top-fetch-cli/env"
	"github.com/kingtingthegreat/top-fetch-cli/fetch"
	"github.com/kingtingthegreat/top-fetch-cli/output"
)

func fetchAndDisplay(cfg config.Config) {
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

func main() {
	start := time.Now()
	env.LoadEnv()
	cfg, err := config.ParseArgs()
	if err != nil {
		log.Fatal(err.Error())
	}

	if cfg.Timeout < 0 {
		fetchAndDisplay(cfg)
	} else {
		c := make(chan bool)
		go func() { fetchAndDisplay(cfg); c <- true }()
		for time.Now().Before(start.Add(time.Duration(cfg.Timeout) * time.Millisecond)) {
			select {
			case <-c:
				return
			default:
				continue
			}
		}
		log.Fatal("Exceed the ", cfg.Timeout, " millisecond time limit")
	}
}
