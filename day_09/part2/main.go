package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type memory struct {
	numOfBlocks int
	id          int
	index       int
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readInput() []int {
	inputData, err := os.ReadFile("./day_09/input.txt")
	check(err)

	lines := strings.Split(strings.TrimSpace(string(inputData)), "\n")
	if len(lines) > 1 {
		panic(errors.New("input contains more than one line"))
	}

	nums := make([]int, len(lines[0]))
	for i, char := range lines[0] {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			panic(fmt.Errorf("invalid character in input: %v", err))
		}
		nums[i] = num
	}

	return nums
}

func isEven(a int) bool {
	return a%2 == 0
}

func computeResult(diskMap []int) int {

	resultLen := 0
	files := make([]memory, (len(diskMap)+1)/2)
	freeSpaces := make([]memory, (len(diskMap))/2)
	index := 0
	for i := 0; i < len(diskMap); i++ {
		if isEven(i) {
			files[i/2] = memory{numOfBlocks: diskMap[i], id: i / 2, index: index}
			resultLen += diskMap[i]
		} else {
			freeSpaces[i/2] = memory{numOfBlocks: diskMap[i], id: i / 2, index: index}
			resultLen += diskMap[i]
		}
		index += diskMap[i]
	}

	for i := len(files) - 1; i >= 0; i-- {
		for j := 0; j < len(freeSpaces); j++ {
			if files[i].index < freeSpaces[j].index {
				break
			}
			if freeSpaces[j].numOfBlocks >= files[i].numOfBlocks {
				freeSpaces[j].numOfBlocks -= files[i].numOfBlocks
				files[i].index = freeSpaces[j].index
				freeSpaces[j].index += files[i].numOfBlocks
				break
			}
		}
	}

	result := 0
	for _, file := range files {
		for i := file.index; i < file.index+file.numOfBlocks; i++ {
			result += file.id * i
		}
	}

	return result
}

func main() {
	diskMap := readInput()
	result := computeResult(diskMap)
	fmt.Println("checksum:", result)
}
