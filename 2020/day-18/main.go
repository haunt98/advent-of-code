package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/haunt98/advent-of-code/pkg/parser"
	"github.com/haunt98/panda/pkg/stack"
)

func main() {
	fmt.Println("2020/day-18")

	lines, err := parser.ParseLines("2020/day-18/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part_1(lines)
	part_2(lines)
}

func part_1(lines []string) {
	result := 0

	exprs := parseExpressions(lines)
	for _, expr := range exprs {
		rpn := parseRPN(expr, getPrecedence1)
		result += calculateRPN(rpn)
	}

	fmt.Println(result)
}

func part_2(lines []string) {
	result := 0

	exprs := parseExpressions(lines)
	for _, expr := range exprs {
		rpn := parseRPN(expr, getPrecedence2)
		result += calculateRPN(rpn)
	}

	fmt.Println(result)
}

func parseExpressions(lines []string) [][]string {
	exprs := make([][]string, 0, 100)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		expr := parseExpression(line)
		exprs = append(exprs, expr)
	}
	return exprs
}

func parseExpression(line string) []string {
	result := make([]string, 0, 10)
	for _, c := range line {
		if c == ' ' {
			continue
		}

		result = append(result, string(c))
	}
	return result
}

type getPrecedenceFn func(e string) int

// Reverse Polish Notation
// https://en.wikipedia.org/wiki/Shunting-yard_algorithm
func parseRPN(expr []string, getPrecedenceFn getPrecedenceFn) []string {
	operatorStack := stack.NewStackString()
	result := make([]string, 0, len(expr))
	for _, e := range expr {
		if isNumber(e) {
			result = append(result, e)
			continue
		}

		if isOperator(e) {
			for {
				op, ok := operatorStack.Pop()
				if !ok {
					break
				}

				if op == "(" {
					operatorStack.Push(op)
					break
				}

				if getPrecedenceFn(op) < getPrecedenceFn(e) {
					operatorStack.Push(op)
					break
				}

				result = append(result, op)
			}

			operatorStack.Push(e)
			continue
		}

		if e == "(" {
			operatorStack.Push(e)
			continue
		}

		if e == ")" {
			for {
				op, ok := operatorStack.Pop()
				if !ok {
					break
				}

				if op == "(" {
					break
				}

				result = append(result, op)
			}
		}
	}

	for {
		op, ok := operatorStack.Pop()
		if !ok {
			break
		}

		result = append(result, op)
	}

	return result
}

func getPrecedence1(e string) int {
	return 0
}

func getPrecedence2(e string) int {
	switch e {
	case "+", "-":
		return 2
	case "*", "/":
		return 1
	default:
		return 0
	}
}

func calculateRPN(rpn []string) int {
	s := stack.NewStackInt()

	for _, e := range rpn {
		if isOperator(e) {
			second, ok := s.Pop()
			if !ok {
				return 0
			}

			first, ok := s.Pop()
			if !ok {
				return 0
			}

			switch e {
			case "+":
				s.Push(first + second)
			case "-":
				s.Push(first - second)
			case "*":
				s.Push(first * second)
			case "/":
				s.Push(first / second)
			default:
				return 0
			}
			continue
		}

		num, err := strconv.Atoi(e)
		if err != nil {
			return 0
		}

		s.Push(num)
	}

	result, ok := s.Pop()
	if !ok {
		return 0
	}

	return result
}

func isNumber(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}

	return true
}

func isOperator(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/"
}
