package day2

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	inputs := readInput()
	fmt.Println(calculateValidPasswords(inputs))
	fmt.Println(calculateValidPasswords2(inputs))
}

func readInput() []string {
	file, err := os.Open("input.txt")
	checkError(err)

	scanner := bufio.NewScanner(file)
	var inputs []string

	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	return inputs
}

func calculateValidPasswords(passwords []string) int {
	var validPasswords = 0

	re := regexp.MustCompile(" ")
	reNum := regexp.MustCompile("-")

	for i := range passwords {
		split := re.Split(passwords[i], -1)
		// get values
		limits := split[0]
		num1, err := strconv.Atoi(reNum.Split(limits, -1)[0])
		num2, err := strconv.Atoi(reNum.Split(limits, -1)[1])
		checkError(err)
		word := split[1][:len(split[1])-1]

		count := strings.Count(split[2], word)

		if num1 <= count && count <= num2 {
			validPasswords++
		}

	}
	return validPasswords
}

func calculateValidPasswords2(passwords []string) int {
	var validPasswords = 0

	re := regexp.MustCompile(" ")
	reNum := regexp.MustCompile("-")

	for i := range passwords {
		split := re.Split(passwords[i], -1)
		// get values
		limits := split[0]
		pos1, err := strconv.Atoi(reNum.Split(limits, -1)[0])
		pos2, err := strconv.Atoi(reNum.Split(limits, -1)[1])
		checkError(err)
		word := split[1][:len(split[1])-1]

		test := split[2]

		if ((string(test[pos1-1]) == word) || (string(test[pos2-1]) == word)) && !((string(test[pos1-1]) == word) && (string(test[pos2-1]) == word)) {
			validPasswords++
		}

	}
	return validPasswords
}
