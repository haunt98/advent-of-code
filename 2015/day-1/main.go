package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("2015/day-1")

	s := prepareInput()
	part_1(s)
	part_2(s)
}

func prepareInput() string {
	bytes, err := os.ReadFile("2015/day-1/input.txt")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	s := strings.TrimSpace(string(bytes))

	return s
}

func part_1(s string) {
	fmt.Println(goFloor(s))
}

func part_2(s string) {
	fmt.Println(goIndexFloorUntil(s, -1))
}

func goFloor(s string) int {
	floor := 0

	for _, r := range s {
		floor += instructionFloor(r)
	}

	return floor
}

func goIndexFloorUntil(s string, stopFloor int) int {
	floor := 0

	for i, r := range s {
		floor += instructionFloor(r)
		if floor == stopFloor {
			// Because index floor start from 1
			return i + 1
		}
	}

	return 0
}

// ( -> go up -> + 1
// ) -> go down -> - 1
func instructionFloor(r rune) int {
	switch r {
	case '(':
		return 1
	case ')':
		return -1
	default:
		return 0
	}
}
