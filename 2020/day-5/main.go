package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	fmt.Println("2020/day-5")

	bytes, err := ioutil.ReadFile("2020/day-5/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	boardingPasses := strings.Fields(strings.TrimSpace(string(bytes)))

	fmt.Println(part_1(boardingPasses))
	fmt.Println(part_2(boardingPasses))
}

func part_1(boardingPasses []string) int {
	max := getSeatID(boardingPasses[0])

	for i := 1; i < len(boardingPasses); i++ {
		seatID := getSeatID(boardingPasses[i])
		if max < seatID {
			max = seatID
		}
	}

	return max
}

func part_2(boardingPasses []string) int {
	m := make(map[int]struct{})

	for _, boardingPass := range boardingPasses {
		m[getSeatID(boardingPass)] = struct{}{}
	}

	// max id is 1023 because 1024 = 2^10

	// skip very front and very back
	for id := 1; id < 1023; id++ {
		if _, ok := m[id]; !ok {
			if _, ok := m[id-1]; !ok {
				continue
			}

			if _, ok := m[id+1]; !ok {
				continue
			}

			return id
		}
	}

	return 0
}

// seat id is binary
// BFFFBBFRRR -> 1000110111
func getSeatID(boardingPass string) int {
	if len(boardingPass) != 10 {
		return 0
	}

	id := 0
	for i := 0; i < 10; i++ {
		if boardingPass[i] == 'F' {
			id = id * 2
			continue
		}

		if boardingPass[i] == 'B' {
			id = id*2 + 1
			continue
		}

		if boardingPass[i] == 'L' {
			id = id * 2
			continue
		}

		if boardingPass[i] == 'R' {
			id = id*2 + 1
			continue
		}
	}

	return id
}
