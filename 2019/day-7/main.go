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
	fmt.Println("2019/day-7")

	bytes, err := ioutil.ReadFile("2019/day-7/input.txt")
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
	x := []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26,
		27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}
	perm := []int{9, 8, 7, 6, 5}
	fmt.Println(amplifier_loop(x, perm))
}

func part_1(nums []int) {
	if isDebug() {
		fmt.Println(nums)
	}

	perms := permutation([]int{0, 1, 2, 3, 4})

	max := 0

	for _, perm := range perms {
		output, ok := amplifier(nums, perm)
		if !ok {
			continue
		}

		if max < output {
			max = output
		}
	}

	fmt.Println(max)
}

func part_2(nums []int) {
	perms := permutation([]int{5, 6, 7, 8, 9})

	max := 0

	for _, perm := range perms {
		output, ok := amplifier_loop(nums, perm)
		if !ok {
			continue
		}

		if max < output {
			max = output
		}
	}

	fmt.Println(max)
}

func amplifier(nums, phases []int) (int, bool) {
	temp := make([]int, len(nums))
	var input int

	for i, phase := range phases {
		copy(temp, nums)

		output, ok := runIntcode(temp, []int{phase, input})
		if !ok {
			return 0, false
		}

		input = output

		if i == len(phases)-1 {
			return input, true
		}
	}

	return 0, false
}

func amplifier_loop(nums, phases []int) (int, bool) {
	nums_A := make([]int, len(nums))
	copy(nums_A, nums)

	nums_B := make([]int, len(nums))
	copy(nums_B, nums)

	nums_C := make([]int, len(nums))
	copy(nums_C, nums)

	nums_D := make([]int, len(nums))
	copy(nums_D, nums)

	nums_E := make([]int, len(nums))
	copy(nums_E, nums)

	var input_A int
	var input_B int
	var input_C int
	var input_D int
	var input_E int

	inputIndex_A := 0
	inputIndex_B := 0
	inputIndex_C := 0
	inputIndex_D := 0
	inputIndex_E := 0

	phase_A := phases[0]
	phase_B := phases[1]
	phase_C := phases[2]
	phase_D := phases[3]
	phase_E := phases[4]

	index_A := 0
	index_B := 0
	index_C := 0
	index_D := 0
	index_E := 0

	var output_A int
	var output_B int
	var output_C int
	var output_D int
	var output_E int

	var isExit bool
	var isCrash bool

	for {
		input_A = output_E

		fmt.Println(nums_A, index_A, phase_A, input_A, inputIndex_A)

		output_A, index_A, inputIndex_A, isExit, isCrash = runIntcode_2(nums_A, index_A, []int{phase_A, input_A}, inputIndex_A)
		if isCrash {
			log.Fatal("wtf A")
		}

		input_B = output_A

		output_B, index_B, inputIndex_B, isExit, isCrash = runIntcode_2(nums_B, index_B, []int{phase_B, input_B}, inputIndex_B)
		if isCrash {
			log.Fatal("wtf B")
		}

		input_C = output_B

		output_C, index_C, inputIndex_C, isExit, isCrash = runIntcode_2(nums_C, index_C, []int{phase_C, input_C}, inputIndex_C)
		if isCrash {
			log.Fatal("wtf C")
		}

		input_D = output_C

		output_D, index_D, inputIndex_D, isExit, isCrash = runIntcode_2(nums_D, index_D, []int{phase_D, input_D}, inputIndex_D)
		if isCrash {
			log.Fatal("wtf D")
		}

		input_E = output_D

		output_E, index_E, inputIndex_E, isExit, isCrash = runIntcode_2(nums_E, index_E, []int{phase_E, input_E}, inputIndex_E)
		if isCrash {
			log.Fatal("wtf E")
		}

		if isExit {
			return output_E, true
		}
	}

	return 0, false
}

type instruction struct {
	opcode int
	mode_1 int
	mode_2 int
	mode_3 int
}

