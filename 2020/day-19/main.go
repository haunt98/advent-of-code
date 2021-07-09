package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/haunt98/advent-of-code/pkg/debug"
	"github.com/haunt98/advent-of-code/pkg/parser"
)

func main() {
	fmt.Println("2020/day-19")

	pars, err := parser.ParseParagraphs("2020/day-19/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	ruleLines := strings.Split(pars[0], "\n")
	if debug.IsDebug() {
		fmt.Println(ruleLines)
	}
	msgsLines := strings.Split(pars[1], "\n")
	if debug.IsDebug() {
		fmt.Println(msgsLines)
	}

	part_1(ruleLines, msgsLines)
	part_2(ruleLines, msgsLines)
}

func part_1(ruleLines, msgsLines []string) {
	rules := make(map[int]rule)
	for _, line := range ruleLines {
		r := parseRule(line)
		rules[r.id] = r
	}

	// count all line match rule id 0
	count := 0
	for _, line := range msgsLines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		indexes := getValidIndexes(line, 0, rules)
		if len(indexes) == 0 {
			continue
		}

		for _, index := range indexes {
			if index == len(line) {
				count++
				break
			}
		}
	}
	fmt.Println(count)
}

func part_2(ruleLines, msgsLines []string) {
	rules := make(map[int]rule)
	for _, line := range ruleLines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		r := parseRule(line)
		rules[r.id] = r
	}

	rules[8] = rule{
		id:   8,
		kind: kindComplex,
		orRuleIDs: [][]int{
			{42},
			{42, 8},
		},
	}
	rules[11] = rule{
		id:   11,
		kind: kindComplex,
		orRuleIDs: [][]int{
			{42, 31},
			{42, 11, 31},
		},
	}

	// count all line match rule id 0
	count := 0
	for _, line := range msgsLines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		indexes := getValidIndexes(line, 0, rules)
		if len(indexes) == 0 {
			continue
		}

		for _, index := range indexes {
			if index == len(line) {
				count++
				break
			}
		}
	}
	fmt.Println(count)
}

const (
	kindSimple = iota + 1
	kindComplex
)

type rule struct {
	id   int
	kind int
	// literal string
	value string
	//  rules or rules or rules
	orRuleIDs [][]int
}

func parseRule(line string) rule {
	idRest := strings.Split(line, ":")
	id, err := strconv.Atoi(strings.TrimSpace(idRest[0]))
	if err != nil {
		log.Fatal(err)
	}

	rest := strings.TrimSpace(idRest[1])

	// simple
	if strings.Contains(rest, `"`) {
		rest = strings.TrimPrefix(rest, `"`)
		rest = strings.TrimSuffix(rest, `"`)
		rest = strings.TrimSpace(rest)
		return rule{
			id:    id,
			kind:  kindSimple,
			value: rest,
		}
	}

	// complex
	orRuleIDs := make([][]int, 0, 10)

	rawOrRuleIDs := strings.Split(rest, "|")
	for _, rawRuleIDs := range rawOrRuleIDs {
		rawRuleIDs = strings.TrimSpace(rawRuleIDs)
		if rawRuleIDs == "" {
			continue
		}

		ruleIDs := strings.Split(rawRuleIDs, " ")
		realRuleIDs := make([]int, 0, 10)
		for _, ruleID := range ruleIDs {
			ruleID = strings.TrimSpace(ruleID)
			if ruleID == "" {
				continue
			}

			realRuleID, err := strconv.Atoi(ruleID)
			if err != nil {
				log.Fatal(err)
			}

			realRuleIDs = append(realRuleIDs, realRuleID)
		}

		orRuleIDs = append(orRuleIDs, realRuleIDs)
	}

	return rule{
		id:        id,
		kind:      kindComplex,
		orRuleIDs: orRuleIDs,
	}
}

// return indexes where msg[0:index] is matched with ruleID
// get all valid indexes, do not stop on first ok (handle loop rule)
// example 1: 2 3 | 3 2
// if 2 3 is matched, do not stop, check 3 2 too
func getValidIndexes(msg string, ruleID int, rules map[int]rule) []int {
	r, ok := rules[ruleID]
	if !ok {
		return nil
	}

	if r.kind == kindSimple {
		if strings.HasPrefix(msg, r.value) {
			return []int{1}
		}

		return nil
	}

	result := make([]int, 0, 10)

	for _, ruleIDs := range r.orRuleIDs {
		indexes := []int{0}

		for _, ruleID := range ruleIDs {
			temp := make([]int, 0, 10)
			for _, index := range indexes {
				if index >= len(msg) {
					continue
				}

				newValidIndexes := getValidIndexes(msg[index:], ruleID, rules)
				if len(newValidIndexes) == 0 {
					continue
				}

				for _, newValidIndex := range newValidIndexes {
					temp = append(temp, index+newValidIndex)
				}
			}
			indexes = temp
			if len(indexes) == 0 {
				break
			}
		}

		if len(indexes) == 0 {
			continue
		}

		result = append(result, indexes...)
	}

	if len(result) == 0 {
		return nil
	}

	return result
}
