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
			if puzzle[i][j] == 'X' {
				if j >= 3 && puzzle[i][j-1] == 'M' && puzzle[i][j-2] == 'A' && puzzle[i][j-3] == 'S' {
					result++
				}
				if j < len(puzzle[i])-3 && puzzle[i][j+1] == 'M' && puzzle[i][j+2] == 'A' && puzzle[i][j+3] == 'S' {
					result++
				}
				if i >= 3 && puzzle[i-1][j] == 'M' && puzzle[i-2][j] == 'A' && puzzle[i-3][j] == 'S' {
					result++
				}
				if i+3 < len(puzzle) && puzzle[i+1][j] == 'M' && puzzle[i+2][j] == 'A' && puzzle[i+3][j] == 'S' {
					result++
				}
				if i >= 3 && j >= 3 && puzzle[i-1][j-1] == 'M' && puzzle[i-2][j-2] == 'A' && puzzle[i-3][j-3] == 'S' {
					result++
				}
				if j < len(puzzle[i])-3 && i < len(puzzle)-3 && puzzle[i+1][j+1] == 'M' && puzzle[i+2][j+2] == 'A' && puzzle[i+3][j+3] == 'S' {
					result++
				}
				if i >= 3 && j+3 < len(puzzle[i]) && puzzle[i-1][j+1] == 'M' && puzzle[i-2][j+2] == 'A' && puzzle[i-3][j+3] == 'S' {
					result++
				}
				if j >= 3 && i+3 < len(puzzle) && puzzle[i+1][j-1] == 'M' && puzzle[i+2][j-2] == 'A' && puzzle[i+3][j-3] == 'S' {
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
