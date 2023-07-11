package pkg

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	invalidNumberOfLines = errors.New("No required lines in file")
	badFormatOfLine      = errors.New("Bad format of file on line:")
)

func checkEvent(line string) {

}

func checkDataOfFile(dataOfFile string) (string, error) {
	lines := strings.Split(dataOfFile, "\r\n") //windows&linux
	if len(lines) < 3 {
		return "", invalidNumberOfLines
	}

	//first 3 lines check
	numberOfTables, err := strconv.Atoi(lines[0])
	if err != nil && numberOfTables < 1 {
		return lines[0], badFormatOfLine
	}

	price, err := strconv.Atoi(lines[2])
	if err != nil && price < 1 {
		return lines[2], badFormatOfLine
	}

	fmt.Println(numberOfTables) //

	//others lines check
	//var wg sync.WaitGroup

	for i := 3; i <= len(lines); i++ {
		//wg.Add(1)
		//i := i
		//go func() {
		//	defer wg.Done()
		//	checkEvent(lines[i])
		//}()
		checkEvent(lines[i])
	}

	//wg.Wait()

	for numberOfLine, line := range lines {
		fmt.Println(numberOfLine, ":", line)
	}

	return "", nil
}
