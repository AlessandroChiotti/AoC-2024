package main

import (
	"fmt"
	"os"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput() string {
	data, err := os.ReadFile("./day_03/input.txt")
	check(err)

	memory := string(data)
	return memory
}

func computeResult(instructions []string) int {
	result := 0
	for _, instruction := range instructions {
		var a, b int
		_, err := fmt.Sscanf(instruction, "mul(%d,%d)", &a, &b)
		check(err)
		result += a * b
	}
	return result
}

func extractValidInstructions(memory string) []string {
	whiteSpacesRegex := regexp.MustCompile(`\s+`)
	noWhiteSpacesMemory := whiteSpacesRegex.ReplaceAllString(memory, "")

	disabledPortionsRegex := regexp.MustCompile(`(?m)don\'t\(\)(.*?)do\(\)|don\'t\(\)(.*?)$`)
	enabledMemoryPortions := disabledPortionsRegex.ReplaceAllString(noWhiteSpacesMemory, "")
	instructionsRegex, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	check(err)
	instructions := instructionsRegex.FindAllString(enabledMemoryPortions, -1)

	return instructions
}

func main() {
	memory := readInput()
	instructions := extractValidInstructions(memory)
	fmt.Println("number of instructions: ", len(instructions))
	result := computeResult(instructions)
	fmt.Println("result: ", result)
}
