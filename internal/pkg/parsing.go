package pkg

import (
	"fmt"
	"os"
)

func Parse(file string) {
	dataOfFile, err := os.ReadFile("internal/pkg/testFile/" + file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	line, err := checkDataOfFile(string(dataOfFile))
	if err != nil {
		fmt.Printf("%v%v\r\n", err, line)
		os.Exit(1)
	}

	fmt.Println(string(dataOfFile))
}
