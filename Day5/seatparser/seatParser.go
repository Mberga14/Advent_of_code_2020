package seatparser

import (
	"math"
	"sort"
)

// GetSeatID : get ID of single seat
func GetSeatID(seatString string) []int {
	row := seatString[:7]
	column := seatString[7:len(seatString)]

	rowRune := []rune(row)

	columnRune := []rune(column)

	rowNumber := bynarySearch(127, 0, rowRune)
	columnNumber := bynarySearch(7, 0, columnRune)

	return []int{rowNumber, columnNumber} //this will return the row and column
}

func bynarySearch(top int, bottom int, instructions []rune) int {
	//rowValue := 0
	//calculate row
	middle := 0
	for _, rune := range instructions {
		middle = (1 + (top+bottom)/2)
		if rune == 'F' || rune == 'L' { //take lower half
			top = middle - 1

		}

		if rune == 'B' || rune == 'R' { //take upper half
			bottom = middle
		}
	}

	return top
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

// FindMissingSeatID : finds the missing seat IDs
func FindMissingSeatID(seats []string) int {
	missingSeatID := 0
	var IDs []int
	for _, seat := range seats {
		rowCol := GetSeatID(seat)
		calculatedSeatID := CalculateSeatID(rowCol[0], rowCol[1])
		IDs = append(IDs, calculatedSeatID)
	}

	sort.Ints(IDs)

	for i := 1; i < len(IDs); i++ {
		if IDs[i]-IDs[i-1] != 1 {
			missingSeatID = IDs[i] - 1
		}
	}
	return missingSeatID
}

// CalculateSeatID : calculate the ID of the given seat
func CalculateSeatID(row int, column int) int {
	return (row * 8) + column
}
