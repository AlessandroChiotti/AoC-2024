package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	empty    rune = '.'
	antinode rune = '#'
)

type matrix struct {
	data [][]rune
	rows int
	cols int
}

type position struct {
	x, y int
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readInput() matrix {
	inputData, err := os.ReadFile("./day_08/input.txt")
	check(err)

	lines := strings.Split(strings.TrimSpace(string(inputData)), "\n")
	data := make([][]rune, len(lines))

	for i, line := range lines {
		data[i] = []rune(line)
	}

	return matrix{
		data: data,
		rows: len(data),
		cols: len(data[0]),
	}
}

func extractAntennasPositions(puzzle matrix) map[rune][]position {
	antennas := make(map[rune][]position)

	for y, row := range puzzle.data {
		for x, char := range row {
			if char != empty {
				antennas[char] = append(antennas[char], position{x: x, y: y})
			}
		}
	}
	return antennas
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return abs(a)
}

func addAntinodes(antinodes map[position]int, start, delta position, rows, cols int) {
	current := position{x: start.x + delta.x, y: start.y + delta.y}

	for current.x >= 0 && current.x < cols && current.y >= 0 && current.y < rows {
		antinodes[current]++
		current.x += delta.x
		current.y += delta.y
	}
}

func computeResult(puzzle matrix) int {
	antennas := extractAntennasPositions(puzzle)
	antinodes := make(map[position]int)

	for _, antennaPositions := range antennas {
		for i := 0; i < len(antennaPositions)-1; i++ {
			for j := i + 1; j < len(antennaPositions); j++ {
				p1, p2 := antennaPositions[i], antennaPositions[j]

				dx, dy := p2.x-p1.x, p2.y-p1.y
				factor := gcd(dx, dy)
				dx /= factor
				dy /= factor

				addAntinodes(antinodes, p1, position{x: dx, y: dy}, puzzle.rows, puzzle.cols)
				addAntinodes(antinodes, p2, position{x: -dx, y: -dy}, puzzle.rows, puzzle.cols)

			}
		}
	}

	return len(antinodes)
}

func main() {
	puzzle := readInput()
	result := computeResult(puzzle)
	fmt.Println("result:", result)
}
