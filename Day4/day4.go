package day4

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data := readInput()

	result := checkValidity(data)

	fmt.Println(result)
}

func checkValidity(data []string) []int {

	m := map[string]bool{
		// key + neccesary or not
		"byr": true,
		"iyr": true,
		"eyr": true,
		"hgt": true,
		"hcl": true,
		"ecl": true,
		"pid": true,
		"cid": false,
	}
	validDocuments := 0
	validDocumentsStrong := 0
	for _, document := range data {
		isValid := true
		var contained map[string]bool
		contained = make(map[string]bool)
		for key := range m {
			if strings.Contains(document, key) {
				contained[key] = true
			} else {
				contained[key] = false
			}
		}
		for key, value := range contained {
			if !value && m[key] {
				isValid = false
			}
		}

		if isValid {

			var boolarray []bool

			parseSpaces := strings.Split(document, " ")

			for _, parseSpace := range parseSpaces {
				parsedColon := strings.Split(parseSpace, ":")

				if parsedColon[0] == "" {
					continue
				}
				if parsedColon[0] == "cid" {
					continue
				}

				validity := checkValidityStrong(parsedColon[1], parsedColon[0])

				boolarray = append(boolarray, validity)
			}
			isValid := true
			for _, value := range boolarray {
				if !value {
					isValid = false
				}
			}

			if isValid {
				validDocumentsStrong++
			}
			validDocuments++
		}
	}

	return []int{validDocuments, validDocumentsStrong}
}

func checkValidityStrong(passportValue string, limitation string) bool {

	if limitation == "byr" {
		value, _ := strconv.Atoi(passportValue)
		return compareNumber(value, 1920, 2002)
	} else if limitation == "iyr" {
		value, _ := strconv.Atoi(passportValue)
		return compareNumber(value, 2010, 2020)
	} else if limitation == "eyr" {
		value, _ := strconv.Atoi(passportValue)
		return compareNumber(value, 2020, 2030)
	} else if limitation == "hgt" {
		return getHeight(passportValue)
	} else if limitation == "hcl" {
		return getHairColor(passportValue)
	} else if limitation == "ecl" {
		return getEyeColor(passportValue)
	} else if limitation == "pid" {
		return getPID(passportValue)
	}
	return false
}

func compareNumber(num int, low int, high int) bool {
	if low <= num && num <= high {
		return true
	}
	return false
}

func getHeight(passportValue string) bool {
	number := passportValue[:len(passportValue)-2]
	metric := passportValue[len(passportValue)-2:]
	value, _ := strconv.Atoi(number)

	if metric == "cm" {
		return compareNumber(value, 150, 193)
	}
	return compareNumber(value, 59, 76)
}

func getHairColor(passportValue string) bool {
	re := regexp.MustCompile("^#[0-9a-f]{6}$")

	if re.Match([]byte(passportValue)) {
		return true
	}
	return false
}

func getEyeColor(passportValue string) bool {
	allowedNames := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

	for _, allowedName := range allowedNames {
		if strings.Contains(passportValue, allowedName) {
			return true
		}
	}
	return false
}

func getPID(passportValue string) bool {
	re := regexp.MustCompile("^[0-9]{9}$")

	if re.Match([]byte(passportValue)) {
		return true
	}
	return false
}
func readInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var information []string
	var data string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			information = append(information, data)
			data = ""
		} else {
			data += line + " "
		}
	}

	return information
}
