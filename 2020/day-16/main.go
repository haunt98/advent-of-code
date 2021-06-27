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
		fmt.Printf("rules: %+v\n", rules)
	}

	yourTicket := parseYourTicket(linesYourTicket)
	if isDebug() {
		fmt.Printf("your ticket: %+v\n", yourTicket)
	}

	nearbyTickets := parseNearbyTickets(linesNearbyTickets)
	if isDebug() {
		fmt.Printf("nearby tickets: %+v\n", nearbyTickets)
	}

	part_1(rules, nearbyTickets)
	part_2(rules, yourTicket, nearbyTickets)
}

func part_1(rules []rule, nearbyTickets []ticket) {
	sum := 0
	for _, t := range nearbyTickets {
		ok, errorRate := isValidTicket(t, rules)
		if !ok {
			sum += errorRate
		}
	}
	fmt.Println(sum)
}

func isValidTicket(t ticket, rules []rule) (ok bool, errorRate int) {
	// if 1 field not satisfy any rule
	// then ticket is not valid
	for _, field := range t.data {
		validField := false

		for _, r := range rules {
			if isValidField(field, r) {
				validField = true
				break
			}
		}

		if !validField {
			return false, field
		}
	}

	return true, 0
}

func isValidField(field int, r rule) bool {
	return (r.rangeRules[0].from <= field && field <= r.rangeRules[0].to) ||
		(r.rangeRules[1].from <= field && field <= r.rangeRules[1].to)
}

func part_2(rules []rule, yourTicket ticket, nearbyTickets []ticket) {
	nearbyTickets = filterNearbyTickets(rules, nearbyTickets)
	if isDebug() {
		fmt.Printf("new nearby tickets: %+v\n", nearbyTickets)
	}

	possibleRules := getPossibleRules(rules, nearbyTickets)
	if isDebug() {
		fmt.Println("possile rules")
		for i, rules := range possibleRules {
			fmt.Printf("field %d\n", i)
			for name := range rules {
				fmt.Printf("[%s] ", name)
			}
			fmt.Println()
		}
	}

	solution := make([]rule, len(rules))
	solution = csp(rules, possibleRules, solution, 0)
	if isDebug() {
		fmt.Printf("solution: %+v\n", solution)
		for i := range solution {
			fmt.Printf("[%s] ", solution[i].name)
		}
		fmt.Println()
	}

	if len(solution) != len(rules) {
		fmt.Println("mystery unsolved")
		return
	}

	result := 1
	for i := range yourTicket.data {
		if strings.HasPrefix(strings.TrimSpace(solution[i].name), "departure") {
			if isDebug() {
				fmt.Println(solution[i].name, yourTicket.data[i])
			}
			result *= yourTicket.data[i]
		}
	}
	fmt.Println(result)
}

// go from field with least possible rule
func csp(rules []rule, possibleRules []map[string]rule, prevSolution []rule, step int) []rule {
	if step == len(rules) {
		return prevSolution
	}

	if isDebug() {
		fmt.Println("step", step)
		fmt.Println("possibleRules", possibleRules)
		fmt.Println("prevSolution", prevSolution)
	}

	var min_i int
	var min_set bool
	for i := range possibleRules {
		// already in solution
		if len(possibleRules[i]) == 0 {
			continue
		}

		if !min_set {
			min_i = i
			min_set = true
			continue
		}

		if len(possibleRules[i]) < len(possibleRules[min_i]) {
			min_i = i
		}
	}

	if isDebug() {
		fmt.Println("min_set", min_set, "min_i", min_i)
	}

	if !min_set {
		return nil
	}

	for name, r := range possibleRules[min_i] {
		newPrevSolution := make([]rule, len(rules))
		copy(newPrevSolution, prevSolution)
		newPrevSolution[min_i] = r

		newPossibleRules := make([]map[string]rule, len(rules))
		copyRules(newPossibleRules, possibleRules)
		// will set in solution
		newPossibleRules[min_i] = make(map[string]rule)

		// delete r from other possible rules
		for i := range newPossibleRules {
			if len(newPossibleRules[i]) == 0 {
				continue
			}

			delete(newPossibleRules[i], name)
		}

		if isDebug() {
			fmt.Println("prevSolution", prevSolution)
			fmt.Println("newPossibleRules", newPossibleRules)
		}

		solution := csp(rules, newPossibleRules, newPrevSolution, step+1)
		if solution != nil {
			return solution
		}
	}

	return nil
}

func copyRules(dst []map[string]rule, src []map[string]rule) {
	if len(dst) != len(src) {
		log.Fatal("keep it up soldier")
	}

	for i, srcMap := range src {
		dst[i] = make(map[string]rule)
		for k, v := range srcMap {
			dst[i][k] = v
		}
	}
}

func filterNearbyTickets(rules []rule, nearbyTickets []ticket) []ticket {
	newNearbyTickets := make([]ticket, 0, len(nearbyTickets))

	for _, t := range nearbyTickets {
		if ok, _ := isValidTicket(t, rules); !ok {
			continue
		}

		newNearbyTickets = append(newNearbyTickets, t)
	}

	return newNearbyTickets
}

func getPossibleRules(rules []rule, nearbyTickets []ticket) []map[string]rule {
	lenRules := len(rules)

	possibleRules := make([]map[string]rule, lenRules)

	for i := 0; i < lenRules; i++ {
		possibleRules[i] = make(map[string]rule)
		// rule must satisfy all data[i] of nearby tickets
		for _, r := range rules {
			satisfy := true

			for _, t := range nearbyTickets {
				if !isValidField(t.data[i], r) {
					satisfy = false
					break
				}
			}

			if satisfy {
				possibleRules[i][r.name] = r
			}
		}
	}

	return possibleRules
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
