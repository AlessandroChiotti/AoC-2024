package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type file struct {
	numOfBlocks int
	id          int
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

func computeResult(diskMap []int) []int {

	resultLen := 0
	files := make([]file, (len(diskMap)+1)/2)
	for i := 0; i < len(diskMap); i += 2 {
		files[i/2] = file{numOfBlocks: diskMap[i], id: i / 2}
		resultLen += diskMap[i]
	}

	var result []int
	j := len(files) - 1
	l := 0
	for i := 0; l <= j; i++ { // TODO: fix exit condition
		if isEven(i) {
			for k := 0; k < diskMap[i]; k++ {
				result = append(result, files[l].id)
				files[l].numOfBlocks--
				if files[l].numOfBlocks == 0 {
					l++
					break
				}
			}
		} else {
			for k := 0; k < diskMap[i]; k++ {
				result = append(result, files[j].id)
				files[j].numOfBlocks--
				if files[j].numOfBlocks == 0 {
					j--
				}
			}
		}
	}
	return result
}

func main() {
	diskMap := readInput()
	result := computeResult(diskMap)
	checksum := 0
	for i := 0; i < len(result); i++ {
		checksum += i * result[i]
	}
	fmt.Println("checksum:", checksum)
}
