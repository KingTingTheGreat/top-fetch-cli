package output

import (
	"os"
	"strings"

	"github.com/kingtingthegreat/top-fetch-cli/config"
	"github.com/kingtingthegreat/top-fetch-cli/fatal"
)

func centerTrackText(trackText string, dim int) string {
	if len(trackText) > dim {
		return trackText
	}

	leftPad := (dim - len(trackText)) / 2
	rightPad := dim - len(trackText) - leftPad
	return strings.Repeat(" ", leftPad) + trackText + strings.Repeat(" ", rightPad)
}

func Output(ansiImage, trackText string) {
	cfg := config.Config()
	trackText = centerTrackText(trackText, int(cfg.Dim))

	// write to desired output
	if cfg.File != "" {
		outputFile, err := WriteToFile(ansiImage, trackText)
		if err != nil {
			fatal.Fatal(err.Error())
		}
		os.Stdout.WriteString(outputFile)
	} else {
		os.Stdout.WriteString(ansiImage + trackText)
	}
}
