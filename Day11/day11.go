package day11

import (
	"bufio"
	"os"
)

// ParseInput : parse the input for day11
func ParseInput(fileName string) [][]rune {

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	lines := [][]rune{}
	for scanner.Scan() {
		line := ([]rune)(scanner.Text())

		// 76 = L, 46 = . , 35 = #
		lines = append(lines, line)
	}
	return lines
}

// Arrange : part 1 day 11
func Arrange(fileName string, part int) int {
	seats := ParseInput(fileName)

	seats, isSame := arrageSeats(seats, part)
	count := 0

	for !isSame {
		seats, isSame = arrageSeats(seats, part)
	}

	for idx := range seats {
		for jdx := range seats[idx] {
			if seats[idx][jdx] == '#' {
				count++
			}

		}
	}

	return count

}

func arrageSeats(seats [][]rune, part int) ([][]rune, bool) {

	seatsReworked := make([][]rune, len(seats))

	for idx := range seats {
		seatsReworked[idx] = make([]rune, len(seats[idx]))
		for jdx := range seats[idx] {
			seatsReworked[idx][jdx] = seats[idx][jdx]
		}
	}

	isSame := true
	for idx := range seats {
		for jdx := range seats[idx] {
			occupied := 0
			if part == 1 {
				occupied = checkOccupied(seats, idx, jdx)
			} else {
				occupied = checkOccupiedPart2(seats, idx, jdx)
			}

			if seats[idx][jdx] == 'L' && occupied == 0 {
				seatsReworked[idx][jdx] = '#'
				isSame = false
			}

			if part == 1 {
				if seats[idx][jdx] == '#' && occupied > 3 {
					seatsReworked[idx][jdx] = 'L'
					isSame = false
				}
			} else {
				if seats[idx][jdx] == '#' && occupied > 4 {
					seatsReworked[idx][jdx] = 'L'
					isSame = false
				}
			}
		}
	}

	return seatsReworked, isSame
}

func checkOccupied(seats [][]rune, i int, j int) int {
	counter := 0
	if checkValidity(i-1, j, len(seats), len(seats[0])) && seats[i-1][j] == '#' { //check up
		counter++
	}

	if checkValidity(i-1, j+1, len(seats), len(seats[0])) && seats[i-1][j+1] == '#' { //chech up right
		counter++
	}

	if checkValidity(i, j+1, len(seats), len(seats[0])) && seats[i][j+1] == '#' { //check right
		counter++
	}

	if checkValidity(i+1, j+1, len(seats), len(seats[0])) && seats[i+1][j+1] == '#' { // check down right
		counter++
	}

	if checkValidity(i+1, j, len(seats), len(seats[0])) && seats[i+1][j] == '#' { //check down
		counter++
	}

	if checkValidity(i+1, j-1, len(seats), len(seats[0])) && seats[i+1][j-1] == '#' { // check down left
		counter++
	}

	if checkValidity(i, j-1, len(seats), len(seats[0])) && seats[i][j-1] == '#' { // check left
		counter++
	}

	if checkValidity(i-1, j-1, len(seats), len(seats[0])) && seats[i-1][j-1] == '#' { // check up left
		counter++
	}

	return counter
}

