package pkg

import (
	"os"
	"strings"
)

func CheckArgs() (file string, err error) {
	if len(os.Args) != 2 {
		return strings.Join(os.Args, " "), invalidNumberOfArgs
	}

	if !strings.Contains(os.Args[1], ".txt") {
		return os.Args[1], invalidFileFormat
	}

	return os.Args[1], nil
}
