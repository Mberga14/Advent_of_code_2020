package day16

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type ticket struct {
	attributes []int
}

type rule struct {
	attributeName string
	rules         []interval
	used          bool
}

type interval struct {
	low  int
	high int
}

var invalidTicketIndexes []int

//ParseInput : parse input for day16
func ParseInput(fileName string) ([]rule, ticket, []ticket) {
	file, _ := os.Open(fileName)
	invalidTicketIndexes = []int{}

	scanner := bufio.NewScanner(file)

	limitations := []rule{}
	// parse rules
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		line := strings.Split(scanner.Text(), ": ")
		rules := strings.Split(line[1], " or ")
		limits := rule{attributeName: line[0]}
		for _, value := range rules {
			limit := strings.Split(value, "-")
			limitLow, _ := strconv.Atoi(limit[0])
			limitHigh, _ := strconv.Atoi(limit[1])

			limits.rules = append(limits.rules, interval{low: limitLow, high: limitHigh})
		}

		limitations = append(limitations, limits)
	}

	// parse my ticket

	scanner.Scan() // read the "your ticket:" line
	scanner.Scan()

	myTicketString := strings.Split(scanner.Text(), ",")
	attr := []int{}
	for _, value := range myTicketString {
		attribute, _ := strconv.Atoi(value)
		attr = append(attr, attribute)
	}
	myTicket := ticket{attributes: attr}

	//parse others tickets
	scanner.Scan() // read the "nearby tickets:" line
	scanner.Scan()
	tickets := []ticket{}
	for scanner.Scan() {
		otherTicketString := strings.Split(scanner.Text(), ",")
		attr := []int{}
		for _, value := range otherTicketString {
			attribute, _ := strconv.Atoi(value)
			attr = append(attr, attribute)
		}
		tickets = append(tickets, ticket{attributes: attr})
	}

	return limitations, myTicket, tickets
}

//ReturnErrorRate : day16 part1
func ReturnErrorRate(fileName string) int {
	rules, _, tickets := ParseInput(fileName)

	errorRate := 0
	// go over rules for each ticket
	for idx, nearbyTicket := range tickets {

		for _, attribute := range nearbyTicket.attributes {
			isValid := true
			for _, rule := range rules {
				isValid = checkValidity(rule.rules, attribute)
				if isValid {
					break
				}
			}

			if !isValid {
				errorRate += attribute
				invalidTicketIndexes = append(invalidTicketIndexes, idx)
			}
		}
	}

	return errorRate
}

//GuessAttributes : part2
func GuessAttributes(fileName string) int {
	rules, myTicket, tickets := ParseInput(fileName)
	ReturnErrorRate(fileName)

	correctTickets := []ticket{}
	indexHolder := 0
	for i := range tickets {
		if len(invalidTicketIndexes) == 0 {
			correctTickets = tickets
			break
		}
		if i != invalidTicketIndexes[indexHolder] {
			correctTickets = append(correctTickets, tickets[i])
		} else {
			if indexHolder == len(invalidTicketIndexes)-1 {
				break
			}
			indexHolder++
		}
	}

	interestingRules := []rule{}

	for _, rule := range rules {
		if strings.Contains(rule.attributeName, "departure") {
			interestingRules = append(interestingRules, rule)
		}
	}
	//now we have valid tickets
	sum := 1
	for i := range correctTickets[0].attributes { // get attributes
		attr := getTicketValueAtIndex(correctTickets, i)

		for idx, rule := range interestingRules {
			ruleValid := true
			for _, attribute := range attr {
				if !checkValidity(rule.rules, attribute) {
					ruleValid = false
					break
				}
			}

			if ruleValid && !rules[idx].used { // all attributes are valid for this rule
				//fmt.Println(rules[idx].attributeName + " is " + strconv.Itoa(myTicket.attributes[i]))
				sum *= myTicket.attributes[i]
				rules[idx].used = true
				break
			}
		}
	}
	return sum
}

func getTicketValueAtIndex(tickets []ticket, index int) []int {
	values := []int{}
	for _, ticket := range tickets {
		values = append(values, ticket.attributes[index])
	}

	return values
}

func remove(s []ticket, i int) []ticket {
	s[i] = s[len(s)-1]
	// We do not need to put s[i] at the end, as it will be discarded anyway
	return s[:len(s)-1]
}

func checkValidity(intervals []interval, fieldValue int) bool {
	valid := []bool{}
	for _, interval := range intervals {
		if !(fieldValue <= interval.high && fieldValue >= interval.low) {
			valid = append(valid, false)
		} else {
			valid = append(valid, true)
		}
	}
	isAllFalse := true
	for _, value := range valid {
		if value { //value is false
			isAllFalse = false
		}
	}
	return !isAllFalse
}
