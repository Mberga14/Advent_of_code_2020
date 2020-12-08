package day8

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

	validResult := []operation{
		{operand: "nop", value: 0},
		{operand: "acc", value: 1},
		{operand: "jmp", value: 4},
		{operand: "acc", value: 3},
		{operand: "jmp", value: -3},
		{operand: "acc", value: -99},
		{operand: "acc", value: 1},
		{operand: "jmp", value: -4},
		{operand: "acc", value: 6},
	}

	fmt.Println(string(colorYellow), "Function output:", functionOutput)
	fmt.Println(string(colorGreen), "Valid output:   ", validResult)
	fmt.Println(string(colorReset), "")

	assert.Equal(t, validResult[0], functionOutput[0], "The two Arrays should be the same.")

}

func TestFindAccValueCycle(t *testing.T) {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"

	testData := "input_test.txt"

	functionOutput, _ := FindAccValueCycle(ParseInput(testData))

	validResult := 5

	fmt.Println(string(colorYellow), "Function output:", functionOutput)
	fmt.Println(string(colorGreen), "Valid output:   ", validResult)
	fmt.Println(string(colorReset), "")

	assert.Equal(t, validResult, functionOutput, "The two numbers should be the same.")

}

func TestFindCorruptInstructionAndFix(t *testing.T) {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"

	testData := "input_test.txt"

	functionOutput := FindCorruptInstructionAndFix(ParseInput(testData))
	validResult := 8

	fmt.Println(string(colorYellow), "Function output:", functionOutput)
	fmt.Println(string(colorGreen), "Valid output:   ", validResult)
	fmt.Println(string(colorReset), "")

	assert.Equal(t, validResult, functionOutput, "The two numbers should be the same.")

}
