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
	fmt.Println("2020/day-16/main.go")

	bytes, err := ioutil.ReadFile("2020/day-16/input.txt")
	if err != nil {
		log.Fatal("god help your input")
	}

	bigLines := strings.Split(strings.TrimSpace(string(bytes)), "\n\n")
	if len(bigLines) != 3 {
		log.Fatal("god help your big lines")
	}
	linesRule := strings.Split(bigLines[0], "\n")
	linesYourTicket := strings.Split(bigLines[1], "\n")
	linesNearbyTickets := strings.Split(bigLines[2], "\n")

	rules := parseRules(linesRule)
	if isDebug() {
		fmt.Printf("%+v\n", rules)
	}

	yourTicket := parseYourTicket(linesYourTicket)
	if isDebug() {
		fmt.Printf("%+v\n", yourTicket)
	}

	nearbyTickets := parseNearbyTickets(linesNearbyTickets)
	if isDebug() {
		fmt.Printf("%+v\n", nearbyTickets)
	}

	part_1(rules, nearbyTickets)
}

func part_1(rules []rule, nearbyTickets []ticket) {
	sum := 0
	for _, t := range nearbyTickets {
		for _, field := range t.data {
			completelyInvalid := true

			for _, r := range rules {
				if len(r.rangeRules) != 2 {
					log.Fatal("wrong rules")
				}

				if r.rangeRules[0].from <= field &&
					r.rangeRules[0].to >= field {
					completelyInvalid = false
					break
				}

				if r.rangeRules[1].from <= field &&
					r.rangeRules[1].to >= field {
					completelyInvalid = false
					break
				}
			}

			if completelyInvalid {
				sum += field
			}
		}
	}
	fmt.Println(sum)
}

func part_2(rules []rule, yourTicket ticket, nearbyTickets []ticket) {
	nearbyTickets = filerNearbyTickets(rules, nearbyTickets)
}

func filerNearbyTickets(rules []rule, nearbyTickets []ticket) []ticket {
	newNearbyTickets := make([]ticket, 0, 1000)

	for _, t := range nearbyTickets {
		for _, field := range t.data {
			completelyInvalid := true

			for _, r := range rules {
				if len(r.rangeRules) != 2 {
					log.Fatal("wrong rules")
				}

				if r.rangeRules[0].from <= field &&
					r.rangeRules[0].to >= field {
					completelyInvalid = false
					break
				}

				if r.rangeRules[1].from <= field &&
					r.rangeRules[1].to >= field {
					completelyInvalid = false
					break
				}
			}

			if !completelyInvalid {
				newNearbyTickets = append(newNearbyTickets, t)
			}
		}
	}

	return newNearbyTickets
}

type rule struct {
	name       string
	rangeRules []rangeRule
}

type rangeRule struct {
	from, to int
}

func parseRules(lines []string) []rule {
	rules := make([]rule, 0, 1000)

	for _, line := range lines {
		r := parseRule(line)
		rules = append(rules, r)
	}

	return rules
}

func parseRule(line string) rule {
	nameAndRanges := strings.Split(line, ":")
	name := strings.TrimSpace(nameAndRanges[0])

	rawRanges := strings.Split(strings.TrimSpace(nameAndRanges[1]), "or")
	raRules := make([]rangeRule, 0, 10)
	for _, rawRange := range rawRanges {
		var from, to int
		fmt.Sscanf(strings.TrimSpace(rawRange), "%d-%d", &from, &to)
		raRule := rangeRule{
			from: from,
			to:   to,
		}
		raRules = append(raRules, raRule)
	}

	r := rule{
		name:       name,
		rangeRules: raRules,
	}

	return r
}

type ticket struct {
	data []int
}

func parseYourTicket(lines []string) ticket {
	if len(lines) != 2 {
		log.Fatal("god hep your ticket")
	}

	return parseTicket(lines[1])
}

func parseNearbyTickets(lines []string) []ticket {
	if len(lines) < 2 {
		log.Fatal("god hep your nearby ticket")
	}

	tickets := make([]ticket, 0, 100)
	for i := 1; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}
		t := parseTicket(line)
		tickets = append(tickets, t)
	}
	return tickets
}

func parseTicket(line string) ticket {
	rawData := strings.Split(strings.TrimSpace(line), ",")
	data := make([]int, 0, 10)
	for i := range rawData {
		rawData[i] = strings.TrimSpace(rawData[i])
		if rawData[i] == "" {
			continue
		}

		singleData, err := strconv.Atoi(rawData[i])
		if err != nil {
			log.Fatal("god hep your parse ticket")
		}
		data = append(data, singleData)
	}
	return ticket{
		data: data,
	}
}

func isDebug() bool {
	return os.Getenv("DEBUG") != ""
}
