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

const (
	right       rune = '>'
	left        rune = '<'
	up          rune = '^'
	down        rune = 'v'
	obstruction rune = '#'
	empty       rune = '.'
)

type matrix struct {
	data [][]rune
	rows int
	cols int
}

type position struct {
	x      int
	y      int
	cursor rune
}

type position2 struct {
	x int
	y int
}

func readInput() (matrix, position) {
	inputData, err := os.ReadFile("./day_06/input.txt")
	check(err)

	var startPosition position
	lines := strings.Split(strings.TrimSpace(string(inputData)), "\n")
	data := make([][]rune, len(lines))

	for i, line := range lines {
		data[i] = make([]rune, 0, len(line))
		for j, char := range line {
			data[i] = append(data[i], char)
			if char == '^' || char == '>' || char == '<' || char == 'v' {
				startPosition = position{x: j, y: i, cursor: char}
				data[i][j] = empty
			}
		}
	}

	return matrix{
		data: data,
		rows: len(data),
		cols: len(data[0]),
	}, startPosition
}

func getNewDirection(cursor rune) (rune, error) {
	if cursor == right {
		return down, nil
	}
	if cursor == down {
		return left, nil
	}
	if cursor == left {
		return up, nil
	}
	if cursor == up {
		return right, nil
	}
	return 0, fmt.Errorf("cursor not found: %c", cursor)

}

func getNewPosition(curPos position) (position, error) {
	cursor := curPos.cursor
	if cursor == right {
		return position{cursor: cursor, x: curPos.x + 1, y: curPos.y}, nil
	}
	if cursor == down {
		return position{cursor: cursor, x: curPos.x, y: curPos.y + 1}, nil
	}
	if cursor == left {
		return position{cursor: cursor, x: curPos.x - 1, y: curPos.y}, nil
	}
	if cursor == up {
		return position{cursor: cursor, x: curPos.x, y: curPos.y - 1}, nil
	}
	return position{}, fmt.Errorf("cursor not found: %c", cursor)

}

func move(puzzle matrix, curPos position, result map[position2]int) map[position2]int {
	if curPos.x == 0 || curPos.x == puzzle.cols-1 || curPos.y == 0 || curPos.y == puzzle.rows-1 {
		return result
	}

	newPosition, err := getNewPosition(curPos)

	check(err)

	nextCell := puzzle.data[newPosition.y][newPosition.x]
	if nextCell == obstruction {
		newCursor, err := getNewDirection(curPos.cursor)
		check(err)

		newPosition, err = getNewPosition(position{x: curPos.x, y: curPos.y, cursor: newCursor})
		check(err)
	}

	result[position2{x: newPosition.x, y: newPosition.y}] += 1
	result = move(puzzle, newPosition, result)
	return result
}

func computeResult(puzzle matrix, startPosition position) map[position2]int {
	result := make(map[position2]int)
	return move(puzzle, startPosition, result)
}

func main() {
	puzzle, startPosition := readInput()
	result := computeResult(puzzle, startPosition)
	fmt.Println("result: ", len(result))
}
