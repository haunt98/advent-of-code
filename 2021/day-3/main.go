package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/make-go-great/panda-go/parser"
)

func main() {
	fmt.Println("2021/day-3")

	lines, err := parser.ParseLines("2021/day-3/input.txt")
	if err != nil {
		log.Fatalln("err", err)
	}

	fmt.Println("part_1:", part_1(lines))
	fmt.Println("part_2", part_2(lines))
}

func part_1(lines []string) int64 {
	most, least := getCommonBits(lines)
	return most * least
}

func part_2(lines []string) int64 {
	most, least := getCommonBitsThenFilter(lines)
	return most * least
}

func getCommonBits(lines []string) (most, least int64) {
	if len(lines) == 0 {
		return
	}

	// Count most, least common bits

	bitsLen := len(lines[0])
	bits := make([]int, bitsLen)

	for _, line := range lines {
		for i, c := range line {
			if c == '1' {
				bits[i] += 1
			} else if c == '0' {
				bits[i] -= 1
			}
		}
	}

	// Get most, least common bits

	mostBitsStr := ""
	leastBitsStr := ""
	for _, b := range bits {
		if b > 0 {
			mostBitsStr += "1"
			leastBitsStr += "0"
		} else {
			mostBitsStr += "0"
			leastBitsStr += "1"
		}
	}

	// Convert binary to decimal

	var err error
	most, err = strconv.ParseInt(mostBitsStr, 2, 64)
	if err != nil {
		log.Fatalln("err", err)
	}

	least, err = strconv.ParseInt(leastBitsStr, 2, 64)
	if err != nil {
		log.Fatalln("err", err)
	}

	return
}

func getCommonBitsThenFilter(lines []string) (most, least int64) {
	if len(lines) == 0 {
		return
	}

	bitsLen := len(lines[0])

	// Filter

	mostLines := make([]string, len(lines))
	copy(mostLines, lines)

	leastLines := make([]string, len(lines))
	copy(leastLines, lines)

	for i := 0; i < bitsLen; i++ {
		if len(mostLines) != 1 {
			mostLines, _ = getCommonBitsAtPosition(mostLines, i)
		}

		if len(leastLines) != 1 {
			_, leastLines = getCommonBitsAtPosition(leastLines, i)
		}
	}

	// Convert binary to decimal

	var err error
	most, err = strconv.ParseInt(mostLines[0], 2, 64)
	if err != nil {
		log.Fatalln("err", err)
	}

	least, err = strconv.ParseInt(leastLines[0], 2, 64)
	if err != nil {
		log.Fatalln("err", err)
	}

	return
}

// 1001
// 1101
// 1011
// position 2
// most: 0 -> 1001, 1011
// least: 1 -> 1101
func getCommonBitsAtPosition(lines []string, position int) (mostLines, leastLines []string) {
	if len(lines) == 0 {
		return
	}

	// Count most, least common bits at position

	countBit := 0

	for _, line := range lines {
		if line[position] == '1' {
			countBit += 1
		} else if line[position] == '0' {
			countBit -= 1
		}
	}

	// Get most, least common bits at position

	var mostBit rune
	var leastBit rune
	if countBit > 0 {
		mostBit = '1'
		leastBit = '0'
	} else if countBit < 0 {
		mostBit = '0'
		leastBit = '1'
	} else {
		mostBit = '1'
		leastBit = '0'
	}

	mostLines = make([]string, 0, len(lines))
	for _, line := range lines {
		if line[position] == byte(mostBit) {
			mostLines = append(mostLines, line)
		}
	}

	leastLines = make([]string, 0, len(lines))
	for _, line := range lines {
		if line[position] == byte(leastBit) {
			leastLines = append(leastLines, line)
		}
	}

	return mostLines, leastLines
}
