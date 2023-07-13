package pkg

import (
	"os"
	"strings"
)

func CheckArgs() (file string, err error) {
	if len(os.Args) != 2 {
		return strings.Join(os.Args, " "), invalidNumberOfArgs
		//log.Printf("Error: %v %v\r\n", invalidNumberOfArgs, os.Args)
		//os.Exit(1)
	}

	if !strings.Contains(os.Args[1], ".txt") {
		return os.Args[1], invalidFileFormat
		//log.Printf("Error: %v: %v\r\n", invalidFileFormat, os.Args[1])
		//os.Exit(1)
	}

	return os.Args[1], nil
}
