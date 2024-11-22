package main

import (
	"time"

	"github.com/kingtingthegreat/top-fetch-cli/config"
	"github.com/kingtingthegreat/top-fetch-cli/convert"
	"github.com/kingtingthegreat/top-fetch-cli/env"
	"github.com/kingtingthegreat/top-fetch-cli/fatal"
	"github.com/kingtingthegreat/top-fetch-cli/fetch"
	"github.com/kingtingthegreat/top-fetch-cli/output"
)

func fetchAndDisplay(web bool) {
	var imageUrl, trackText string

	if web {
		imageUrl, trackText = fetch.WebFetch()
	} else {
		imageUrl, trackText = fetch.LocalFetch()
	}
	// log.Println("converting")
	ansiImage, err := convert.UrlToAnsi(imageUrl)
	if err != nil {
		fatal.Fatal(err.Error())
	}

	output.Output(ansiImage, trackText)
}

func main() {
	start := time.Now()
	env.LoadEnv()
	err := config.ParseArgs()
	if err != nil {
		fatal.Fatal(err.Error())
	}

	cfg := config.Config()

	if cfg.Timeout < 0 {
		fetchAndDisplay(cfg.Web)
	} else {
		c := make(chan bool)
		go func() { fetchAndDisplay(cfg.Web); c <- true }()
		for time.Now().Before(start.Add(time.Duration(cfg.Timeout) * time.Millisecond)) {
			select {
			case <-c:
				return
			default:
				continue
			}
		}
		fatal.Fatal("Exceed the ", cfg.Timeout, " millisecond time limit")
	}
}
