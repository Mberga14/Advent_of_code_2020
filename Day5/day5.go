package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seatStrings := parseInput("input.txt")

	for _, seat := range seatStrings {
		fmt.Println(seat)
		seatparser.getID
	}

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
