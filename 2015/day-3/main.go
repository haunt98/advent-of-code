package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("2015/day-3")

	s := prepare()
	part_1(s)
	part_2(s)
}

func prepare() string {
	bytes, err := os.ReadFile("2015/day-3/input.txt")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return strings.TrimSpace(string(bytes))
}

func part_1(s string) {
	fmt.Println(countAtLeastOnePresent(s))
}

func part_2(s string) {
	fmt.Println(countAtLeastOnePresentButTakeTurn(s))
}

type position struct {
	x int
	y int
}

func (p position) encode() string {
	return fmt.Sprintf("%d.%d", p.x, p.y)
}

func countAtLeastOnePresent(s string) int {
	passed := make(map[string]int)

	pos := position{
		x: 0,
		y: 0,
	}

	passed[pos.encode()] = 1

	count := 1

	for _, r := range s {
		pos = move(pos, r)
		if _, ok := passed[pos.encode()]; ok {
			passed[pos.encode()]++
		} else {
			passed[pos.encode()] = 1
		}

		if passed[pos.encode()] == 1 {
			count++
		}
	}

	return count
}

func countAtLeastOnePresentButTakeTurn(s string) int {
	passed := make(map[string]int)

	pos := position{
		x: 0,
		y: 0,
	}

	passed[pos.encode()] = 1

	count := 1

	santaPos := pos
	roboPos := pos
	isSanta := true

	for _, r := range s {
		if isSanta {
			pos = santaPos
		} else {
			pos = roboPos
		}

		pos = move(pos, r)
		if _, ok := passed[pos.encode()]; ok {
			passed[pos.encode()]++
		} else {
			passed[pos.encode()] = 1
		}

		if passed[pos.encode()] == 1 {
			count++
		}

		if isSanta {
			santaPos = pos
			isSanta = false
		} else {
			roboPos = pos
			isSanta = true
		}
	}

	return count
}

func move(pos position, r rune) position {
	newPost := position{
		x: pos.x,
		y: pos.y,
	}

	switch r {
	case '^':
		newPost.y += 1
	case 'v':
		newPost.y -= 1
	case '>':
		newPost.x += 1
	case '<':
		newPost.x -= 1
	}

	return newPost
}
