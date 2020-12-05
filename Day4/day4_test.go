package day4

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSomething : test
func TestGetSeatID(t *testing.T) {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"

	testData := "160cm"

	functionOutput := getHeight(testData)

	validResult := true

	fmt.Println(string(colorYellow), "Function output:", functionOutput)
	fmt.Println(string(colorGreen), "Valid output:   ", validResult)
	fmt.Println(string(colorReset), "")

	assert.Equal(t, validResult, functionOutput, "The two booleans should be the same.")

}
