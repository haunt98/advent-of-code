package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("2019/day-6")

	bytes, err := ioutil.ReadFile("2019/day-6/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(bytes), "\n")

	part_1(lines)
	part_2(lines)
}

func part_1(lines []string) {
	m := readLines(lines)
	if isDebug() {
		fmt.Println(m)
	}

	alreadyCount := make(map[string]int)

	for orbit := range m {
		countOrbit(m, alreadyCount, orbit)
	}

	if isDebug() {
		fmt.Println(alreadyCount)
	}

	result := 0

	for _, count := range alreadyCount {
		result += count
	}

	fmt.Println(result)
}

func part_2(lines []string) {
	m := readLines(lines)

	pathYOU := traverse(m, "YOU")
	if isDebug() {
		fmt.Println(pathYOU)
	}

	pathSAN := traverse(m, "SAN")
	if isDebug() {
		fmt.Println(pathSAN)
	}

	min := 0
	minEnable := false

	for i := range pathYOU {
		for j := range pathSAN {
			if pathYOU[i] == pathSAN[j] {
				if isDebug() {
					fmt.Println(i, j)
				}

				step := i + j
				if !minEnable || min > step {
					min = step
					minEnable = true
				}
			}
		}
	}

	fmt.Println(min)
}

func isDebug() bool {
	return os.Getenv("DEBUG") != ""
}

func countOrbit(m map[string]string, alreadyCount map[string]int, target string) {
	if _, ok := m[target]; !ok {
		alreadyCount[target] = 0
		return
	}

	if _, ok := alreadyCount[target]; ok {
		return
	}

	countOrbit(m, alreadyCount, m[target])

	alreadyCount[target] = alreadyCount[m[target]] + 1
}

func traverse(m map[string]string, target string) []string {
	arr := make([]string, 0, 1000)

	for {
		root, ok := m[target]
		if !ok {
			break
		}

		arr = append(arr, root)

		target = root
	}

	return arr
}

func readLines(lines []string) map[string]string {
	m := make(map[string]string)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		root, orbit := readLine(line)
		m[orbit] = root
	}

	return m
}

func readLine(line string) (root string, orbit string) {
	rootOrbit := strings.Split(line, ")")
	root = strings.TrimSpace(rootOrbit[0])
	orbit = strings.TrimSpace(rootOrbit[1])
	return
}
