package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type policy struct {
	min, max int
	letter   string
}

func main() {
	fmt.Println("2020/day-2")

	file, err := os.Open("2020/day-2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	m := make(map[string]policy)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		var password string
		var policy policy

		fmt.Sscanf(line, "%d-%d %1s: %s", &policy.min, &policy.max, &policy.letter, &password)

		m[password] = policy
	}

	fmt.Println(part_1(m))
	fmt.Println(part_2(m))
}

func part_1(m map[string]policy) int {
	result := 0

	for password, policy := range m {
		count := strings.Count(password, policy.letter)
		if count >= policy.min && count <= policy.max {
			result++
		}
	}

	return result
}

func part_2(m map[string]policy) int {
	result := 0

	for password, policy := range m {
		if string(password[policy.min-1]) == policy.letter &&
			string(password[policy.max-1]) != policy.letter {
			result++
			continue
		}

		if string(password[policy.min-1]) != policy.letter &&
			string(password[policy.max-1]) == policy.letter {
			result++
			continue
		}
	}

	return result
}
