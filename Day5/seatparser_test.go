package day5

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

	testData := "FBFBBFFRLR"

	functionOutput := GetSeatID(testData)

	validResult := []int{44, 5}

	fmt.Println(string(colorYellow), "Function output:", functionOutput)
	fmt.Println(string(colorGreen), "Valid output:   ", validResult)
	fmt.Println(string(colorReset), "")

	assert.Equal(t, validResult, functionOutput, "The two Arrays should be the same.")

}
