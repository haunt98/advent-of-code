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

	if len(pars) != 2 {
		log.Fatal("len pars must be 2")
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
}

func part_1(ruleLines, msgsLines []string) {
	rules := make(map[int]rule)
	for _, line := range ruleLines {
		r := parseRule(line)
		rules[r.id] = r
	}

	// get all or value for rule id 0
	orValues := make(map[int][]string)
	orValue0 := getOrValue(0, rules, orValues)
	maskedOrValue0 := make(map[string]struct{})
	for _, value0 := range orValue0 {
		maskedOrValue0[value0] = struct{}{}
	}

	// count all line match rule id 0
	count := 0
	for _, line := range msgsLines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if _, ok := maskedOrValue0[line]; ok {
			count++
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

func getOrValue(id int, rules map[int]rule, orValues map[int][]string) []string {
	r, ok := rules[id]
	if !ok {
		log.Fatal("no id ? why ?")
	}

	orValue, ok := orValues[id]
	if ok {
		return orValue
	}

	if r.kind == kindSimple {
		result := []string{r.value}
		orValues[id] = result
		return result
	}

	result := make([]string, 0, 10)

	for _, ruleIDs := range r.orRuleIDs {
		matrix := make([][]string, 0, 10)
		for _, id := range ruleIDs {
			orValue = getOrValue(id, rules, orValues)
			matrix = append(matrix, orValue)
		}
		result = append(result, composeMatrix(matrix)...)
	}

	return result
}

// matrix = pool + pool + pool
// pool = ["a", "b"]
// "a" + ["b", "c"] = ["ab", "ac"]
func composeMatrix(matrix [][]string) []string {
	result := []string{""}

	for _, pool := range matrix {
		temp := make([]string, 0, len(result))
		for _, v := range result {
			for _, newV := range pool {
				temp = append(temp, v+newV)
			}
		}
		result = temp
	}

	return result
}
