package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	fmt.Println("2020/day-6")

	bytes, err := ioutil.ReadFile("2020/day-6/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	groups := strings.Split(string(bytes), "\n\n")

	fmt.Println(part_1(groups))
	fmt.Println(part_2(groups))

	fmt.Println(countAnswer("ab\nac\n"))
}

func part_1(groups []string) int {
	result := 0

	for _, group := range groups {
		_, exist, _ := countAnswer(group)
		result += exist
	}

	return result
}

func part_2(groups []string) int {
	result := 0

	for _, group := range groups {
		_, _, same := countAnswer(group)
		result += same
	}

	return result
}

func countAnswer(group string) (m map[rune]int, exist int, same int) {
	m = map[rune]int{
		'a': 0,
		'b': 0,
		'c': 0,
		'x': 0,
		'y': 0,
		'z': 0,
	}

	group = strings.TrimSpace(group)
	answers := strings.Split(group, "\n")

	for _, answer := range answers {
		for _, question := range answer {
			m[question]++
		}
	}

	for question := range m {
		if m[question] > 0 {
			exist++
		}

		if m[question] == len(answers) {
			same++
		}
	}

	return m, exist, same
}
