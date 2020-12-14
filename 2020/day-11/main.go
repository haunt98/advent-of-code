package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	fmt.Println("2020/day-11")

	bytes, err := ioutil.ReadFile("2020/day-11/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	part_1(lines)
	part_2(lines)
}

func part_1(lines []string) {
	seats := make([]string, 0, 1000)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		seats = append(seats, line)
	}

	for {
		newSeats := change(seats)

		if isEqual(newSeats, seats) {
			break
		}

		seats = newSeats
	}

	occupiedCount := countAppear(seats, "#")

	fmt.Println(occupiedCount)
}

func part_2(lines []string) {
	seats := make([]string, 0, 1000)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		seats = append(seats, line)
	}

	for {
		newSeats := change_2(seats)

		if isEqual(newSeats, seats) {
			break
		}

		seats = newSeats
	}

	occupiedCount := countAppear(seats, "#")

	fmt.Println(occupiedCount)
}

func change(seats []string) []string {
	newSeats := make([]string, len(seats))

	for i, seat := range seats {
		newSeat := ""

		for j := range seat {
			if string(seat[j]) == "." {
				newSeat += "."
				continue
			}

			if string(seat[j]) == "L" {
				if countAdjacent(seats, i, j, "#") == 0 {
					newSeat += "#"
				} else {
					newSeat += "L"
				}
				continue
			}

			if string(seat[j]) == "#" {
				if countAdjacent(seats, i, j, "#") >= 4 {
					newSeat += "L"
				} else {
					newSeat += "#"
				}
				continue
			}
		}

		newSeats[i] = newSeat
	}

	return newSeats
}

func change_2(seats []string) []string {
	newSeats := make([]string, len(seats))

	for i, seat := range seats {
		newSeat := ""

		for j := range seat {
			if string(seat[j]) == "." {
				newSeat += "."
				continue
			}

			if string(seat[j]) == "L" {
				if countOccupiedAdjacent_2(seats, i, j) == 0 {
					newSeat += "#"
				} else {
					newSeat += "L"
				}
				continue
			}

			if string(seat[j]) == "#" {
				if countOccupiedAdjacent_2(seats, i, j) >= 5 {
					newSeat += "L"
				} else {
					newSeat += "#"
				}
				continue
			}
		}

		newSeats[i] = newSeat
	}

	return newSeats
}

func isEqual(seats_1, seats_2 []string) bool {
	if len(seats_1) != len(seats_2) {
		log.Fatal("impossible")
	}

	for i := range seats_1 {
		if seats_1[i] != seats_2[i] {
			return false
		}
	}

	return true
}

// XXX
// X?X
// XXX
func countAdjacent(seats []string, i, j int, target string) int {
	if i < 0 || i >= len(seats) || j < 0 || j >= len(seats[i]) {
		return 0
	}

	count := 0

	if i-1 >= 0 {
		if string(seats[i-1][j]) == target {
			count++
		}

		if j-1 >= 0 && string(seats[i-1][j-1]) == target {
			count++
		}

		if j+1 < len(seats[i-1]) && string(seats[i-1][j+1]) == target {
			count++
		}
	}

	if j-1 >= 0 && string(seats[i][j-1]) == target {
		count++
	}

	if j+1 < len(seats[i]) && string(seats[i][j+1]) == target {
		count++
	}

	if i+1 < len(seats) {
		if string(seats[i+1][j]) == target {
			count++
		}

		if j-1 >= 0 && string(seats[i+1][j-1]) == target {
			count++
		}

		if j+1 < len(seats[i+1]) && string(seats[i+1][j+1]) == target {
			count++
		}
	}

	return count
}

func countOccupiedAdjacent_2(seats []string, i, j int) int {
	if i < 0 || i >= len(seats) || j < 0 || j >= len(seats[i]) {
		return 0
	}

	count := 0

	// up
	temp_i := i - 1
	for temp_i >= 0 {
		if string(seats[temp_i][j]) == "#" {
			count++
			break
		}

		if string(seats[temp_i][j]) == "L" {
			break
		}

		temp_i--
	}

	// down
	temp_i = i + 1
	for temp_i < len(seats) {
		if string(seats[temp_i][j]) == "#" {
			count++
			break
		}

		if string(seats[temp_i][j]) == "L" {
			break
		}

		temp_i++
	}

	// left
	temp_j := j - 1
	for temp_j >= 0 {
		if string(seats[i][temp_j]) == "#" {
			count++
			break
		}

		if string(seats[i][temp_j]) == "L" {
			break
		}

		temp_j--
	}

	// right
	temp_j = j + 1
	for temp_j < len(seats[i]) {
		if string(seats[i][temp_j]) == "#" {
			count++
			break
		}

		if string(seats[i][temp_j]) == "L" {
			break
		}

		temp_j++
	}

	// up and right
	temp_i = i - 1
	temp_j = j + 1
	for temp_i >= 0 && temp_j < len(seats[i]) {
		if string(seats[temp_i][temp_j]) == "#" {
			count++
			break
		}

		if string(seats[temp_i][temp_j]) == "L" {
			break
		}

		temp_i--
		temp_j++
	}

	// up and left
	temp_i = i - 1
	temp_j = j - 1
	for temp_i >= 0 && temp_j >= 0 {
		if string(seats[temp_i][temp_j]) == "#" {
			count++
			break
		}

		if string(seats[temp_i][temp_j]) == "L" {
			break
		}

		temp_i--
		temp_j--
	}

	// down and right
	temp_i = i + 1
	temp_j = j + 1
	for temp_i < len(seats) && temp_j < len(seats[i]) {
		if string(seats[temp_i][temp_j]) == "#" {
			count++
			break
		}

		if string(seats[temp_i][temp_j]) == "L" {
			break
		}

		temp_i++
		temp_j++
	}

	// down and left
	temp_i = i + 1
	temp_j = j - 1
	for temp_i < len(seats) && temp_j >= 0 {
		if string(seats[temp_i][temp_j]) == "#" {
			count++
			break
		}

		if string(seats[temp_i][temp_j]) == "L" {
			break
		}

		temp_i++
		temp_j--
	}

	return count
}

func countAppear(seats []string, target string) int {
	count := 0

	for i := 0; i < len(seats); i++ {
		for j := 0; j < len(seats[i]); j++ {
			if string(seats[i][j]) == target {
				count++
			}
		}
	}

	return count
}
