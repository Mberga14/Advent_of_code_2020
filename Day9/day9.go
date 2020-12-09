package day9

import (
	"bufio"
	"os"
	"strconv"
)

func parseInput(fileName string) []int {

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

	return numbers

}

//EncryptionWeakness : part2 solution
func EncryptionWeakness(fileName string, preambleLength int) int {
	rogueNumber, numbers := GetRougeNumber(fileName, preambleLength)

	low := 0
	high := 0
	sliceCount := sumSlice(numbers[low:high])

	for sliceCount != rogueNumber {
		if sliceCount < rogueNumber { //number is smaller, we should add more numbers to the slice
			high++
			sliceCount = sumSlice(numbers[low:high])
			continue
		}

		if sliceCount > rogueNumber { // number is bigger, the set is wrong
			low++
			high = low
			sliceCount = sumSlice(numbers[low:high])
		}
	}
	lowest, highest := minMax(numbers[low:high])
	return lowest + highest
}

//  https://stackoverflow.com/questions/34259800/is-there-a-built-in-min-function-for-a-slice-of-int-arguments-or-a-variable-numb - thanks Michael Dorner!
func minMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func sumSlice(slice []int) int {
	count := 0
	for _, value := range slice {
		count += value
	}
	return count
}

//GetRougeNumber : Part1 of day9
func GetRougeNumber(fileName string, preambleLength int) (int, []int) {
	var slidingWindow []int
	numbers := parseInput(fileName)

	slidingWindow = numbers[:preambleLength]
	low := 0
	high := preambleLength
	testSubject := numbers[preambleLength]
	for calculateFit(slidingWindow, testSubject) {
		low++
		high++
		slidingWindow = numbers[low:high]
		testSubject = numbers[high]
	}

	return testSubject, numbers[:high]
}

func calculateFit(window []int, testSubject int) bool {
	for idx := range window {
		for jdx := range window {
			if idx == jdx {
				continue
			}
			if window[idx]+window[jdx] == testSubject {
				return true
			}
		}
	}
	return false
}
