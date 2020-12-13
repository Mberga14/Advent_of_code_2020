package day13

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

const maxUint = ^uint(0)
const maxInt = int(maxUint >> 1)

// FindEarliestBus : part1
func FindEarliestBus(fileName string) int {

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	timeToDepart, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	busses := scanner.Text()
	bussesSplit := strings.Split(busses, ",")
	sort.Strings(bussesSplit)
	minTime := maxInt
	busNumber := 0
	for _, value := range bussesSplit {
		valueInt, err := strconv.Atoi(value) // read value, since the array is sorted once we get to 'x' char we break
		if err != nil {
			break
		}
		arrivalTimeMod := timeToDepart % valueInt
		arrivalTime := abs(arrivalTimeMod - valueInt)

		if arrivalTime < minTime {
			minTime = arrivalTime
			busNumber = valueInt
		}
	}
	return busNumber * minTime
}

type rule struct {
	value  int
	offSet int
}

//CalculateInterval : part2
func CalculateInterval(fileName string) int {

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	_, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	busses := scanner.Text()
	bussesSplit := strings.Split(busses, ",")

	bussesInt := []int{}
	for _, value := range bussesSplit {
		valueInt, err := strconv.Atoi(value)
		if err != nil {
			bussesInt = append(bussesInt, -1)
			continue
		}
		bussesInt = append(bussesInt, valueInt)
	}

	rules := []rule{}

	cnt := 0

	maxValue := -1
	for idx, value := range bussesInt {
		if value == -1 {
			cnt++
		} else {
			rules = append(rules, rule{value: value, offSet: idx})
			if value > maxValue {
				maxValue = value
			}
			cnt = 0
		}
	}

	position := 0
	increment := 0

	for _, offset := range rules {
		for true {
			position += increment
			if (position+offset.offSet)%offset.value == 0 {
				break
			}
			increment *= offset.value
		}
	}

	return position
}

// Abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// GCD : greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM :  find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
