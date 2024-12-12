package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const BLINKS int = 25

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readInput() []int {
	dataInput, err := os.ReadFile("./day_11/input.txt")
	check(err)

	lines := strings.Split(strings.TrimSpace(string(dataInput)), "\n")
	var data []int
	for _, s := range strings.Split(lines[0], " ") {
		num, _ := strconv.Atoi(s)
		data = append(data, num)
	}

	return data
}

func numOfDigits(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func computeStones(generatedStones []int, blink int) int {
	if blink == BLINKS {
		return len(generatedStones)
	}

	var newStones []int
	for i := 0; i < len(generatedStones); i++ {
		stone := generatedStones[i]
		if stone == 0 {
			newStones = append(newStones, 1)
		} else if numOfDigits(stone)%2 == 0 {
			digits := numOfDigits(stone)
			x := 1
			for j := 0; j < digits/2; j++ {
				x *= 10
			}
			newStones = append(newStones, stone/x)
			newStones = append(newStones, stone%x)
		} else {
			newStones = append(newStones, stone*2024)
		}
	}

	blink += 1
	return computeStones(newStones, blink)
}

func computeResult(stones []int) int {
	result := 0

	for _, stone := range stones {
		generatedStones := make([]int, 1)
		generatedStones[0] = stone
		result += computeStones(generatedStones, 0)
	}

	return result
}

func main() {
	stones := readInput()
	fmt.Println(stones)
	result := computeResult(stones)
	fmt.Println("result:", result)
}
