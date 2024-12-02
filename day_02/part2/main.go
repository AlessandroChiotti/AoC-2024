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



func readInput() ([][]int) {
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

func testSafety(report []int) bool {
	var asc bool = report[1] > report[0]
	var i int
	for i = 1; i < len(report); i++ {
		level1 := report[i-1]
		level2 := report[i]

		if (asc != (level2 > level1)) || level2 == level1 {
			return false
		}
		if math.Abs(float64(level2 - level1)) > 3 {
			return false
		}
	}
	return true
}

func RemoveIndex(s []int, index int) []int {
    ret := make([]int, 0)
    ret = append(ret, s[:index]...)
    return append(ret, s[index+1:]...)
}


func countSafeReports(reports [][]int) int {
	safeReports := 0
	for _, report := range reports {
		if len(report) == 0 {
			continue
		}

		isReportSafe := testSafety(report)
		if(isReportSafe) {
			safeReports++
			continue
		}

		for i:= 0 ; i < len(report); i++ {
			partialReport := RemoveIndex(report, i)
			if(testSafety(partialReport)) {
				safeReports++
				break
			}
		}
	}
	return safeReports
}

func main() {
	reports := readInput()
	fmt.Println(len(reports))
	fmt.Println(int(countSafeReports(reports)))
}
