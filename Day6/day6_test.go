package day6

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSomething : test
func TestParsing(t *testing.T) {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"

	testData := "input_test.txt"

	functionOutput := ParseInput(testData)

	validResult := [][]string{{"qzbw", "qez"}}

	fmt.Println(string(colorYellow), "Function output:", functionOutput)
	fmt.Println(string(colorGreen), "Valid output:   ", validResult)
	fmt.Println(string(colorReset), "")

	assert.Equal(t, validResult[0], functionOutput[0], "The two Arrays should be the same.")

}

func TestCalculateNumberAnsweredQuestions(t *testing.T) {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"

	testData := []string{"qzbw", "qez"}

	functionOutput := CalculateNumberAnsweredQuestions(testData, "Part1")

	validResult := 5

	fmt.Println(string(colorYellow), "Function output:", functionOutput)
	fmt.Println(string(colorGreen), "Valid output:   ", validResult)
	fmt.Println(string(colorReset), "")

	assert.Equal(t, validResult, functionOutput, "The two numbers should be the same.")

}

func TestCalculateNumberAnsweredQuestionsPart2(t *testing.T) {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"

	testData := []string{"a", "a", "a"}

	functionOutput := CalculateNumberAnsweredQuestions(testData, "Part2")

	validResult := 1

	fmt.Println(string(colorYellow), "Function output:", functionOutput)
	fmt.Println(string(colorGreen), "Valid output:   ", validResult)
	fmt.Println(string(colorReset), "")

	assert.Equal(t, validResult, functionOutput, "The two numbers should be the same.")
}

func TestSumOfAnsweredQuestions(t *testing.T) {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"

	testData := "input_test.txt"

	functionOutput := SumOfAnsweredQuestions(testData, "Part2")

	validResult := 2

	fmt.Println(string(colorYellow), "Function output:", functionOutput)
	fmt.Println(string(colorGreen), "Valid output:   ", validResult)
	fmt.Println(string(colorReset), "")

	assert.Equal(t, validResult, functionOutput, "The two numbers should be the same.")
}
