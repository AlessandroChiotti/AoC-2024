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
	right          rune = '>'
	left           rune = '<'
	up             rune = '^'
	down           rune = 'v'
	obstruction    rune = '#'
	newObstruction rune = 'O'
	empty          rune = '.'
)

type matrix struct {
	data [][]rune
	rows int
	cols int
}

type cursorPosition struct {
	x      int
	y      int
	cursor rune
}

type position struct {
	x int
	y int
}

func readInput() (matrix, cursorPosition) {
	inputData, err := os.ReadFile("./day_06/input.txt")
	check(err)

	var startPosition cursorPosition
	lines := strings.Split(strings.TrimSpace(string(inputData)), "\n")
	data := make([][]rune, len(lines))

	for i, line := range lines {
		data[i] = make([]rune, 0, len(line))
		for j, char := range line {
			data[i] = append(data[i], char)
			if char == '^' || char == '>' || char == '<' || char == 'v' {
				startPosition = cursorPosition{x: j, y: i, cursor: char}
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

func getNewCursorPosition(curPos cursorPosition) (cursorPosition, error) {
	cursor := curPos.cursor
	if cursor == right {
		return cursorPosition{cursor: cursor, x: curPos.x + 1, y: curPos.y}, nil
	}
	if cursor == down {
		return cursorPosition{cursor: cursor, x: curPos.x, y: curPos.y + 1}, nil
	}
	if cursor == left {
		return cursorPosition{cursor: cursor, x: curPos.x - 1, y: curPos.y}, nil
	}
	if cursor == up {
		return cursorPosition{cursor: cursor, x: curPos.x, y: curPos.y - 1}, nil
	}
	return cursorPosition{}, fmt.Errorf("cursor not found: %c", cursor)

}

func move(puzzle matrix, curPos cursorPosition, occupiedPositions map[position]rune) (map[position]rune, bool) {
	if curPos.x == 0 || curPos.x == puzzle.cols-1 || curPos.y == 0 || curPos.y == puzzle.rows-1 {
		return occupiedPositions, false
	}

	newCursorPosition, err := getNewCursorPosition(curPos)
	newPosition := position{x: newCursorPosition.x, y: newCursorPosition.y}
	check(err)

	nextCell := puzzle.data[newPosition.y][newPosition.x]
	for nextCell == obstruction || nextCell == newObstruction {
		newCursor, err := getNewDirection(newCursorPosition.cursor)
		check(err)

		newCursorPosition, err = getNewCursorPosition(cursorPosition{x: curPos.x, y: curPos.y, cursor: newCursor})
		check(err)
		newPosition = position{x: newCursorPosition.x, y: newCursorPosition.y}
		nextCell = puzzle.data[newPosition.y][newPosition.x]
	}

	if occupiedPositions[newPosition] == newCursorPosition.cursor {
		return occupiedPositions, true
	}

	occupiedPositions[newPosition] = newCursorPosition.cursor
	var hasCycle bool
	occupiedPositions, hasCycle = move(puzzle, newCursorPosition, occupiedPositions)
	return occupiedPositions, hasCycle
}

func computeResult(puzzle matrix, startPosition cursorPosition) int {
	result := 0
	occupiedPositions := make(map[position]rune)
	occupiedPositions, _ = move(puzzle, startPosition, occupiedPositions)
	fmt.Println(len(occupiedPositions))

	for occupiedPosition := range occupiedPositions {
		if occupiedPosition == (position{x: startPosition.x, y: startPosition.y}) {
			continue
		}

		puzzle.data[occupiedPosition.y][occupiedPosition.x] = newObstruction

		occupiedPositionsWithObstruction := make(map[position]rune)
		_, hasCycle := move(puzzle, startPosition, occupiedPositionsWithObstruction)
		if hasCycle {
			result += 1
		}

		puzzle.data[occupiedPosition.y][occupiedPosition.x] = empty
	}
	return result
}

func main() {
	puzzle, startPosition := readInput()
	result := computeResult(puzzle, startPosition)
	fmt.Println("result: ", result)
}
