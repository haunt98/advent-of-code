package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("2019/day-5")

	bytes, err := ioutil.ReadFile("2019/day-5/input.txt")
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
	part_1(temp)

	copy(temp, nums)
	part_2(temp)
}

func part_1(nums []int) {
	calcIntcode_1(nums, 1)
}

func part_2(nums []int) {
	calcIntcode_2(nums, 5)
}

type instruction struct {
	opcode int
	mode_1 int
	mode_2 int
	mode_3 int
}

func calcIntcode_1(nums []int, input int) {
	for i := 0; i < len(nums); {
		instr := splitOpcode(nums[i])

		if instr.opcode == 99 {
			return
		}

		switch instr.opcode {
		case 1:
			calcOpcode_1(nums, i, instr)
			i += 4
		case 2:
			calcOpcode_2(nums, i, instr)
			i += 4
		case 3:
			calcOpcode_3(nums, i, instr, input)
			i += 2
		case 4:
			calcOpcode_4(nums, i, instr)
			i += 2
		default:
			return
		}
	}
}

func calcIntcode_2(nums []int, input int) {
	for i := 0; i < len(nums); {
		instr := splitOpcode(nums[i])
		if isDebug() {
			fmt.Printf("nums %+v\n", nums)
			fmt.Printf("index %d %+v\n", i, instr)
		}

		if instr.opcode == 99 {
			return
		}

		switch instr.opcode {
		case 1:
			calcOpcode_1(nums, i, instr)
			i += 4
		case 2:
			calcOpcode_2(nums, i, instr)
			i += 4
		case 3:
			calcOpcode_3(nums, i, instr, input)
			i += 2
		case 4:
			calcOpcode_4(nums, i, instr)
			i += 2
		case 5:
			newIndex, ok := calcOpcode_5(nums, i, instr)
			if !ok {
				return
			}

			i = newIndex
		case 6:
			newIndex, ok := calcOpcode_6(nums, i, instr)
			if !ok {
				return
			}

			i = newIndex
		case 7:
			if ok := calcOpcode_7(nums, i, instr); !ok {
				return
			}

			i += 4
		case 8:
			if ok := calcOpcode_8(nums, i, instr); !ok {
				return
			}

			i += 4
		default:
			return
		}
	}
}

func splitOpcode(rawOpcode int) instruction {
	opcode := rawOpcode % 100
	mode_1 := (rawOpcode / 100) % 10
	mode_2 := (rawOpcode / 1000) % 10
	mode_3 := (rawOpcode / 10000) % 10

	return instruction{
		opcode: opcode,
		mode_1: mode_1,
		mode_2: mode_2,
		mode_3: mode_3,
	}
}

func calcOpcode_1(nums []int, index int, instr instruction) bool {
	val_1, ok := getValue(nums, index+1, instr.mode_1)
	if !ok {
		return false
	}

	val_2, ok := getValue(nums, index+2, instr.mode_2)
	if !ok {
		return false
	}

	// 3

	if instr.mode_3 != 0 {
		return false
	}

	if index+3 >= len(nums) || nums[index+3] >= len(nums) {
		return false
	}

	nums[nums[index+3]] = val_1 + val_2

	return true
}

func calcOpcode_2(nums []int, index int, instr instruction) bool {
	val_1, ok := getValue(nums, index+1, instr.mode_1)
	if !ok {
		return false
	}

	val_2, ok := getValue(nums, index+2, instr.mode_2)
	if !ok {
		return false
	}

	// 3

	if instr.mode_3 != 0 {
		return false
	}

	if index+3 >= len(nums) || nums[index+3] >= len(nums) {
		return false
	}

	nums[nums[index+3]] = val_1 * val_2

	return true
}

func calcOpcode_3(nums []int, index int, instr instruction, input int) bool {
	if instr.mode_1 != 0 {
		return false
	}

	if index+1 >= len(nums) || nums[index+1] >= len(nums) {
		return false
	}

	nums[nums[index+1]] = input

	if isDebug() {
		fmt.Printf("nums[%d] = %d\n", nums[index+1], nums[nums[index+1]])
	}

	return true
}

func calcOpcode_4(nums []int, index int, instr instruction) bool {
	val_1, ok := getValue(nums, index+1, instr.mode_1)
	if !ok {
		return false
	}

	fmt.Printf("OUTPUT: %d\n", val_1)

	return true
}

func calcOpcode_5(nums []int, index int, instr instruction) (newIndex int, ok bool) {
	val_1, ok := getValue(nums, index+1, instr.mode_1)
	if !ok {
		return 0, false
	}

	if isDebug() {
		fmt.Printf("val_1 %d\n", val_1)
	}

	if val_1 == 0 {
		return index + 3, true
	}

	val_2, ok := getValue(nums, index+2, instr.mode_2)
	if !ok {
		return 0, false
	}

	if isDebug() {
		fmt.Printf("val_2 %d\n", val_2)
	}

	return val_2, true
}

func calcOpcode_6(nums []int, index int, instr instruction) (newIndex int, ok bool) {
	val_1, ok := getValue(nums, index+1, instr.mode_1)
	if !ok {
		return 0, false
	}

	if isDebug() {
		fmt.Printf("val_1 %d\n", val_1)
	}

	if val_1 != 0 {
		return index + 3, true
	}

	val_2, ok := getValue(nums, index+2, instr.mode_2)
	if !ok {
		return 0, false
	}

	if isDebug() {
		fmt.Printf("val_2 %d\n", val_2)
	}

	return val_2, true
}

func calcOpcode_7(nums []int, index int, instr instruction) bool {
	val_1, ok := getValue(nums, index+1, instr.mode_1)
	if !ok {
		return false
	}

	if isDebug() {
		fmt.Printf("val_1 %d\n", val_1)
	}

	val_2, ok := getValue(nums, index+2, instr.mode_2)
	if !ok {
		return false
	}

	if isDebug() {
		fmt.Printf("val_2 %d\n", val_2)
	}

	if val_1 < val_2 {
		nums[nums[index+3]] = 1
	} else {
		nums[nums[index+3]] = 0
	}

	if isDebug() {
		fmt.Printf("nums[%d] = %d\n", nums[index+3], nums[nums[index+3]])
	}

	return true
}

func calcOpcode_8(nums []int, index int, instr instruction) bool {
	val_1, ok := getValue(nums, index+1, instr.mode_1)
	if !ok {
		return false
	}

	if isDebug() {
		fmt.Printf("val_1 %d\n", val_1)
	}

	val_2, ok := getValue(nums, index+2, instr.mode_2)
	if !ok {
		return false
	}

	if isDebug() {
		fmt.Printf("val_2 %d\n", val_2)
	}

	if val_1 == val_2 {
		nums[nums[index+3]] = 1
	} else {
		nums[nums[index+3]] = 0
	}

	if isDebug() {
		fmt.Printf("nums[%d] = %d\n", nums[index+3], nums[nums[index+3]])
	}

	return true
}

func getValue(nums []int, index, mode int) (int, bool) {
	if mode == 0 {
		if nums[index] >= len(nums) {
			return 0, false
		}

		return nums[nums[index]], true
	}

	if mode == 1 {
		if index >= len(nums) {
			return 0, false
		}

		return nums[index], true
	}

	return 0, false
}

func isDebug() bool {
	return os.Getenv("DEBUG") != ""
}
