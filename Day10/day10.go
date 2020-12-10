package day10

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

//ParseInput : parse input
func ParseInput(fileName string) []int {

	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var numbers []int

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())

		if err != nil {
			panic(err)
		}

		numbers = append(numbers, value)

	}
	sort.Ints(numbers) //return sorted numbers
	return numbers
}

//CalCulateJoltDifference : part1 day10
func CalCulateJoltDifference(fileName string) int {

	adapters := ParseInput(fileName)

	joltDifference := []int{0, 0, 0}
	for idx := 0; idx < len(adapters)-1; idx++ {
		joltDifference[(adapters[idx+1]-adapters[idx])-1] = joltDifference[(adapters[idx+1]-adapters[idx])-1] + 1
	}
	return (joltDifference[0] + 1) * (joltDifference[2] + 1)
}

//CalculateAdapterPossibilities : part2 day10
func CalculateAdapterPossibilities(fileName string) int {

	adapters := ParseInput(fileName)
	adapters = append([]int{0}, adapters...)
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	result := calculateAdapterPossibilities(adapters)

	return result

}

func calculateAdapterPossibilities(adapters []int) int {

	dp := make([]int, len(adapters))
	dp[0] = 1
	j := 0
	count := 1
	for idx := 1; idx < len(adapters)-1; idx++ {
		for adapters[idx]-adapters[j] > 3 {
			count -= dp[j]
			j++
		}

		dp[idx] = count
		count += dp[idx]
	}
	return dp[len(dp)-2]
}
