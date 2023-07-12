package main

import (
	"fmt"
	"yadro/internal/pkg"
)

func main() {
	fmt.Println(pkg.Parse(pkg.CheckArgs()))
}
