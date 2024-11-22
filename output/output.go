package output

import (
	"log"
	"os"

	"github.com/kingtingthegreat/top-fetch-cli/config"
)

func Output(cfg config.Config, ansiImage, trackText string) {

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
