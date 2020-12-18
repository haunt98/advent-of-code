package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("2020/day-14")

	bytes, err := ioutil.ReadFile("2020/day-14/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := strings.TrimSpace(string(bytes))
	lines := strings.Split(data, "\n")

	programs := parseLines(lines)

	part_1(programs)
	part_2(programs)
}

func part_1(programs []program) {
	systemMemories := make(map[int64]int64)
	for _, prog := range programs {
		for _, mem := range prog.memories {
			systemMemories[mem.address] = transform(mem.value, prog.mask)
		}
	}

	var sum int64
	for _, value := range systemMemories {
		sum += value
	}

	fmt.Println(sum)
}

func part_2(programs []program) {
	systemMemories := make(map[int64]int64)
	for _, prog := range programs {
		for _, mem := range prog.memories {
			newAddresses := transform_2(mem.address, prog.mask)
			for _, address := range newAddresses {
				systemMemories[address] = mem.value
			}
		}
	}

	var sum int64
	for _, value := range systemMemories {
		sum += value
	}

	fmt.Println(sum)
}

type program struct {
	mask     string
	memories []memory
}

type memory struct {
	address int64
	value   int64
}

func parseLines(lines []string) []program {
	programs := make([]program, 0, 1000)

	for i := 0; i < len(lines); {
		if !strings.HasPrefix(lines[i], "mask") {
			i++
			continue
		}

		var mask string
		fmt.Sscanf(lines[i], "mask = %s", &mask)

		memories := make([]memory, 0, 10)
		j := i + 1
		for j := i + 1; j < len(lines); j++ {
			if !strings.HasPrefix(lines[j], "mem") {
				break
			}

			var address, value int64
			fmt.Sscanf(lines[j], "mem[%d] = %d", &address, &value)
			mem := memory{
				address: address,
				value:   value,
			}
			memories = append(memories, mem)
		}
		j--
		if j == i {
			i++
		} else {
			i = j
		}

		prog := program{
			mask:     mask,
			memories: memories,
		}
		programs = append(programs, prog)
	}

	return programs
}

func transform(v int64, mask string) int64 {
	bits := int2bits(v)

	if len(bits) != len(mask) {
		log.Fatal("something wrong, oh god")
	}

	newBits := ""

	for i, c := range mask {
		if string(c) == "X" {
			newBits += string(bits[i])
			continue
		}

		newBits += string(mask[i])
	}

	return bits2int(newBits)
}

func transform_2(v int64, mask string) []int64 {
	bits := int2bits(v)

	if len(bits) != len(mask) {
		log.Fatal("something wrong, oh god")
	}

	newManyBits := make([]string, 0, 10)
	newManyBits = append(newManyBits, "")

	for i, c := range mask {
		switch string(c) {
		case "0":
			for j := range newManyBits {
				newManyBits[j] += string(bits[i])
			}
		case "1":
			for j := range newManyBits {
				newManyBits[j] += "1"
			}
		case "X":
			tempManyBits := make([]string, len(newManyBits))
			copy(tempManyBits, newManyBits)

			for j := range newManyBits {
				newManyBits[j] += "0"
			}

			for j := range tempManyBits {
				tempManyBits[j] += "1"
			}

			newManyBits = append(newManyBits, tempManyBits...)
		}
	}

	result := make([]int64, len(newManyBits))
	for i := range newManyBits {
		result[i] = bits2int(newManyBits[i])
	}
	return result
}

func int2bits(v int64) string {
	result := ""

	for i := 0; i < 36; i++ {
		b := v % 2
		v /= 2
		result = strconv.FormatInt(b, 10) + result
	}

	return result
}

func bits2int(bits string) int64 {
	var result int64

	for i := 0; i < 36; i++ {
		b, err := strconv.ParseInt(string(bits[i]), 10, 64)
		if err != nil {
			log.Fatal("help me please")
		}

		result = result*2 + b
	}

	return result
}
