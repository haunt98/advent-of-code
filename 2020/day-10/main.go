package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("2020/day-10")

	bytes, err := ioutil.ReadFile("2020/day-10/example.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	adapters := make([]int, 0, 1000)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		adapter, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		adapters = append(adapters, adapter)
	}

	sort.Slice(adapters, func(i, j int) bool {
		return adapters[i] < adapters[j]
	})

	part_1(adapters)
	part_2(adapters)
}

func part_1(adapters []int) {
	allRated := append([]int{0}, adapters...)
	allRated = append(allRated, adapters[len(adapters)-1]+3)

	diff_1 := 0
	diff_3 := 0

	for i := 0; i < len(allRated)-1; i++ {
		diff := allRated[i+1] - allRated[i]
		if diff == 1 {
			diff_1++
		} else if diff == 3 {
			diff_3++
		}
	}

	fmt.Println(diff_1 * diff_3)
}

func part_2(adapters []int) {
	allRated := append([]int{0}, adapters...)
	allRated = append(allRated, adapters[len(adapters)-1]+3)

	m := make(map[int]int)
	count := countArrange(allRated, 0, m)

	fmt.Println(count)
}

func countArrange(allRated []int, index int, remember map[int]int) int {
	if index >= len(allRated) {
		return 0
	}

	if result, ok := remember[index]; ok {
		return result
	}

	if index == len(allRated)-1 {
		remember[index] = 1
		return 1
	}

	count := 0

	if index+1 < len(allRated) && allRated[index+1]-allRated[index] <= 3 {
		count += countArrange(allRated, index+1, remember)
	}

	if index+2 < len(allRated) && allRated[index+2]-allRated[index] <= 3 {
		count += countArrange(allRated, index+2, remember)
	}

	if index+3 < len(allRated) && allRated[index+3]-allRated[index] <= 3 {
		count += countArrange(allRated, index+3, remember)
	}

	remember[index] = count
	return count
}
