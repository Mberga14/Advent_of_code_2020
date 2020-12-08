package day8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type operation struct {
	operand              string
	value                int
	used                 bool
	parentOperationIndex int
}

// ParseInput : parses input of input.txt
func ParseInput(fileName string) []operation {

	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var operations []operation

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, " ")

		operand := splitLine[0]
		value, err := strconv.Atoi(splitLine[1])

		if err != nil {
			panic(err)
		}

		operations = append(operations, operation{operand: operand, value: value, used: false})

	}

	return operations

}

//FindAccValueCycle : returns the value of acc register on cycle detection
func FindAccValueCycle(instructions []operation) (int, bool) {

	fp := 0 // Frame Pointer
	accValue := 0
	cycleDetected := false

	for !cycleDetected {
		if fp == len(instructions) {
			return accValue, false
		}
		instruction := instructions[fp]
		if instruction.used {
			return accValue, true
		}
		switch instruction.operand {
		case "acc":
			instructions[fp].used = !instructions[fp].used
			accValue += instruction.value
			fp++
		case "jmp":
			instructions[fp].used = !instructions[fp].used
			fp += instruction.value
		case "nop":
			fp++
		}
	}

	return accValue, false
}

//FindCorruptInstructionAndFix : finds the corrupt instruction, fixes the instructions and executes the instructions to the end
func FindCorruptInstructionAndFix(instructions []operation) int {
	accValue := 0
	cycleDetected := false

	accValue, cycleDetected = FindAccValueCycle(instructions)
	fp := 0
	for cycleDetected {
		for i := range instructions {
			instructions[i].used = false
			instructions[i].parentOperationIndex = 0
		}
		if instructions[fp].operand == "jmp" {
			instruction := instructions[fp]
			instructions[fp].operand = "nop"
			accValue, cycleDetected = FindAccValueCycle(instructions)
			if cycleDetected {
				instructions[fp] = instruction
				fp++
			} else {
				return accValue
			}
		} else if instructions[fp].operand == "nop" {
			instruction := instructions[fp]
			instructions[fp].operand = "jmp"
			accValue, cycleDetected = FindAccValueCycle(instructions)
			if cycleDetected {
				instructions[fp] = instruction
				fp++
			} else {
				return accValue
			}
		} else {
			fp++
		}

	}
	fmt.Println(instructions)
	return accValue
}
