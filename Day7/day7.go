package day7

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Bag : has a color and carrying capacity map in form of map[string]int
type Bag struct {
	Color   string
	CanHold map[string]int
}

// ParseInput : will parse input in the type of <color> bags contain <number> <color> bag, <number> <color> bags.
// can be parsed with contain
func ParseInput(fileName string) []Bag {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var bags []Bag
	for scanner.Scan() {
		line := scanner.Text()

		splitByContain := strings.Split(line, " contain ")
		bagColor := splitByContain[0][:len(splitByContain[0])-5]

		splitByComma := strings.Split(splitByContain[1], ", ")

		canHold := make(map[string]int)
		for _, value := range splitByComma {
			splitBySpace := strings.Split(value, " ")
			carryCapacity := 0

			if !strings.Contains(splitBySpace[0], "no") {
				carryCapacity, err = strconv.Atoi(splitBySpace[0])
			}

			capacityColor := strings.Join(splitBySpace[1:len(splitBySpace)-1], " ")

			canHold[capacityColor] = carryCapacity
		}

		bags = append(bags, Bag{Color: bagColor, CanHold: canHold})

	}

	return bags
}

func canHoldBag(currentBag Bag, desiredBag string, bags []Bag) bool {
	for bag, capacity := range currentBag.CanHold {
		if capacity > 0 && bag == desiredBag {
			return true
		}
		b := findBag(bag, bags)
		if capacity > 0 && canHoldBag(b, desiredBag, bags) {
			return true
		}
	}

	return false
}

func findBag(color string, bags []Bag) Bag {

	for _, bag := range bags {
		if bag.Color == color {
			return bag
		}
	}
	return Bag{}
}

// GetSumOfBags : gets sum of bags that can hhold your bag
func GetSumOfBags(desiredBag string, fileName string) int {

	bags := ParseInput(fileName)
	count := 0
	for _, bag := range bags {
		if canHoldBag(bag, desiredBag, bags) {
			count++
		}
	}

	return count
}

func calculateContainingBags(desiredBag string, bags []Bag) int {
	b := findBag(desiredBag, bags)

	canHold := true

	for _, capacity := range b.CanHold {
		if capacity == 0 {
			canHold = false
		}
	}

	if !canHold {
		return 0
	}
	total := 0
	for bag, value := range b.CanHold {
		total += value + (value * calculateContainingBags(bag, bags))
	}

	return total

}

//GetAmmountOfBagsInDesiredBag : part2
func GetAmmountOfBagsInDesiredBag(desiredBag string, fileName string) int {

	bags := ParseInput(fileName)
	count := calculateContainingBags(desiredBag, bags)

	return count

}
