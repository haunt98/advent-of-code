package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("2020/day-7")

	bytes, err := ioutil.ReadFile("2020/day-7/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(bytes), "\n")

	part_1(lines)
	part_2(lines)
}

func part_1(lines []string) {
	m := readLines(lines)

	result := 0

	canContain := make(map[bag]struct{})

	for b, _ := range m {
		if isContain(m, canContain, b, "shiny gold") {
			result++
		}
	}

	fmt.Println(result)
}

func part_2(lines []string) {
	m := readLines(lines)

	result := countContain(m, "shiny gold")

	fmt.Println(result)
}

type bag string
type capacity map[bag]int

func readLines(lines []string) map[bag]capacity {
	m := make(map[bag]capacity)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		bag, capacity := readLine(line)

		m[bag] = capacity
	}

	return m
}

func readLine(line string) (bag, capacity) {
	rawBagCapacity := strings.Split(line, "contain")

	rawBag := strings.TrimSpace(rawBagCapacity[0])
	rawBagFields := strings.Split(rawBag, " ")

	b := strings.Join(rawBagFields[:2], " ")

	if strings.TrimSpace(rawBagCapacity[1]) == "no other bags." {
		return bag(b), nil
	}

	c := make(map[bag]int)

	rawCapacity := strings.Split(rawBagCapacity[1], ",")

	for _, rawSingleCapacity := range rawCapacity {
		r := strings.TrimSpace(rawSingleCapacity)

		rFields := strings.Split(r, " ")

		num, err := strconv.Atoi(rFields[0])
		if err != nil {
			log.Fatal(err)
		}

		rB := strings.Join(rFields[1:3], " ")

		c[bag(rB)] = num

	}

	return bag(b), c
}

func isContain(m map[bag]capacity, canContain map[bag]struct{}, root, target bag) bool {
	if _, ok := canContain[root]; ok {
		return true
	}

	if val, ok := m[root]; !ok || val == nil {
		return false
	}

	for b, _ := range m[root] {
		if b == target {
			return true
		}

		if _, ok := canContain[b]; ok {
			return true
		}

		if isContain(m, canContain, b, target) {
			canContain[b] = struct{}{}
			return true
		}

	}

	return false
}

func countContain(m map[bag]capacity, target bag) int {
	if _, ok := m[target]; !ok {
		return 0
	}

	result := 0

	for b, num := range m[target] {
		result += num * (1 + countContain(m, b))
	}

	return result
}
