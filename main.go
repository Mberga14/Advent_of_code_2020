package main

import (
	"fmt"

	day5 "github.com/Mberga14/Advent_of_code_2020/Day5"
	day6 "github.com/Mberga14/Advent_of_code_2020/Day6"
)

func main() {
	fmt.Printf("Solution for 5th day:\n\tPart 1: %d\n\tPart 2: %d\n", day5.GetMaxSeatID(), day5.FindMissingSeatID())
	fmt.Println("Solution for day6 is: ", day6.ParseInput())
}
