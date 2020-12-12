package day12

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

	validResult := []instruction{
		{direction: "F", ammount: 10},
		{direction: "N", ammount: 3},
		{direction: "F", ammount: 7},
		{direction: "R", ammount: 90},
		{direction: "F", ammount: 11},
	}

	fmt.Println(string(colorYellow), "Function output:", functionOutput)
	fmt.Println(string(colorGreen), "Valid output:   ", validResult)
	fmt.Println(string(colorReset), "")

	assert.Equal(t, validResult[0], functionOutput[0], "The two Arrays should be the same.")

}

func TestCalculateDistance(t *testing.T) {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"

	testData := "input_test.txt"

	functionOutput := CalculateDistance(testData)

	validResult := 25

	fmt.Println(string(colorYellow), "Function output:", functionOutput)
	fmt.Println(string(colorGreen), "Valid output:   ", validResult)
	fmt.Println(string(colorReset), "")

	assert.Equal(t, validResult, functionOutput, "The two numbers should be the same.")

}

func TestCalculateDistanceWaypoint(t *testing.T) {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"

	testData := "input_test.txt"

	functionOutput := CalculateDistanceWaypoint(testData)
	validResult := 286

	fmt.Println(string(colorYellow), "Function output:", functionOutput)
	fmt.Println(string(colorGreen), "Valid output:   ", validResult)
	fmt.Println(string(colorReset), "")

	assert.Equal(t, validResult, functionOutput, "The two numbers should be the same.")

}
