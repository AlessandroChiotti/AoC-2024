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

const (
	plus rune = '+'
	mul  rune = 'x'
)

type equation struct {
	result   int
	operands []int
}

func readInput() []equation {
	inputData, err := os.ReadFile("./day_07/input.txt")
	check(err)

	lines := strings.Split(strings.TrimSpace(string(inputData)), "\n")
	equations := make([]equation, len(lines))

	for i, line := range lines {
		equationParts := strings.Split(line, ": ")
		equations[i].result, err = strconv.Atoi(equationParts[0])
		check(err)
		operands := strings.Split(equationParts[1], " ")
		equations[i].operands = make([]int, 0, len(operands))
		for _, operand := range operands {
			op, err := strconv.Atoi(operand)
			check(err)
			equations[i].operands = append(equations[i].operands, op)
		}
	}

	return equations
}

func computeEquation(operands []int, operators []rune) int {
	result := operands[0]
	for i := 1; i < len(operands); i++ {
		operator := operators[i-1]
		operand := operands[i]
		if operator == plus {
			result = result + operand
		} else if operator == mul {
			result = result * operand
		} else {
			panic(fmt.Errorf("operator %c not recognized", operator))
		}
	}
	return result
}

func isEquationCorrect(equation equation, operators []rune) bool {
	if len(operators) == len(equation.operands)-1 {
		equationResult := computeEquation(equation.operands, operators)
		return equation.result == equationResult
	}

	operators = append(operators, plus)
	result := isEquationCorrect(equation, operators)
	if !result {
		operators[len(operators)-1] = mul
		result = isEquationCorrect(equation, operators)
	}

	return result
}

func computeResult(equations []equation) int {
	result := 0

	for _, equation := range equations {
		operators := make([]rune, 0, len(equation.operands)-1)
		if isEquationCorrect(equation, operators) {
			result += equation.result
		}
	}

	return result
}

func main() {
	equations := readInput()
	// fmt.Println(equations)
	result := computeResult(equations)
	fmt.Println("result: ", result)
}
