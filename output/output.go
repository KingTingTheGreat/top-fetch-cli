package output

import (
	"log"
	"os"
	"strings"

	"github.com/kingtingthegreat/top-fetch-cli/config"
)

func centerTrackText(trackText string, dim int) string {
	if len(trackText) > dim {
		return trackText
	}

	leftPad := (dim - len(trackText)) / 2
	rightPad := dim - len(trackText) - leftPad
	return strings.Repeat(" ", leftPad) + trackText + strings.Repeat(" ", rightPad)
}

func Output(cfg config.Config, ansiImage, trackText string) {
	trackText = centerTrackText(trackText, int(cfg.Dim))

	// write to desired output
	if cfg.File != "" {
		outputFile, err := WriteToFile(cfg, ansiImage, trackText)
		if err != nil {
			log.Fatal(err.Error())
		}
		os.Stdout.WriteString(outputFile)
	} else {
		os.Stdout.WriteString(ansiImage + trackText)
	}
}
