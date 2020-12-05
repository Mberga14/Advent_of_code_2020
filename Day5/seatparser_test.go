package main

import (
	"fmt"
	"testing"

	"github.com/Mberga14/Advent_of_code_2020/tree/main/Day5/seatparser"

	"github.com/stretchr/testify/assert"
)

// TestSomething : test
func TestGetSeatID(t *testing.T) {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"

	testData := "FBFBBFFRLR"

	functionOutput := seatparser.GetSeatID(testData)

	validResult := []int{44, 5}

	fmt.Println(string(colorYellow), "Function output:", functionOutput)
	fmt.Println(string(colorGreen), "Valid output:   ", validResult)
	fmt.Println(string(colorReset), "")

	assert.Equal(t, validResult, functionOutput, "The two Arrays should be the same.")

}
