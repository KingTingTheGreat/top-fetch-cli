package fatal

import (
	"log"
	"os"
)

func Fatal(silent bool, v ...any) {
	if !silent {
		log.Fatal(v...)
	} else {
		os.Exit(1)
	}
}
