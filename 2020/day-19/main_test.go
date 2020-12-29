package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRule(t *testing.T) {
	tests := []struct {
		name string
		line string
		want rule
	}{
		{
			name: "simple",
			line: `3: "b"`,
			want: rule{
				id:    3,
				kind:  kindSimple,
				value: "b",
			},
		},
		{
			name: "complex",
			line: "0: 1 2",
			want: rule{
				id:   0,
				kind: kindComplex,
				orRuleIDs: [][]int{
					{1, 2},
				},
			},
		},
		{
			name: "complex",
			line: "2: 1 3 | 3 1",
			want: rule{
				id:   2,
				kind: kindComplex,
				orRuleIDs: [][]int{
					{1, 3},
					{3, 1},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := parseRule(tc.line)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestGetValidIndexes(t *testing.T) {
	ruleLines := []string{
		"0: 4 1 5",
		"1: 2 3 | 3 2",
		"2: 4 4 | 5 5",
		"3: 4 5 | 5 4",
		`4: "a"`,
		`5: "b"`,
	}
	rules := make(map[int]rule)

	for _, line := range ruleLines {
		r := parseRule(line)
		rules[r.id] = r
	}

	tests := []struct {
		name   string
		msg    string
		ruleID int
		rules  map[int]rule
		want   []int
	}{
		{
			name:   "rule 5",
			msg:    "b",
			ruleID: 5,
			rules:  rules,
			want:   []int{1},
		},
		{
			name:   "rule 5",
			msg:    "ba",
			ruleID: 5,
			rules:  rules,
			want:   []int{1},
		},
		{
			name:   "rule 5",
			msg:    "bb",
			ruleID: 5,
			rules:  rules,
			want:   []int{1},
		},
		{
			name:   "rule 5",
			msg:    "a",
			ruleID: 5,
			rules:  rules,
			want:   nil,
		},
		{
			name:   "rule 5",
			msg:    "ab",
			ruleID: 5,
			rules:  rules,
			want:   nil,
		},
		{
			name:   "rule 3",
			msg:    "ab",
			ruleID: 3,
			rules:  rules,
			want:   []int{2},
		},
		{
			name:   "rule 3",
			msg:    "ba",
			ruleID: 3,
			rules:  rules,
			want:   []int{2},
		},
		{
			name:   "rule 3",
			msg:    "aa",
			ruleID: 3,
			rules:  rules,
			want:   nil,
		},
		{
			name:   "rule 3",
			msg:    "bb",
			ruleID: 3,
			rules:  rules,
			want:   nil,
		},
		{
			name:   "rule 3 with longer msg",
			msg:    "abc",
			ruleID: 3,
			rules:  rules,
			want:   []int{2},
		},
		{
			name:   "rule 3 with longer msg",
			msg:    "bac",
			ruleID: 3,
			rules:  rules,
			want:   []int{2},
		},
		{
			name:   "rule 1",
			msg:    "abaa",
			ruleID: 1,
			rules:  rules,
			want:   []int{4},
		},
		{
			name:   "rule 1",
			msg:    "abbb",
			ruleID: 1,
			rules:  rules,
			want:   []int{4},
		},
		{
			name:   "rule 1",
			msg:    "aba",
			ruleID: 1,
			rules:  rules,
			want:   nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := getValidIndexes(tc.msg, tc.ruleID, tc.rules)
			assert.Equal(t, tc.want, got)
		})
	}
}
