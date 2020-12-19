package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("2020/day-15")

	fmt.Println(part_1([]int{1, 2, 16, 19, 18, 0}))
	fmt.Println(part_2([]int{1, 2, 16, 19, 18, 0}))
}

func part_1(inits []int) int {
	return memoryGame(inits, 2020)
}

func part_2(inits []int) int {
	return memoryGame(inits, 30000000)
}

func memoryGame(inits []int, finalStep int) int {
	step := 1

	// save spoken number
	memories := make(map[int][]int)

	// save step
	steps := make(map[int]int)

	// run from init
	for _, init := range inits {
		// -1 is not yet to be exist
		memories[init] = []int{-1, step}
		steps[step] = init
		step++
	}

	// run freely
	for {
		if step > finalStep {
			break
		}

		// previous
		prev := steps[step-1]
		mem := memories[prev]
		if len(mem) != 2 {
			log.Fatal("god bless you")
		}

		// check turn apart or first appear
		value := 0
		if mem[0] == -1 {
		} else {
			value = mem[1] - mem[0]
		}

		steps[step] = value
		if memValue, ok := memories[value]; !ok {
			memories[value] = []int{-1, step}
		} else {
			if len(memValue) != 2 {
				log.Fatal("god bless you again")
			}
			memories[value] = []int{memValue[1], step}
		}

		step++
	}

	return steps[finalStep]
}
