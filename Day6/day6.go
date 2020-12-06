package day6

import (
	"bufio"
	"os"
)

// ParseInput : parses Input for day6
func ParseInput() [][]string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var groups [][]string

	scanner := bufio.NewScanner(file)
	var answers []string
	for scanner.Scan() {
		if scanner.Text() == "\n" {
			groups = append(groups, answers)
			answers = []string{}
		}

		answers = append(answers, scanner.Text())

	}

	groups = append(groups, answers)
	return groups
}
