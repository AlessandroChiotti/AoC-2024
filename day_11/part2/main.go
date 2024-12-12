package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const BLINKS int = 75

type key struct {
	stone int
	blink int
}

var CACHE = make(map[key]int)

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

func computeStones(stone int, blinks int) int {
	if blinks == 0 {
		return 1
	}
	blinks -= 1

	if CACHE[key{stone: stone, blink: blinks}] != 0 {
		return CACHE[key{stone: stone, blink: blinks}]
	}

	if stone == 0 {
		CACHE[key{stone: stone, blink: blinks}] = computeStones(1, blinks)
	} else if numOfDigits(stone)%2 == 0 {
		digits := numOfDigits(stone)
		x := 1
		for j := 0; j < digits/2; j++ {
			x *= 10
		}
		CACHE[key{stone: stone, blink: blinks}] = computeStones(stone/x, blinks) + computeStones(stone%x, blinks)
	} else {
		CACHE[key{stone: stone, blink: blinks}] = computeStones(stone*2024, blinks)
	}
	return CACHE[key{stone: stone, blink: blinks}]
}

func computeResult(stones []int) int {
	result := 0
	for _, stone := range stones {
		result += computeStones(stone, BLINKS)
	}

	return result
}

func main() {
	stones := readInput()
	fmt.Println(stones)
	result := computeResult(stones)
	fmt.Println("result:", result)
}
