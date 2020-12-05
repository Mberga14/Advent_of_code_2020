package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(dat)
	var numbers []int

	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		check(err)
		numbers = append(numbers, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i, number := range numbers {
		for _, num1 := range numbers[i:] {
			if number+num1 == 2020 {
				fmt.Println(num1 * number)
				break
			}
		}
	}

	for i, num1 := range numbers {
		for j, num2 := range numbers[i:] {
			for _, num3 := range numbers[j:] {
				if num1+num2+num3 == 2020 {
					fmt.Println(num1 * num2 * num3)
					break
				}
			}
		}
	}
}
