package pkg

import (
	"fmt"
	"log"
	"os"
)

func Parse(file string) {
	dataOfFile, err := os.ReadFile("internal/pkg/testFile/" + file)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	line, err := checkDataOfFile(string(dataOfFile))
	if err != nil {
		log.Printf("%v %v\r\n", err, line)
		os.Exit(1)
	}

	//3+ строк чекать и парсиить
	// чекать что номер стола не превосходит имующийся

	fmt.Println(string(dataOfFile))
}
