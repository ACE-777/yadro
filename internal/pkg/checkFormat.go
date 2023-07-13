package pkg

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

func checkPattern(input string, typeOfCheck string) bool {
	var checkLine string
	switch typeOfCheck {
	case "time":
		checkLine = "^[0-9:]+$"
	case "ID":
		checkLine = "^[1-4]+$"
	case "name":
		checkLine = "^[a-z0-9_-]+$"
	}

	r := regexp.MustCompile(checkLine)
	valid := r.MatchString(input)
	return valid
}

func checkEvent(line []string, numberOfTables int) bool {
	if len(line) > 4 {
		return false
	}

	if !checkPattern(line[0], "time") {
		return false
	}

	if !checkPattern(line[1], "ID") {
		return false
	}

	if !checkPattern(line[2], "name") {
		return false
	}

	if len(line) == 4 {
		IDTable, err := strconv.Atoi(line[3])
		if err != nil {
			return false
		}

		if IDTable < 0 || IDTable > numberOfTables {
			return false
		}
	}

	return true
}

func checkDataOfFirstThreeLinesInFile(dataOfFile string) ([]string, int, time.Time, time.Time, int, string, error) {
	lines := strings.Split(dataOfFile, "\r\n") //windows&linux
	if len(lines) < 3 {
		return nil, 0, time.Time{}, time.Time{}, 0, "", InvalidNumberOfLines
	}

	numberOfTables, err := strconv.Atoi(lines[0])
	if err != nil || numberOfTables < 1 {
		return nil, 0, time.Time{}, time.Time{}, 0, lines[0], BadFormatOfLine
	}

	secondLine := strings.Split(lines[1], " ")
	if len(secondLine) != 2 || len(secondLine[0]) != 5 || len(secondLine[1]) != 5 {
		return nil, 0, time.Time{}, time.Time{}, 0, lines[1], BadFormatOfLine
	}

	if !checkPattern(secondLine[0], "time") || !checkPattern(secondLine[1], "time") {
		return nil, 0, time.Time{}, time.Time{}, 0, lines[1], BadFormatOfLine
	}

	price, err := strconv.Atoi(lines[2])
	if err != nil || price < 1 {
		return nil, 0, time.Time{}, time.Time{}, 0, lines[2], BadFormatOfLine
	}

	openTime, err := time.Parse("15:04", secondLine[0])
	if err != nil {
		return lines, numberOfTables, time.Time{}, time.Time{}, price, "", err
	}

	closeTime, err := time.Parse("15:04", secondLine[1])
	if err != nil {
		return lines, numberOfTables, time.Time{}, time.Time{}, price, "", err
	}

	return lines, numberOfTables, openTime, closeTime, price, "", nil
}
