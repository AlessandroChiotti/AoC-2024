package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput() (map[int][]int, map[int][]int, string) {
	data, err := os.ReadFile("./day_05/input.txt")
	check(err)
	nextPagesNotAllowedByPage := make(map[int][]int)
	nextPagesAllowedByPage := make(map[int][]int)

	inputs := strings.Split(string(data), "\n\n")

	rules := strings.Split(inputs[0], "\n")
	for _, rule := range rules {
		ruleParts := strings.Split(rule, "|")
		n1, err := strconv.Atoi(ruleParts[0])
		check(err)
		n2, err := strconv.Atoi(ruleParts[1])
		check(err)
		nextPagesNotAllowedByPage[n2] = append(nextPagesNotAllowedByPage[n2], n1)
		nextPagesAllowedByPage[n1] = append(nextPagesAllowedByPage[n1], n2)

	}

	return nextPagesNotAllowedByPage, nextPagesAllowedByPage, inputs[1]
}

func correctInvalidUpdate(pageNumbers []string, nextPagesAllowedByPage map[int][]int) string {

	var correctedUpdate string
	allowedInUpdateByPage := make(map[int]int)
	for _, pageNumber := range pageNumbers {
		if pageNumber == "" {
			continue
		}
		page, err := strconv.Atoi(pageNumber)
		check(err)
		nextPagesAllowed := nextPagesAllowedByPage[page]
		allowedInUpdateByPage[page] = 0
		for _, nextPageAllowed := range nextPagesAllowed {
			for _, otherNumber := range pageNumbers {
				if pageNumber == "" || pageNumber == otherNumber {
					continue
				}
				otherPage, err := strconv.Atoi(otherNumber)
				check(err)

				if otherPage == nextPageAllowed {
					allowedInUpdateByPage[page] = allowedInUpdateByPage[page] + 1
				}

			}
		}
	}

	keys := make([]int, 0, len(allowedInUpdateByPage))

	for key := range allowedInUpdateByPage {
		keys = append(keys, key)
	}

	slices.SortFunc(keys, func(i, j int) int {
		return allowedInUpdateByPage[j] - allowedInUpdateByPage[i]
	})

	stringKeys := make([]string, len(keys))
	for i, key := range keys {
		stringKeys[i] = strconv.Itoa(key)
	}
	correctedUpdate = strings.Join(stringKeys, ",")

	return correctedUpdate
}

func computeCorrectUpdates(nextPagesNotAllowedByPage map[int][]int, nextPagesAllowedByPage map[int][]int, updates string) []string {

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
			continue
		}

		correctedUpdate := correctInvalidUpdate(pageNumbers, nextPagesAllowedByPage)

		correctUpdates = append(correctUpdates, correctedUpdate)

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
	nextPagesNotAllowedByPage, nextPagesAllowedByPage, updates := readInput()
	correctUpdates := computeCorrectUpdates(nextPagesNotAllowedByPage, nextPagesAllowedByPage, updates)
	result := computeResult(correctUpdates)
	fmt.Println("result: ", result)
}
