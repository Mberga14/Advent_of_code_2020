package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Mberga14/Advent_of_code_2020/tree/main/Day5/seatparser"
)

func main() {
	seatStrings := parseInput("input.txt")

	result := seatparser.GetMaxSeatID(seatStrings)

	fmt.Println("Highest seat ID is :", result)

	missingID := seatparser.FindMissingSeatID(seatStrings)

	fmt.Println("Your seat ID is :", missingID)

}

func parseInput(fileName string) []string {
	file, err := os.Open(fileName)
	handleError(err)
	var seats []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		seats = append(seats, scanner.Text())
	}

	return seats
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
