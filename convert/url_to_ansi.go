package convert

import (
	"fmt"
	"image"
	"net/http"

	"github.com/kingtingthegreat/ansi-converter/converter"
	"github.com/kingtingthegreat/top-fetch-cli/config"
)

func UrlToAnsi(url string, config config.Config) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to get image from url")
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return "", fmt.Errorf("cannot decode image from response")
	}

	return converter.Convert(img, config.Dim, config.Char, config.FontRatio), nil
}
