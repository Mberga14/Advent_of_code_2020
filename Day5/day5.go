package day5

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	result := GetMaxSeatID()

	fmt.Println("Highest seat ID is :", result)

	missingID := FindMissingSeatID()

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
