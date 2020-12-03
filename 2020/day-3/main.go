package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	defaultLen = 1000
)

func main() {
	fmt.Println("2020/day-3")

	file, err := os.Open("2020/day-3/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	areaMap := make([]string, 0, defaultLen)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		areaMap = append(areaMap, line)
	}

	fmt.Println(part_1(areaMap))
	fmt.Println(part_2(areaMap))
}

func part_1(areaMap []string) int {
	count := 0
	position := 0

	for i := 1; i < len(areaMap); i++ {
		position += 3
		realPosition := position % len(areaMap[i])
		if areaMap[i][realPosition] == '#' {
			count++
		}
	}

	return count
}

func part_2(areaMap []string) int {
	// 1_1 is right 1 down 1
	count_1_1 := 0
	count_3_1 := 0
	count_5_1 := 0
	count_7_1 := 0
	count_1_2 := 0

	position_1_1 := 0
	position_3_1 := 0
	position_5_1 := 0
	position_7_1 := 0
	position_1_2 := 0

	for i := 1; i < len(areaMap); i++ {
		position_1_1 += 1
		realPos_1_1 := position_1_1 % len(areaMap[i])
		if areaMap[i][realPos_1_1] == '#' {
			count_1_1++
		}

		position_3_1 += 3
		realPos_3_1 := position_3_1 % len(areaMap[i])
		if areaMap[i][realPos_3_1] == '#' {
			count_3_1++
		}

		position_5_1 += 5
		realPos_5_1 := position_5_1 % len(areaMap[i])
		if areaMap[i][realPos_5_1] == '#' {
			count_5_1++
		}

		position_7_1 += 7
		realPos_7_1 := position_7_1 % len(areaMap[i])
		if areaMap[i][realPos_7_1] == '#' {
			count_7_1++
		}

		if i%2 == 0 {
			position_1_2 += 1
			realPos_1_2 := position_1_2 % len(areaMap[i])
			if areaMap[i][realPos_1_2] == '#' {
				count_1_2++
			}
		}
	}

	return count_1_1 * count_3_1 * count_5_1 * count_7_1 * count_1_2
}
