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
	fmt.Println("2020/day-13")

	bytes, err := ioutil.ReadFile("2020/day-13/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	timestamp, err := strconv.ParseUint(lines[0], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	rawBuses := strings.Split(strings.TrimSpace(lines[1]), ",")

	part_1(timestamp, rawBuses)
	part_2(timestamp, rawBuses)
}

func part_1(timestamp uint64, rawBuses []string) {
	buses := make([]uint64, 0, 1000)

	for _, rawBus := range rawBuses {
		rawBus = strings.TrimSpace(rawBus)
		if rawBus == "" || rawBus == "x" {
			continue
		}

		bus, err := strconv.ParseUint(rawBus, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		buses = append(buses, bus)
	}

	var minBus uint64
	var minArrive uint64
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

func part_2(timestamp uint64, rawBuses []string) {
	buses := make(map[uint64]uint64)

	for i, rawBus := range rawBuses {
		if rawBus == "x" {
			continue
		}

		bus, err := strconv.ParseUint(rawBus, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		buses[bus] = uint64(i)
	}

	if isDebug() {
		fmt.Println(buses)
	}

	var maxBus uint64

	for bus, _ := range buses {
		if maxBus < bus {
			maxBus = bus
		}
	}

	result := (100000000000000/maxBus)*maxBus + maxBus - buses[maxBus]

	for {
		if isDebug() {
			fmt.Println(result)
		}

		ok := true

		for bus, t := range buses {
			if (result+t)%bus != 0 {
				ok = false
				break
			}
		}

		if ok {
			break
		}

		result += maxBus
	}

	fmt.Println(result)
}

func isDebug() bool {
	return os.Getenv("DEBUG") != ""
}
