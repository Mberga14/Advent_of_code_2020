package seatparser

import (
	"fmt"
	"math"
)

// GetSeatID : get ID of single seat
func GetSeatID(seatString string) []int {
	row := seatString[:7]
	column := seatString[7:len(seatString)]

	fmt.Println(row, " ", column)
	return []int{0, 0} //this will return the row and column
}

// GetMaxSeatID : get max ID of seats
func GetMaxSeatID(seats []string) int {
	maxSeatID := math.MinInt32
	for _, seat := range seats {
		rowCol := GetSeatID(seat)
		calculatedSeatID := CalculateSeatID(rowCol[0], rowCol[1])
		if calculatedSeatID > maxSeatID {
			maxSeatID = calculatedSeatID
		}
	}
	return maxSeatID
}

// CalculateSeatID : calculate the ID of the given seat
func CalculateSeatID(row int, column int) int {
	return (row * 8) + column
}
