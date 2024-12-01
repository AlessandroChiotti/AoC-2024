package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)



func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readInput() ([]int, []int) {
	l1 := []int{}
	l2 := []int{}
	data, err := os.ReadFile("./day_01/input.txt")
	check(err)
	
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		numbers := strings.Fields(line)
		if len(numbers) != 2 {
			continue
		}
		n1, err := strconv.Atoi(numbers[0])
		check(err)
		n2, err := strconv.Atoi(numbers[1])
		check(err)
		l1 = append(l1, n1)
		l2 = append(l2, n2)
   }

	return l1, l2
}

func computeSimilarity(l1 []int,  m map[int]int) int {
	similarity := 0
	for i := 0; i < len(l1); i++ {
		similarity += l1[i] * m[l1[i]]
	}
	return similarity
}

func main() {
	l1, l2 := readInput()
	m := make(map[int]int)
	for i := 0; i < len(l2); i++ {
		m[l2[i]] += 1
	}
	fmt.Println(int(computeSimilarity(l1, m)))
}
