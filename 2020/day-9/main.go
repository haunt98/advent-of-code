package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("2020/day-9")

	bytes, err := ioutil.ReadFile("2020/day-9/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(bytes), "\n")

	nums := make([]int, 0, 10000)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		nums = append(nums, num)
	}

	part_1(nums)
	part_2(nums)
}

func part_1(nums []int) {
	imposter := findImposter(nums, 25)

	fmt.Println(imposter)
}

func part_2(nums []int) {
	imposter := findImposter(nums, 25)

	arr := existSum_2(nums, imposter)
	if len(arr) == 0 {
		log.Fatal("wrong len")
	}

	min := arr[0]
	max := arr[0]

	for i := 1; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		}

		if max < arr[i] {
			max = arr[i]
		}
	}

	result := min + max

	fmt.Println(result)
}

func findImposter(nums []int, preamleCount int) int {
	for i := preamleCount; i < len(nums); i++ {
		if !existSum_1(nums[i-preamleCount:i], nums[i]) {
			return nums[i]
		}
	}

	return 0
}

func existSum_1(arr []int, target int) bool {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				return true
			}
		}
	}

	return false
}

func existSum_2(arr []int, target int) []int {
	for count := len(arr) - 1; count >= 2; count-- {
		for i := 0; i < len(arr); i++ {
			if i+count > len(arr) {
				break
			}

			if sum(arr[i:i+count]) == target {
				return arr[i : i+count]
			}
		}
	}

	return nil
}

func sum(arr []int) int {
	result := 0

	for i := range arr {
		result += arr[i]
	}

	return result
}
