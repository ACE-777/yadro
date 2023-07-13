package main

import (
	"fmt"
	"log"
	"os"

	"yadro/internal/pkg"
)

func main() {
	file, err := pkg.CheckArgs()
	if err != nil {
		log.Printf("Error: %v: %v\r\n", err, file)
		os.Exit(1)
	}

	result, err := pkg.Parse(file)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(result)
}
