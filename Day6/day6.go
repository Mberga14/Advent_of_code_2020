package day6

import (
	"bufio"
	"os"
	"strings"
)

// ParseInput : parses Input for day6
func ParseInput(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	var groups [][]string

	scanner := bufio.NewScanner(file)
	var answers []string
	for scanner.Scan() {
		if scanner.Text() == "" {
			groups = append(groups, answers)
			answers = []string{}
			continue
		}

		answers = append(answers, scanner.Text())

	}

	groups = append(groups, answers)
	return groups
}

// CalculateNumberAnsweredQuestions : calculates ammount of people with same answer in group
func CalculateNumberAnsweredQuestions(answers []string, limitation string) int {

	if len(answers) == 1 {
		return len(answers[0])
	}

	var answerStrings [][]string
	for _, answer := range answers {
		answerString := strings.Split(answer, "")
		answerStrings = append(answerStrings, answerString)
	}

	answeredMap := make(map[string]int)

	for i := range answers {
		for _, letter := range answerStrings[i] {
			if _, isPresent := answeredMap[letter]; !isPresent {
				answeredMap[letter] = answeredMap[letter] + 1
			} else {
				answeredMap[letter] = answeredMap[letter] + 1
			}
		}
	}

	if limitation == "Part2" {
		part2answer := 0
		for _, value := range answeredMap {
			if value == len(answers) {
				part2answer++
			}
		}

		return part2answer
	}

	return len(answeredMap)
}

// SumOfAnsweredQuestions : sum of CalculateNumberAnsweredQuestions for groups
func SumOfAnsweredQuestions(fileName string, limitation string) int {
	groups := ParseInput(fileName)

	counter := 0

	for _, group := range groups {
		counter += CalculateNumberAnsweredQuestions(group, limitation)
	}

	return counter
}
