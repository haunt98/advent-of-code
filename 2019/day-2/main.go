package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("2019/day-2")

	bytes, err := ioutil.ReadFile("2019/day-2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := strings.TrimSpace(string(bytes))

	arr := strings.Split(data, ",")

	nums := make([]int, len(arr))

	for i := range arr {
		num, err := strconv.Atoi(arr[i])
		if err != nil {
			log.Fatal(err)
		}

		nums[i] = num
	}

	temp := make([]int, len(nums))

	copy(temp, nums)
	fmt.Println(part_1(temp))

	copy(temp, nums)
	fmt.Println(part_2(temp))
}

func part_1(nums []int) int {
	result, ok := calcIntcode(nums, 12, 2)
	if !ok {
		return 0
	}

	return result
}

func part_2(nums []int) int {
	temp := make([]int, len(nums))

	for noun := 0; noun < len(nums); noun++ {
		for verb := 0; verb < len(nums); verb++ {
			copy(temp, nums)

			result, ok := calcIntcode(temp, noun, verb)
			if !ok {
				continue
			}

			if result != 19690720 {
				continue
			}

			return noun*100 + verb
		}
	}

	return 0
}

func calcIntcode(nums []int, noun, verb int) (int, bool) {
	nums[1] = noun
	nums[2] = verb

	for i := 0; i < len(nums); i += 4 {
		if nums[i] == 99 {
			break
		}

		if nums[i+1] >= len(nums) ||
			nums[i+2] >= len(nums) ||
			nums[i+3] >= len(nums) {
			return 0, false
		}

		if nums[i] == 1 {
			nums[nums[i+3]] = nums[nums[i+1]] + nums[nums[i+2]]
			continue
		}

		if nums[i] == 2 {
			nums[nums[i+3]] = nums[nums[i+1]] * nums[nums[i+2]]
			continue
		}

		return 0, false
	}

	return nums[0], true
}
