package main

import (
	"fmt"
	"log"

	"github.com/make-go-great/panda-go/parser"
)

func main() {
	fmt.Println("2021/day-1")

	arr, err := parser.ParseIntLines("2021/day-1/input.txt")
	if err != nil {
		log.Fatalln("Failed to parse int lines", err)
	}

	fmt.Println("part 1:", part_1(arr))
	fmt.Println("part 2:", part_2(arr))
}

func part_1(arr []int) int {
	return countIncrease(arr)
}

func part_2(arr []int) int {
	return countIncrease3Windows(arr)
}

// 1 2 3
// 2 > 1 and 3 > 1 => count =2
func countIncrease(arr []int) int {
	if len(arr) <= 1 {
		return 0
	}

	count := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			count += 1
		}
	}

	return count
}

func countIncrease3Windows(arr []int) int {
	if len(arr) <= 3 {
		return 0
	}

	count := 0
	for i := 3; i < len(arr); i++ {
		if arr[i] > arr[i-3] {
			count += 1
		}
	}

	return count
}
