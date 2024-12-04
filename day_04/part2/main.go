package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput() [][]rune {
	data, err := os.ReadFile("./day_04/input.txt")
	check(err)

	lines := strings.Split(string(data), "\n")
	puzzle := make([][]rune, len(lines)-1)
	for i, line := range lines {
		for _, char := range line {
			puzzle[i] = append(puzzle[i], char)
		}
	}
	return puzzle
}

func computeResult(puzzle [][]rune) int {
	result := 0
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[i]); j++ {
			if puzzle[i][j] == 'A' {
				if !(i >= 1 && j >= 1 && i < len(puzzle)-1 && j < len(puzzle[i])-1) {
					continue
				}
				if (puzzle[i-1][j+1] == 'M' && puzzle[i+1][j-1] == 'S' || puzzle[i+1][j-1] == 'M' && puzzle[i-1][j+1] == 'S') && (puzzle[i-1][j-1] == 'M' && puzzle[i+1][j+1] == 'S' || puzzle[i+1][j+1] == 'M' && puzzle[i-1][j-1] == 'S') {
					result++
				}
			}
		}

	}
	return result
}

func main() {
	puzzle := readInput()
	result := computeResult(puzzle)
	fmt.Println("result: ", result)
}
