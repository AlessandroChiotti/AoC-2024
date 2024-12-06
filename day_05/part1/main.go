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

func readInput() (map[int][]int, string) {
	data, err := os.ReadFile("./day_05/input.txt")
	check(err)
	nextPagesNotAllowedByPage := make(map[int][]int)

	inputs := strings.Split(string(data), "\n\n")

	rules := strings.Split(inputs[0], "\n")
	for _, rule := range rules {
		ruleParts := strings.Split(rule, "|")
		n1, err := strconv.Atoi(ruleParts[0])
		check(err)
		n2, err := strconv.Atoi(ruleParts[1])
		check(err)
		nextPagesNotAllowedByPage[n2] = append(nextPagesNotAllowedByPage[n2], n1)
	}

	return nextPagesNotAllowedByPage, inputs[1]
}

func computeCorrectUpdates(nextPagesNotAllowedByPage map[int][]int, updates string) []string {

	var correctUpdates []string

	updatesParts := strings.Split(updates, "\n")
	for _, update := range updatesParts {
		pageNumbers := strings.Split(update, ",")
		isInvalid := false
		for i, pageNumber := range pageNumbers {
			if pageNumber == "" {
				continue
			}
			page, err := strconv.Atoi(pageNumber)
			check(err)
			for j := i + 1; j < len(pageNumbers); j++ {
				nextPage, err := strconv.Atoi(pageNumbers[j])
				check(err)

				notAllowedPages := nextPagesNotAllowedByPage[page]
				for _, notAllowedPage := range notAllowedPages {
					if notAllowedPage == nextPage {
						isInvalid = true
						break
					}
				}
				if isInvalid {
					break
				}
			}
			if isInvalid {
				break
			}
		}
		if !isInvalid {
			correctUpdates = append(correctUpdates, update)
		}
	}

	return correctUpdates
}

func computeResult(correctUpdates []string) int {
	result := 0
	for _, update := range correctUpdates {
		pageNumbers := strings.Split(update, ",")
		index := (len(pageNumbers) - 1) / 2
		if pageNumbers[index] == "" {
			continue
		}
		page, err := strconv.Atoi(pageNumbers[index])
		check(err)
		result += page
	}
	return result
}

func main() {
	nextPagesNotAllowedByPage, updates := readInput()
	correctUpdates := computeCorrectUpdates(nextPagesNotAllowedByPage, updates)
	result := computeResult(correctUpdates)
	fmt.Println("result: ", result)
}
