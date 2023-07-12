package pkg

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type SpotOfTable struct {
	client    string
	timeStart time.Time
	busy      bool
}

type Clients struct {
	ID   int
	name string
}

type FinalProfit struct {
	wholeTime float64
	profit    int
	IDOfTable int
}

func Parse(file string) string {
	dataOfFile, err := os.ReadFile("internal/pkg/testFile/" + file)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	lines, numberOfTables, openTime, closeTime, price, InvalidLine, err := checkDataOfFirstThreeLinesInFile(string(dataOfFile))
	if err != nil {
		log.Printf("%v %v\r\n", err, InvalidLine)
		os.Exit(1)
	}

	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("%v\r\n", openTime.Format("15:04")))

	isClientInTheClub, clientsNumberOfTables, tables, ProfitOfTables :=
		makingOutputFromThirdLine(lines, &builder, numberOfTables, openTime, price)

	clientRemainsInClub := 0
	RemainsClients := []Clients{}
	for client, isInTheClub := range isClientInTheClub {
		if isInTheClub {
			finalProfit := FinalProfit{
				profit:    ProfitOfTables[clientsNumberOfTables[client]].profit + int(math.Ceil(closeTime.Sub(tables[clientsNumberOfTables[client]].timeStart).Hours()))*price,
				wholeTime: ProfitOfTables[clientsNumberOfTables[client]].wholeTime + closeTime.Sub(tables[clientsNumberOfTables[client]].timeStart).Hours(),
			}
			ProfitOfTables[clientsNumberOfTables[client]] = finalProfit

			RemainsClients = append(RemainsClients, Clients{clientRemainsInClub, client})
			clientRemainsInClub++
		}
	}

	sort.Slice(RemainsClients, func(i, j int) bool {
		return RemainsClients[i].name < RemainsClients[j].name
	})

	for nameOfRemainClient := range RemainsClients {
		builder.WriteString(fmt.Sprintf("%v %v %v\r\n", closeTime.Format("15:04"), "11",
			RemainsClients[nameOfRemainClient].name))
	}

	builder.WriteString(fmt.Sprintf("%v\r\n", closeTime.Format("15:04")))

	builder = sortingProfitInfo(ProfitOfTables, &builder)

	return builder.String()
}

