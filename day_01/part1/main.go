package main

import (
	"fmt"
	"math"
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


func computeTotalDistance(l1, l2 []int) float64 {
	totalDistance := 0.0
	for i := 0; i < len(l1); i++ {
		totalDistance += math.Abs(float64(l1[i] - l2[i]))
	}
	return totalDistance
}

func main() {
	l1, l2 := readInput()
	fmt.Println(int(computeTotalDistance(l1, l2)))
}
