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

func readInput() [][]int {
	data, err := os.ReadFile("./day_02/input.txt")
	check(err)

	lines := strings.Split(string(data), "\n")
	var reports = make([][]int, len(lines))
	for i, line := range lines {
		numbers := strings.Fields(line)
		for _, number := range numbers {
			n, err := strconv.Atoi(number)
			check(err)
			reports[i] = append(reports[i], n)
		}
	}

	return reports
}

func countSafeReports(reports [][]int) int {
	safeReports := 0
	for _, report := range reports {
		if len(report) == 0 {
			continue
		}

		asc := report[1] > report[0]
		var j int
		for j = 1; j < len(report); j++ {
			level1 := report[j-1]
			level2 := report[j]
			if (asc != (level2 > level1)) || level2 == level1 {
				break
			}
			if math.Abs(float64(level2-level1)) > 3 {
				break
			}
		}
		if j == len(report) {
			safeReports++
		}
	}
	return safeReports
}

func main() {
	reports := readInput()
	fmt.Println(len(reports))
	fmt.Println(int(countSafeReports(reports)))
}