func checkOccupiedPart2(seats [][]rune, i int, j int) int {
	counter := 0

	l := len(seats)
	w := len(seats[0])

	iTemp := i
	jTemp := j

	seatFound := false

	//go up
	for !seatFound {
		if !checkValidity(iTemp-1, jTemp, l, w) {
			break
		}
		if checkValidity(iTemp-1, jTemp, l, w) && (seats[iTemp-1][jTemp] == '#' || seats[iTemp-1][jTemp] == 'L') {
			if seats[iTemp-1][jTemp] == '#' {
				counter++
			}
			seatFound = !seatFound
		} else {
			iTemp--
		}
	}
	//reset temp values
	iTemp = i
	jTemp = j
	seatFound = false

	//go up and right
	for !seatFound {
		if !checkValidity(iTemp-1, jTemp+1, l, w) {
			break
		}
		if checkValidity(iTemp-1, jTemp+1, l, w) && (seats[iTemp-1][jTemp+1] == '#' || seats[iTemp-1][jTemp+1] == 'L') {
			if seats[iTemp-1][jTemp+1] == '#' {
				counter++
			}
			seatFound = !seatFound
		} else {
			iTemp--
			jTemp++
		}
	}

	//reset temp values
	iTemp = i
	jTemp = j
	seatFound = false

	//go right
	for !seatFound {
		if !checkValidity(iTemp, jTemp+1, l, w) {
			break
		}
		if checkValidity(iTemp, jTemp+1, l, w) && (seats[iTemp][jTemp+1] == '#' || seats[iTemp][jTemp+1] == 'L') {
			if seats[iTemp][jTemp+1] == '#' {
				counter++
			}
			seatFound = !seatFound
		} else {
			jTemp++
		}
	}

	//reset temp values
	iTemp = i
	jTemp = j
	seatFound = false

	//go right down
	for !seatFound {
		if !checkValidity(iTemp+1, jTemp+1, l, w) {
			break
		}
		if checkValidity(iTemp+1, jTemp+1, l, w) && (seats[iTemp+1][jTemp+1] == '#' || seats[iTemp+1][jTemp+1] == 'L') {
			if seats[iTemp+1][jTemp+1] == '#' {
				counter++
			}
			seatFound = !seatFound
		} else {
			iTemp++
			jTemp++
		}
	}

	//reset temp values
	iTemp = i
	jTemp = j
	seatFound = false

	//go down
	for !seatFound {
		if !checkValidity(iTemp+1, jTemp, l, w) {
			break
		}
		if checkValidity(iTemp+1, jTemp, l, w) && (seats[iTemp+1][jTemp] == '#' || seats[iTemp+1][jTemp] == 'L') {
			if seats[iTemp+1][jTemp] == '#' {
				counter++
			}
			seatFound = !seatFound
		} else {
			iTemp++
		}
	}

	//reset temp values
	iTemp = i
	jTemp = j
	seatFound = false

	//go down left
	for !seatFound {
		if !checkValidity(iTemp+1, jTemp-1, l, w) {
			break
		}
		if checkValidity(iTemp+1, jTemp-1, l, w) && (seats[iTemp+1][jTemp-1] == '#' || seats[iTemp+1][jTemp-1] == 'L') {
			if seats[iTemp+1][jTemp-1] == '#' {
				counter++
			}
			seatFound = !seatFound
		} else {
			iTemp++
			jTemp--
		}
	}

	//reset temp values
	iTemp = i
	jTemp = j
	seatFound = false

	//go left
	for !seatFound {
		if !checkValidity(iTemp, jTemp-1, l, w) {
			break
		}
		if checkValidity(iTemp, jTemp-1, l, w) && (seats[iTemp][jTemp-1] == '#' || seats[iTemp][jTemp-1] == 'L') {
			if seats[iTemp][jTemp-1] == '#' {
				counter++
			}
			seatFound = !seatFound
		} else {
			jTemp--
		}
	}

	//reset temp values
	iTemp = i
	jTemp = j
	seatFound = false

	//go up left
	for !seatFound {
		if !checkValidity(iTemp-1, jTemp-1, l, w) {
			break
		}
		if checkValidity(iTemp-1, jTemp-1, l, w) && (seats[iTemp-1][jTemp-1] == '#' || seats[iTemp-1][jTemp-1] == 'L') {
			if seats[iTemp-1][jTemp-1] == '#' {
				counter++
			}
			seatFound = !seatFound
		} else {
			jTemp--
			iTemp--
		}
	}

	return counter

}

func checkValidity(i int, j int, length int, width int) bool {
	return (i >= 0 && i < length) && (j >= 0 && j < width)
}
