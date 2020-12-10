package day10

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

	validResult := []int{
		1,
		4,
		5,
		6,
		7,
		10,
		11,
		12,
		15,
		16,
		19,
	}

	fmt.Println(string(colorYellow), "Function output:", functionOutput)
	fmt.Println(string(colorGreen), "Valid output:   ", validResult)
	fmt.Println(string(colorReset), "")

	assert.Equal(t, validResult[0], functionOutput[0], "The two Arrays should be the same.")

}

func TestCalCulateJoltDifference(t *testing.T) {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"

	testData := "input_test.txt"

	functionOutput := CalCulateJoltDifference(testData)

	validResult := 35

	fmt.Println(string(colorYellow), "Function output:", functionOutput)
	fmt.Println(string(colorGreen), "Valid output:   ", validResult)
	fmt.Println(string(colorReset), "")

	assert.Equal(t, validResult, functionOutput, "The two numbers should be the same.")

}

func TestCalculateAdapterPossibilities(t *testing.T) {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"

	testData := "input_test.txt"

	functionOutput := CalculateAdapterPossibilities(testData)
	validResult := 8

	fmt.Println(string(colorYellow), "Function output:", functionOutput)
	fmt.Println(string(colorGreen), "Valid output:   ", validResult)
	fmt.Println(string(colorReset), "")

	assert.Equal(t, validResult, functionOutput, "The two numbers should be the same.")

}
