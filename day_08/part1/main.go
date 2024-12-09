package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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
	x int
	y int
}

func readInput() matrix {
	inputData, err := os.ReadFile("./day_08/input.txt")
	check(err)

	lines := strings.Split(strings.TrimSpace(string(inputData)), "\n")
	data := make([][]rune, len(lines))

	for i, line := range lines {
		data[i] = make([]rune, 0, len(line))
		for _, char := range line {
			data[i] = append(data[i], char)
		}
	}

	return matrix{
		data: data,
		rows: len(data),
		cols: len(data[0]),
	}
}

func extractAntennasPositions(puzzle matrix) map[rune][]position {
	antennasPositions := make(map[rune][]position)
	for i := 0; i < len(puzzle.data); i++ {
		for j := 0; j < len(puzzle.data[i]); j++ {
			if puzzle.data[i][j] != empty {
				antennasPositions[puzzle.data[i][j]] = append(antennasPositions[puzzle.data[i][j]], position{x: j, y: i})
			}
		}
	}
	return antennasPositions
}

func computeResult(puzzle matrix) int {

	antennasPositions := extractAntennasPositions(puzzle)
	antinodePositions := make(map[position]int)
	for _, antennaPositions := range antennasPositions {
		for i := 0; i < len(antennaPositions)-1; i++ {
			for j := i + 1; j < len(antennaPositions); j++ {
				var antinode1 position
				var antinode2 position
				dx := int(math.Abs(float64(antennaPositions[j].x) - float64(antennaPositions[i].x)))
				dy := antennaPositions[j].y - antennaPositions[i].y
				if antennaPositions[i].x < antennaPositions[j].x {
					antinode1 = position{x: antennaPositions[i].x - dx, y: antennaPositions[i].y - dy}
					antinode2 = position{x: antennaPositions[j].x + dx, y: antennaPositions[j].y + dy}
				} else if antennaPositions[i].x > antennaPositions[j].x {
					antinode1 = position{x: antennaPositions[i].x + dx, y: antennaPositions[i].y - dy}
					antinode2 = position{x: antennaPositions[j].x - dx, y: antennaPositions[j].y + dy}
				} else {
					panic(errors.New("impossible situation"))
				}
				if antinode1.x >= 0 && antinode1.x < puzzle.cols && antinode1.y >= 0 && antinode1.y < puzzle.rows {
					antinodePositions[antinode1] = antinodePositions[antinode1] + 1
				}

				if antinode2.x >= 0 && antinode2.x < puzzle.cols && antinode2.y >= 0 && antinode2.y < puzzle.rows {
					antinodePositions[antinode2] = antinodePositions[antinode2] + 1
				}

			}
		}

	}

	return len(antinodePositions)
}

func main() {
	puzzle := readInput()
	result := computeResult(puzzle)
	fmt.Println("result: ", result)
}
