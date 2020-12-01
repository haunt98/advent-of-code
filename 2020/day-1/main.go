package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	defaultArrLen = 1000
)

func main() {
	fmt.Println("2020/day-1")

	file, err := os.Open("2020/day-1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	arr := make([]int, 0, defaultArrLen)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		num, err := strconv.Atoi(line)
		if err != nil {
			continue
		}

		arr = append(arr, num)
	}

	fmt.Println("part 1:", part_1(arr))
	fmt.Println("part 2:", part_2(arr))
}

func part_1(arr []int) int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == 2020 {
				return arr[i] * arr[j]
			}
		}
	}

	return 0
}

func part_2(arr []int) int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				if arr[i]+arr[j]+arr[k] == 2020 {
					return arr[i] * arr[j] * arr[k]
				}
			}
		}
	}

	return 0
}
