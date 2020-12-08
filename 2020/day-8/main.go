package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("2020/day-8")

	bytes, err := ioutil.ReadFile("2020/day-8/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(bytes), "\n")

	part_1(lines)
	part_2(lines)
}

func part_1(lines []string) {
	if isDebug() {
		fmt.Println(lines)
	}

	m := make(map[int]instruction)

	index := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		m[index] = parseInstr(line)

		index++
	}

	if isDebug() {
		fmt.Println(m)
	}

	accumulator, isLoop := run(m)
	if isDebug() {
		fmt.Println(isLoop)
	}

	fmt.Println(accumulator)
}

func part_2(lines []string) {
	m := make(map[int]instruction)

	mJMP := make(map[int]struct{})
	mNOP := make(map[int]struct{})

	index := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		instr := parseInstr(line)
		m[index] = instr

		if instr.operation == "jmp" {
			mJMP[index] = struct{}{}
		}

		if instr.operation == "nop" {
			mNOP[index] = struct{}{}
		}

		index++
	}

	if isDebug() {
		fmt.Println(m)
		fmt.Println(mJMP)
		fmt.Println(mNOP)
	}

	// try to replace jmp with nop
	for i := range mJMP {
		m[i] = instruction{
			operation: "nop",
			arg:       m[i].arg,
		}

		accumulator, isLoop := run(m)
		if isLoop == false {
			fmt.Println(accumulator)
			return
		}

		// switch back
		m[i] = instruction{
			operation: "jmp",
			arg:       m[i].arg,
		}
	}

	// try to replace nop with jmp
	for i := range mNOP {
		m[i] = instruction{
			operation: "jmp",
			arg:       m[i].arg,
		}

		accumulator, isLoop := run(m)
		if isLoop == false {
			fmt.Println(accumulator)
			return
		}

		// switch back
		m[i] = instruction{
			operation: "nop",
			arg:       m[i].arg,
		}
	}

	fmt.Println("something wrong, I can smell that")
}

func isDebug() bool {
	return os.Getenv("DEBUG") != ""
}

type instruction struct {
	operation string
	arg       argument
}

type argument struct {
	sign string
	num  int
}

func run(m map[int]instruction) (accumulator int, isLoop bool) {
	cur := 0
	visit := make(map[int]struct{})

	for {
		if cur >= len(m) {
			break
		}

		// loop
		if _, ok := visit[cur]; ok {
			isLoop = true
			break
		}

		visit[cur] = struct{}{}

		instr, ok := m[cur]
		if !ok {
			log.Fatal("wrong cur")
		}

		if isDebug() {
			fmt.Println(cur, instr)
		}

		switch instr.operation {
		case "acc":
			switch instr.arg.sign {
			case "+":
				accumulator += instr.arg.num
			case "-":
				accumulator -= instr.arg.num
			default:
				log.Fatal("wrong sign")
			}

			cur += 1
		case "jmp":
			switch instr.arg.sign {
			case "+":
				cur += instr.arg.num
			case "-":
				cur -= instr.arg.num
			default:
				log.Fatal("wrong sign")
			}
		case "nop":
			cur += 1
		default:
			log.Fatal("wrong operation")
		}
	}

	return
}

func parseInstr(line string) instruction {
	var operation string
	var sign string
	var num int

	fmt.Sscanf(line, "%3s %1s%d", &operation, &sign, &num)

	return instruction{
		operation: operation,
		arg: argument{
			sign: sign,
			num:  num,
		},
	}
}