func runIntcode(nums []int, inputs []int) (int, bool) {
	var output int
	iInput := 0

	for i := 0; i < len(nums); {
		instr := splitOpcode(nums[i])
		if isDebug() {
			fmt.Printf("nums %+v\n", nums)
			fmt.Printf("index %d %+v\n", i, instr)
		}

		if instr.opcode == 99 {
			return output, true
		}

		switch instr.opcode {
		case 1:
			calcOpcode_1(nums, i, instr)
			i += 4
		case 2:
			calcOpcode_2(nums, i, instr)
			i += 4
		case 3:
			if iInput >= len(inputs) {
				return 0, false
			}

			calcOpcode_3(nums, i, instr, inputs[iInput])

			iInput++
			i += 2
		case 4:
			newOutput, ok := calcOpcode_4(nums, i, instr)
			if !ok {
				return 0, false
			}

			output = newOutput
			i += 2
		case 5:
			newIndex, ok := calcOpcode_5(nums, i, instr)
			if !ok {
				return 0, false
			}

			i = newIndex
		case 6:
			newIndex, ok := calcOpcode_6(nums, i, instr)
			if !ok {
				return 0, false
			}

			i = newIndex
		case 7:
			if ok := calcOpcode_7(nums, i, instr); !ok {
				return 0, false
			}

			i += 4
		case 8:
			if ok := calcOpcode_8(nums, i, instr); !ok {
				return 0, false
			}

			i += 4
		default:
			return 0, false
		}
	}

	return 0, false
}

func runIntcode_2(nums []int, index int, inputs []int, inputIndex int) (output, newIndex, newInputIndex int, isExit, isCrash bool) {
	for i := index; i < len(nums); {
		instr := splitOpcode(nums[i])
		if isDebug() {
			fmt.Printf("nums %+v\n", nums)
			fmt.Printf("index %d %+v\n", i, instr)
		}

		if instr.opcode == 99 {
			isExit = true
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
			if inputIndex >= len(inputs) {
				isExit = true
				isCrash = true
				return
			}

			calcOpcode_3(nums, i, instr, inputs[inputIndex])

			inputIndex++
			i += 2
		case 4:
			newOutput, ok := calcOpcode_4(nums, i, instr)
			if !ok {
				isExit = true
				isCrash = true
				return
			}

			output = newOutput
			newIndex = i + 2
			newInputIndex = inputIndex
			return
		case 5:
			var ok bool
			newIndex, ok = calcOpcode_5(nums, i, instr)
			if !ok {
				isExit = true
				isCrash = true
				return
			}

			i = newIndex
		case 6:
			var ok bool
			newIndex, ok = calcOpcode_6(nums, i, instr)
			if !ok {
				isExit = true
				isCrash = true
				return
			}

			i = newIndex
		case 7:
			if ok := calcOpcode_7(nums, i, instr); !ok {
				isExit = true
				isCrash = true
				return
			}

			i += 4
		case 8:
			if ok := calcOpcode_8(nums, i, instr); !ok {
				isExit = true
				isCrash = true
				return
			}

			i += 4
		default:
			isExit = true
			isCrash = true
			return
		}
	}

	isExit = true
	isCrash = true
	return
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

func calcOpcode_4(nums []int, index int, instr instruction) (int, bool) {
	val_1, ok := getValue(nums, index+1, instr.mode_1)
	if !ok {
		return 0, false
	}

	return val_1, true
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

// https://en.wikipedia.org/wiki/Heap%27s_algorithm
func permutation(arr []int) [][]int {
	perms := make([][]int, 0, 1000)

	generatePermutation(len(arr), arr, &perms)

	return perms
}

func generatePermutation(k int, arr []int, perms *[][]int) {
	if k == 1 {
		result := make([]int, len(arr))
		copy(result, arr)
		*perms = append(*perms, result)
	}

	for i := 0; i < k; i++ {
		generatePermutation(k-1, arr, perms)

		if k%2 == 0 {
			arr[i], arr[k-1] = arr[k-1], arr[i]
		} else {
			arr[0], arr[k-1] = arr[k-1], arr[0]
		}
	}
}
