package main

import (
	"fmt"

	day10 "github.com/Mberga14/Advent_of_code_2020/Day10"
	day11 "github.com/Mberga14/Advent_of_code_2020/Day11"
	day12 "github.com/Mberga14/Advent_of_code_2020/Day12"
	day13 "github.com/Mberga14/Advent_of_code_2020/Day13"
	day14 "github.com/Mberga14/Advent_of_code_2020/Day14"
	day15 "github.com/Mberga14/Advent_of_code_2020/Day15"
	day5 "github.com/Mberga14/Advent_of_code_2020/Day5"
	day6 "github.com/Mberga14/Advent_of_code_2020/Day6"
	day7 "github.com/Mberga14/Advent_of_code_2020/Day7"
	day8 "github.com/Mberga14/Advent_of_code_2020/Day8"
	day9 "github.com/Mberga14/Advent_of_code_2020/Day9"
)

func main() {
	fmt.Printf("Solution for 5th day:\n\tPart 1: %d\n\tPart 2: %d\n", day5.GetMaxSeatID(), day5.FindMissingSeatID())
	fmt.Println("Solution for day6\n\tPart 1:", day6.SumOfAnsweredQuestions("Day6\\input.txt", "Part1"))
	fmt.Println("\tPart 2:", day6.SumOfAnsweredQuestions("Day6\\input.txt", "Part2"))
	fmt.Println("Solution for day7\n\tPart 1:", day7.GetSumOfBags("shiny gold", "Day7\\input.txt"))
	fmt.Println("\tPart 2:", day7.GetAmmountOfBagsInDesiredBag("shiny gold", "Day7\\input.txt"))
	day8Part1, _ := day8.FindAccValueCycle(day8.ParseInput("Day8\\input.txt"))
	fmt.Println("Solution for day8\n\tPart 1:", day8Part1)
	fmt.Println("\tPart 2:", day8.FindCorruptInstructionAndFix(day8.ParseInput("Day8\\input.txt")))
	day9Part1, _ := day9.GetRougeNumber("Day9\\input_test.txt", 5)
	fmt.Println("Solution for day9\n\tPart 1:", day9Part1)
	fmt.Println("\tPart 2:", day9.EncryptionWeakness("Day9\\input.txt", 25))
	fmt.Println("Solution for day10\n\tPart 1:", day10.CalCulateJoltDifference("Day10\\input_test.txt"))
	fmt.Println("\tPart 2:", day10.CalculateAdapterPossibilities("Day10\\input.txt"))
	fmt.Println("Solution for day11\n\tPart 1:", day11.Arrange("Day11\\input.txt", 1))
	fmt.Println("\tPart 2:", day11.Arrange("Day11\\input.txt", 2))
	fmt.Println("Solution for day12\n\tPart 1:", day12.CalculateDistance("Day12\\input.txt"))
	fmt.Println("\tPart 2:", day12.CalculateDistanceWaypoint("Day12\\input.txt"))
	fmt.Println("Solution for day13\n\tPart 1:", day13.FindEarliestBus("Day13\\input.txt"))
	//fmt.Println("\tPart 2:", day13.CalculateInterval("Day13\\input_test_2.txt"))
	fmt.Println("Solution for day14\n\tPart 1:", day14.CalculateMemoryValues("Day14\\input.txt"))
	fmt.Println("\tPart 2:", day14.CalculateMemoryValuesP2("Day14\\input.txt"))
	fmt.Println("Solution for day15\n\tPart 1:", day15.Calculate2020thNumberPart1("Day15\\input.txt"))
	//fmt.Println("\tPart 2:", day14.CalculateMemoryValuesP2("Day14\\input.txt"))

}
