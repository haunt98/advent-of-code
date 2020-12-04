package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("2019/day-1")

	file, err := os.Open("2019/day-1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	arr := make([]int, 0, 1000)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		num, err := strconv.Atoi(line)
		if err != nil {
			continue
		}

		arr = append(arr, num)
	}

	fmt.Println(part_1(arr))
	fmt.Println(part_2(arr))
}

func part_1(arr []int) int {
	result := 0

	for _, num := range arr {
		result += num/3 - 2
	}

	return result
}

func part_2(arr []int) int {
	result := 0

	for _, num := range arr {
		for {
			num = num/3 - 2
			if num <= 0 {
				break
			}
			result += num
		}
	}

	return result
}