func outputTime(input float64) string {
	hours := int(input)
	minutes := int(math.Round((input - float64(hours)) * 60))
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

func sortingProfitInfo(ProfitOfTables map[string]FinalProfit, builder *strings.Builder) strings.Builder {
	var finalProfitArray []FinalProfit
	for key, value := range ProfitOfTables {
		IDOfTable, _ := strconv.Atoi(key)
		finalProfitArray = append(finalProfitArray, FinalProfit{value.wholeTime, value.profit, IDOfTable})
	}

	sort.Slice(finalProfitArray, func(i, j int) bool {
		return finalProfitArray[i].IDOfTable < finalProfitArray[j].IDOfTable
	})

	for k := range finalProfitArray {
		if k == len(finalProfitArray)-1 {
			builder.WriteString(fmt.Sprintf("%v %v %v", finalProfitArray[k].IDOfTable,
				finalProfitArray[k].profit, outputTime(finalProfitArray[k].wholeTime)))
			break
		}

		builder.WriteString(fmt.Sprintf("%v %v %v \r\n", finalProfitArray[k].IDOfTable,
			finalProfitArray[k].profit, outputTime(finalProfitArray[k].wholeTime)))
	}

	return *builder
}

func makingOutputFromThirdLine(lines []string, builder *strings.Builder, numberOfTables int, openTime time.Time, price int) (map[string]bool, map[string]string, map[string]SpotOfTable, map[string]FinalProfit) {
	var (
		isClientInTheClub     map[string]bool
		clientsNumberOfTables map[string]string
		tables                map[string]SpotOfTable
		ProfitOfTables        map[string]FinalProfit

		que []string
	)

	isClientInTheClub = make(map[string]bool)
	tables = make(map[string]SpotOfTable)
	clientsNumberOfTables = make(map[string]string)
	ProfitOfTables = make(map[string]FinalProfit)

	que = make([]string, 0)

	for i := 3; i < len(lines); i++ {
		var (
			timeOfCurrentLine time.Time
			timeOfNextLine    time.Time
		)

		currentLine := strings.Split(lines[i], " ")
		if !checkEvent(currentLine, numberOfTables) {
			log.Printf("%s %s\r\n", BadFormatOfLine, lines[i])
			os.Exit(1)
		}

		timeOfCurrentLine, _ = time.Parse("15:04", currentLine[0])
		if i != len(lines)-1 {
			nextLine := strings.Split(lines[i+1], " ")
			timeOfNextLine, _ = time.Parse("15:04", nextLine[0])
			if !timeOfCurrentLine.Before(timeOfNextLine) {
				log.Printf("%s %s\r\n", BadFormatOfLine, lines[i+1])
				os.Exit(1)
			}
		}

		builder.WriteString(fmt.Sprintf("%v\r\n", lines[i]))

		switch currentLine[1] {
		case "1":
			if timeOfCurrentLine.Before(openTime) {
				builder.WriteString(fmt.Sprintf("%v %v %v\r\n", currentLine[0], "13", "NotOpenYet"))
				continue
			}

			if isClientInTheClub[currentLine[2]] {
				builder.WriteString(fmt.Sprintf("%v %v %v\r\n", currentLine[0], "13", "YouShallNotPass"))
				continue
			}

			isClientInTheClub[currentLine[2]] = true
		case "2":
			if !isClientInTheClub[currentLine[2]] {
				builder.WriteString(fmt.Sprintf("%v %v %v\r\n", currentLine[0], "13", "ClientUnknown"))
				continue
			}

			if tables[currentLine[3]].busy {
				builder.WriteString(fmt.Sprintf("%v %v %v\r\n", currentLine[0], "13", "PlaceIsBusy"))
				continue
			}

			IDTable := clientsNumberOfTables[currentLine[2]]
			if tables[IDTable].busy {
				table := SpotOfTable{
					client:    currentLine[2],
					timeStart: timeOfCurrentLine,
					busy:      false,
				}

				finalProfit := FinalProfit{
					profit:    ProfitOfTables[IDTable].profit + int(math.Ceil(timeOfCurrentLine.Sub(tables[IDTable].timeStart).Hours()))*price,
					wholeTime: ProfitOfTables[IDTable].wholeTime + timeOfCurrentLine.Sub(tables[IDTable].timeStart).Hours(),
				}

				ProfitOfTables[IDTable] = finalProfit
				tables[IDTable] = table
			}

			table := SpotOfTable{
				client:    currentLine[2],
				timeStart: timeOfCurrentLine,
				busy:      true,
			}

			tables[currentLine[3]] = table
			clientsNumberOfTables[currentLine[2]] = currentLine[3]
		case "3":
			if len(tables) == 0 {
				builder.WriteString(fmt.Sprintf("%v %v %v\r\n", currentLine[0], "13", "ICanWaitNoLonger!"))
				continue
			}

			flag := false
			for _, table := range tables {
				if !table.busy {
					flag = true
					break
				}
			}

			if flag {
				builder.WriteString(fmt.Sprintf("%v %v %v\r\n", currentLine[0], "13", "ICanWaitNoLonger!"))
				continue
			}

			if len(que) == numberOfTables {
				builder.WriteString(fmt.Sprintf("%v %v %v\r\n", currentLine[0], "11", currentLine[2]))
				isClientInTheClub[currentLine[2]] = false
				continue
			}

			que = append(que, currentLine[2])
		case "4":
			if !isClientInTheClub[currentLine[2]] {
				builder.WriteString(fmt.Sprintf("%v %v %v\r\n", currentLine[0], "13", "ClientUnknown"))
				continue
			}

			isClientInTheClub[currentLine[2]] = false

			IDTable := clientsNumberOfTables[currentLine[2]]
			finalProfit := FinalProfit{
				profit:    ProfitOfTables[IDTable].profit + int(math.Ceil(timeOfCurrentLine.Sub(tables[IDTable].timeStart).Hours()))*price,
				wholeTime: ProfitOfTables[IDTable].wholeTime + timeOfCurrentLine.Sub(tables[IDTable].timeStart).Hours(),
			}

			ProfitOfTables[IDTable] = finalProfit

			if len(que) > 0 {
				que[len(que)-1] = que[0]
				table := SpotOfTable{
					client:    que[len(que)-1],
					timeStart: timeOfCurrentLine,
					busy:      true,
				}

				tables[IDTable] = table
				builder.WriteString(fmt.Sprintf("%v %v %v %v\r\n", currentLine[0], "12", table.client,
					IDTable))
				clientsNumberOfTables[table.client] = IDTable
				que = que[:len(que)-1]
				continue
			}

			table := SpotOfTable{
				client:    currentLine[2],
				timeStart: timeOfCurrentLine,
				busy:      false,
			}

			tables[IDTable] = table
		}

	}

	return isClientInTheClub, clientsNumberOfTables, tables, ProfitOfTables
}
