package output

import (
	"os"
	"strings"

	"github.com/kingtingthegreat/top-fetch-cli/config"
	"github.com/kingtingthegreat/top-fetch-cli/fatal"
)

func centerTrackText(trackText string, dim int, left, right int) string {
	if len(trackText) > dim {
		return trackText
	}

	width := dim + right + left
	leftPad := (width - len(trackText)) / 2
	rightPad := width - len(trackText) - leftPad
	return strings.Repeat(" ", leftPad) + trackText + strings.Repeat(" ", rightPad)
}

func Output(ansiImage, trackText string) {
	cfg := config.Config()
	trackText = centerTrackText(trackText, int(cfg.Dim),
		cfg.PaddingLeft+cfg.MarginLeft, cfg.PaddingRight+cfg.MarginRight)

	outputString := ansiImage + "\n" + trackText + "\n" + strings.Repeat("\n", cfg.MarginBottom)

	// write to desired output
	if cfg.File != "" {
		outputFile, err := WriteToFile(outputString)
		if err != nil {
			fatal.Fatal(err.Error())
		}
		os.Stdout.WriteString(outputFile)
	} else {
		os.Stdout.WriteString(outputString)
	}
}
