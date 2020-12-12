package day12

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

type instruction struct {
	direction string
	ammount   int
}

type location struct {
	x int
	y int
}

//ParseInput : parse input of day12
func ParseInput(fileName string) []instruction {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	instructions := []instruction{}

	for scanner.Scan() {
		line := scanner.Text()

		direction := line[:1]
		ammount, _ := strconv.Atoi(line[1:])
		instructions = append(instructions, instruction{direction: direction, ammount: ammount})
	}

	return instructions

}

//CalculateDistance : day12 part 1
func CalculateDistance(fileName string) int {

	instructions := ParseInput(fileName)
	directionsLeft := []string{"east", "north", "west", "south"}
	directionsRight := []string{"north", "east", "south", "west"}
	direction := "east"
	ferryLocation := location{0, 0}

	for _, instruction := range instructions { //go over all instructions

		switch instruction.direction {
		case "N": // go up
			ferryLocation.y += instruction.ammount
		case "S":
			ferryLocation.y -= instruction.ammount
		case "E":
			ferryLocation.x += instruction.ammount
		case "W":
			ferryLocation.x -= instruction.ammount
		case "L":
			test := instruction.ammount / 90
			currentIndex := SliceIndex(len(directionsLeft), func(i int) bool { return directionsLeft[i] == direction })
			direction = directionsLeft[(test+currentIndex)%4]
		case "R":
			test := instruction.ammount / 90
			currentIndex := SliceIndex(len(directionsRight), func(i int) bool { return directionsRight[i] == direction })
			direction = directionsRight[(test+currentIndex)%4]
		case "F":

			switch direction {
			case "east":
				ferryLocation.x += instruction.ammount
			case "north":
				ferryLocation.y += instruction.ammount
			case "west":
				ferryLocation.x -= instruction.ammount
			case "south":
				ferryLocation.y -= instruction.ammount
			}
		}

	}

	return Abs(Abs(ferryLocation.x) + Abs(ferryLocation.y))
}

//CalculateDistanceWaypoint : day12 part 2
func CalculateDistanceWaypoint(fileName string) int {

	instructions := ParseInput(fileName)
	ferryLocation := location{0, 0}
	waypointLocation := location{10, 1}

	for _, instruction := range instructions { //go over all instructions

		switch instruction.direction {
		case "N": // go up
			waypointLocation.y += instruction.ammount
		case "S":
			waypointLocation.y -= instruction.ammount
		case "E":
			waypointLocation.x += instruction.ammount
		case "W":
			waypointLocation.x -= instruction.ammount
		case "L":
			rad := float64(instruction.ammount) * (math.Pi / 180)

			wlx := waypointLocation.x
			wly := waypointLocation.y

			waypointLocation.x = int(float64(wlx)*math.Cos(rad)) - int(float64(wly)*math.Sin(rad))
			waypointLocation.y = int(float64(wlx)*math.Sin(rad)) + int(float64(wly)*math.Cos(rad))
		case "R":
			rad := (float64(instruction.ammount*-1) * (math.Pi / 180))
			wlx := waypointLocation.x
			wly := waypointLocation.y

			waypointLocation.x = int(float64(wlx)*math.Cos(rad)) - int(float64(wly)*math.Sin(rad))
			waypointLocation.y = int(float64(wlx)*math.Sin(rad)) + int(float64(wly)*math.Cos(rad))
		case "F":

			ferryLocation.x += instruction.ammount * waypointLocation.x
			ferryLocation.y += instruction.ammount * waypointLocation.y
		}

	}

	return Abs(Abs(ferryLocation.x) + Abs(ferryLocation.y))
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// SliceIndex : https://stackoverflow.com/questions/8307478/how-to-find-out-element-position-in-slice
func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}
