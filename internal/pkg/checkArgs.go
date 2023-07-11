package pkg

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	invalidNumberOfArgs = errors.New("invalid number of arguments in command line")
	invalidFileFormat   = errors.New("file must has \"txt\" format")
)

func CheckArgs() (file string) {
	if len(os.Args) != 2 {
		fmt.Printf("Error: %v %v\r\n", invalidNumberOfArgs, os.Args)
		os.Exit(1)
	}

	if !strings.Contains(os.Args[1], ".txt") {
		fmt.Printf("Error: %v: %v\r\n", invalidFileFormat, os.Args[1])
		os.Exit(1)
	}

	return os.Args[1]
}
