package fatal

import (
	"log"
	"os"

	"github.com/kingtingthegreat/top-fetch-cli/config"
)

func Fatal(v ...any) {
	cfg := config.Config()
	if !cfg.Silent {
		log.Fatal(v...)
	} else {
		os.Exit(1)
	}
}
