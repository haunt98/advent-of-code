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

	bytes, err := ioutil.ReadFile("2020/day-10/input.txt")
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
