package pkg

import (
	"log"
	"os"
	"strings"
)

func CheckArgs() (file string) {
	if len(os.Args) != 2 {
		log.Printf("Error: %v %v\r\n", invalidNumberOfArgs, os.Args)
		os.Exit(1)
	}

	if !strings.Contains(os.Args[1], ".txt") {
		log.Printf("Error: %v: %v\r\n", invalidFileFormat, os.Args[1])
		os.Exit(1)
	}

	return os.Args[1]
}
