package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("2019/day-4")

	fmt.Println(part_1())
	fmt.Println(part_2())
	fmt.Println(ruleLargerAdjacent("112233"))
	fmt.Println(ruleLargerAdjacent("123444"))
	fmt.Println(ruleLargerAdjacent("111122"))
}

func part_1() int {
	ruleFns := []ruleFn{
		ruleLen,
		ruleAdjacent,
		ruleIncrease,
	}

	result := 0

	for i := 138241; i <= 674034; i++ {
		v := strconv.FormatInt(int64(i), 10)

		valid := true

		for _, fn := range ruleFns {
			if !fn(v) {
				valid = false
				break
			}
		}

		if valid {
			result++
		}
	}

	return result
}

func part_2() int {
	ruleFns := []ruleFn{
		ruleLen,
		ruleLargerAdjacent,
		ruleIncrease,
	}

	result := 0

	for i := 138241; i <= 674034; i++ {
		v := strconv.FormatInt(int64(i), 10)

		valid := true

		for _, fn := range ruleFns {
			if !fn(v) {
				valid = false
				break
			}
		}

		if valid {
			result++
		}
	}

	return result
}

type ruleFn func(v string) bool

func ruleLen(v string) bool {
	return len(v) == 6
}

func ruleAdjacent(v string) bool {
	for i := 0; i < len(v)-1; i++ {
		if v[i] == v[i+1] {
			return true
		}
	}

	return false
}

func ruleIncrease(v string) bool {
	for i := 0; i < len(v)-1; i++ {
		if v[i] > v[i+1] {
			return false
		}
	}

	return true
}

func ruleLargerAdjacent(v string) bool {
	for i := 0; i < len(v); {
		from := i
		to := i
		j := i + 1

		for ; j < len(v); j++ {
			if v[j] == v[i] {
				to = j
				continue
			}

			break
		}

		i = j

		if to-from == 1 {
			return true
		}

	}

	return false
}
