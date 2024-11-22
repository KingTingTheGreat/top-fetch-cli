package fetch

import (
	"io"
	"net/http"
	"strings"

	"github.com/kingtingthegreat/top-fetch-cli/config"
	"github.com/kingtingthegreat/top-fetch-cli/fatal"
)

func WebFetch(cfg config.Config) (string, string) {
	if cfg.TopFetchId == "" {
		fatal.Fatal(cfg.Silent, "TopFetch id is not set")
	}

	res, err := http.Get("https://top-fetch.vercel.app/track?id=" + cfg.TopFetchId)
	if err != nil {
		fatal.Fatal(cfg.Silent, err.Error())
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fatal.Fatal(cfg.Silent, err.Error())
	}
	info := string(bodyBytes)

	infoList := strings.Split(info, "\x1d")
	if len(infoList) == 1 {
		fatal.Fatal(cfg.Silent, info)
	} else if len(infoList) != 2 {
		fatal.Fatal(cfg.Silent, "something went wrong!! please contact me")
	}

	return infoList[0], infoList[1]
}
