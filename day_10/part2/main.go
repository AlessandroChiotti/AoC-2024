package main

import (
	"fmt"
	"os"
	"strings"
)

type position struct {
	x int
	y int
}

type matrix struct {
	data [][]int
	rows int
	cols int
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readInput() matrix {
	dataInput, err := os.ReadFile("./day_10/input.txt")
	check(err)

	lines := strings.Split(strings.TrimSpace(string(dataInput)), "\n")
	data := make([][]int, len(lines))
	for i, line := range lines {
		for _, char := range line {
			num := int(char - '0')
			data[i] = append(data[i], num)
		}
	}
	return matrix{
		data: data,
		rows: len(data),
		cols: len(data[0]),
	}
}

func getTrailheadPositions(topographicMap matrix) []position {

	var positions []position

	for i := 0; i < topographicMap.rows; i++ {
		for j := 0; j < topographicMap.cols; j++ {
			if topographicMap.data[i][j] == 0 {
				positions = append(positions, position{x: j, y: i})
			}
		}
	}

	return positions
}

func computeScore(currPosition position, prevValue int, topographicMap matrix, trailEndPositions []position) []position {
	x := currPosition.x
	y := currPosition.y
	if x < 0 || x > topographicMap.cols-1 || y < 0 || y > topographicMap.rows-1 {
		return trailEndPositions
	}

	if topographicMap.data[y][x] != prevValue+1 {
		return trailEndPositions
	}

	if topographicMap.data[y][x] == 9 {
		trailEndPositions = append(trailEndPositions, currPosition)
	}

	trailEndPositions = computeScore(position{x: x + 1, y: y}, prevValue+1, topographicMap, trailEndPositions)
	trailEndPositions = computeScore(position{x: x - 1, y: y}, prevValue+1, topographicMap, trailEndPositions)
	trailEndPositions = computeScore(position{x: x, y: y + 1}, prevValue+1, topographicMap, trailEndPositions)
	trailEndPositions = computeScore(position{x: x, y: y - 1}, prevValue+1, topographicMap, trailEndPositions)
	return trailEndPositions

}

func computeResult(topographicMap matrix, trailheadPositions []position) int {

	result := 0

	for _, trailheadPosition := range trailheadPositions {
		var trailEndPositions []position
		trailEndPositions = computeScore(trailheadPosition, -1, topographicMap, trailEndPositions)
		result += len(trailEndPositions)
	}

	return result
}

func main() {
	topographicMap := readInput()
	trailheadPositions := getTrailheadPositions(topographicMap)
	result := computeResult(topographicMap, trailheadPositions)
	fmt.Println("result:", result)
}
