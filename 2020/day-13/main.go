package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("2020/day-13")

	bytes, err := ioutil.ReadFile("2020/day-13/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	timestamp, err := strconv.ParseInt(lines[0], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	rawBuses := strings.Split(strings.TrimSpace(lines[1]), ",")

	part_1(timestamp, rawBuses)
	part_2(rawBuses)
}

func part_1(timestamp int64, rawBuses []string) {
	buses := make([]int64, 0, 1000)

	for _, rawBus := range rawBuses {
		rawBus = strings.TrimSpace(rawBus)
		if rawBus == "" || rawBus == "x" {
			continue
		}

		bus, err := strconv.ParseInt(rawBus, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		buses = append(buses, bus)
	}

	var minBus int64
	var minArrive int64
	var isMinInit bool

	for _, bus := range buses {
		arrive := timestamp + (bus - timestamp%bus)
		if !isMinInit {
			minBus = bus
			minArrive = arrive
			isMinInit = true
			continue
		}

		if minArrive > arrive {
			minBus = bus
			minArrive = arrive
		}
	}

	fmt.Println((minArrive - timestamp) * minBus)
}

func part_2(rawBuses []string) {
	buses := parseBuses_2(rawBuses)
	modulos := parseModulos_2(buses)
	onlyBuses := parseOnlyBuses_2(buses)

	// https://en.wikipedia.org/wiki/Chinese_remainder_theorem#Search_by_sieving
	result := onlyBuses[0] + modulos[onlyBuses[0]]
	step := onlyBuses[0]

	for i := 1; i < len(onlyBuses); i++ {
		for {
			if result%onlyBuses[i] == modulos[onlyBuses[i]] {
				break
			}

			result += step
		}

		step *= onlyBuses[i]
	}

	fmt.Println(result)
}

func isDebug() bool {
	return os.Getenv("DEBUG") != ""
}

func parseBuses_2(rawBuses []string) map[int64]int64 {
	buses := make(map[int64]int64)

	for i, rawBus := range rawBuses {
		if rawBus == "x" {
			continue
		}

		bus, err := strconv.ParseInt(rawBus, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		buses[bus] = int64(i)
	}

	if isDebug() {
		fmt.Println(buses)
	}

	return buses
}

func parseModulos_2(buses map[int64]int64) map[int64]int64 {
	modulos := make(map[int64]int64)

	for bus, i := range buses {
		modulo := int64(-i)
		for modulo < 0 {
			modulo += bus
		}

		modulos[bus] = modulo % bus
	}
	if isDebug() {
		fmt.Println(modulos)
	}

	return modulos
}

func parseOnlyBuses_2(buses map[int64]int64) []int64 {
	onlyBuses := make([]int64, 0, 1000)

	for bus := range buses {
		onlyBuses = append(onlyBuses, bus)
	}

	sort.Slice(onlyBuses, func(i, j int) bool {
		return onlyBuses[i] > onlyBuses[j]
	})

	return onlyBuses
}
