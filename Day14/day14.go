package day14

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

var perms []int

// CalculateMemoryValues : part1 day14
func CalculateMemoryValues(fileName string) int {

	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	var memory = make(map[int]int)

	scanner := bufio.NewScanner(file)
	mask := []string{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " = ")
		if line[0] == "mask" { //is da mask
			maskString := strings.Split(line[1], "")
			mask = maskString
		} else { // is da mem
			value := line[0][4 : len(line[0])-1]
			valueInt, _ := strconv.Atoi(value)
			valueMem, _ := strconv.Atoi(line[1])
			valueBin := strconv.FormatInt(int64(valueMem), 2)

			valueBinSplit := strings.Split(valueBin, "")
			append1 := 36 - len(valueBinSplit)
			for i := 0; i < append1; i++ {
				valueBinSplit = append([]string{"0"}, valueBinSplit...)
			}
			memory[valueInt] = applyBitMask(valueBinSplit, mask)
		}
	}

	cnt := 0

	for _, value := range memory {
		cnt += value
	}
	return cnt
}

// CalculateMemoryValuesP2 : part2 day14
func CalculateMemoryValuesP2(fileName string) int {

	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	var memory = make(map[int]int)

	scanner := bufio.NewScanner(file)
	mask := []string{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " = ")
		if line[0] == "mask" { //is da mask
			maskString := strings.Split(line[1], "")
			mask = maskString
		} else { // is da mem
			value := line[0][4 : len(line[0])-1]
			valueInt, _ := strconv.Atoi(value)
			valueMem, _ := strconv.Atoi(line[1])
			valueBin := strconv.FormatInt(int64(valueInt), 2)

			valueBinSplit := strings.Split(valueBin, "")
			append1 := 36 - len(valueBinSplit)
			for i := 0; i < append1; i++ {
				valueBinSplit = append([]string{"0"}, valueBinSplit...)
			}

			maskFloating, _ := applyBitMaskP2(valueBinSplit, mask)
			maskReworked := make([]string, len(maskFloating))

			for idx := range maskFloating {
				maskReworked[idx] = maskFloating[idx]
			}

			generateMemLocations(maskFloating, maskReworked, 0)

			for _, perm := range perms {
				memory[perm] = valueMem
			}
			perms = []int{}
		}
	}

	cnt := 0

	for _, value := range memory {
		cnt += value
	}
	return cnt
}

func generateMemLocations(maskFloating []string, fixedLocation []string, locationOfBit int) {
	if locationOfBit == len(maskFloating) {
		perms = append(perms, convertFromBit(fixedLocation))
		return
	}
	if maskFloating[locationOfBit] == "X" {
		fixedLocation[locationOfBit] = "1"
		generateMemLocations(maskFloating, fixedLocation, locationOfBit+1)
		fixedLocation[locationOfBit] = "0"
		generateMemLocations(maskFloating, fixedLocation, locationOfBit+1)
	} else {
		generateMemLocations(maskFloating, fixedLocation, locationOfBit+1)
	}
}

func applyBitMaskP2(value []string, mask []string) ([]string, int) {
	newValue := []string{}
	cnt := float64(0)
	for idx := range mask {
		if mask[idx] == "1" {
			newValue = append(newValue, mask[idx])
		} else if mask[idx] == "X" {
			newValue = append(newValue, mask[idx])
			cnt++
		} else {
			newValue = append(newValue, value[idx])
		}
	}
	cnt = math.Pow(2, cnt)
	return newValue, int(cnt)
}

func applyBitMask(value []string, mask []string) int {
	newValue := []string{}
	for idx := range mask {
		if mask[idx] == "X" {
			newValue = append(newValue, value[idx])
		} else {
			newValue = append(newValue, mask[idx])
		}
	}

	return convertFromBit(newValue)
}

func convertFromBit(value []string) int {
	count := float64(0)
	lenValue := len(value) - 1

	for i := lenValue; i >= 0; i-- {
		//valueFloat, _ := strconv.ParseFloat(value[i], 64)
		//valueIdx := float64(valueFloat)
		if value[i] == "1" {
			count += math.Pow(2, float64(len(value)-i-1))
		}
	}
	return int(count)
}
