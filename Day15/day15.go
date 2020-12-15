package day15

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type magicNumber struct {
	position      int
	lastOccurance int
}

//Calculate2020thNumberPart1 : part1 day15
func Calculate2020thNumberPart1(fileName string) int {

	file, _ := os.Open(fileName)

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := strings.Split(scanner.Text(), ",")

	var lineInt []int

	numbers := make(map[int][]int)

	for i, value := range line {
		num, _ := strconv.Atoi(value)
		lineInt = append(lineInt, num)
		numbers[num] = append(numbers[num], i+1)
	}

	for i := len(numbers) - 1; i < 30000001; i++ {
		if val, ok := numbers[lineInt[i]]; ok { //number exists
			if len(val) == 1 { //number happend only once
				lineInt = append(lineInt, 0)
				numbers[0] = append(numbers[0], i+2)
				if len(numbers[0]) > 2 {
					numbers[0] = numbers[0][1:]
				}
			} else if len(val) == 2 {
				lineInt = append(lineInt, numbers[lineInt[i]][1]-numbers[lineInt[i]][0])
				numbers[numbers[lineInt[i]][1]-numbers[lineInt[i]][0]] = append(numbers[numbers[lineInt[i]][1]-numbers[lineInt[i]][0]], i+2)
				if len(numbers[numbers[lineInt[i]][1]-numbers[lineInt[i]][0]]) > 2 {
					numbers[numbers[lineInt[i]][1]-numbers[lineInt[i]][0]] = numbers[numbers[lineInt[i]][1]-numbers[lineInt[i]][0]][1:]
				}
			}
		} else {
			numbers[0] = append(numbers[0], i+1)
		}
	}

	return lineInt[29999999]
}
