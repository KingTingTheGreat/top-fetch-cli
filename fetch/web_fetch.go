package fetch

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/kingtingthegreat/top-fetch-cli/config"
)

func WebFetch(cfg config.Config) (string, string) {
	if cfg.TopFetchId == "" {
		log.Fatal("TopFetch id is not set")
	}

	res, err := http.Get("https://top-fetch.vercel.app/track?id=" + cfg.TopFetchId)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	info := string(bodyBytes)

	infoList := strings.Split(info, "\x1d")
	if len(infoList) == 1 {
		log.Fatal(info)
	} else if len(infoList) != 2 {
		log.Fatal("something went wrong!! please contact me")
	}

	return infoList[0], infoList[1]
}
