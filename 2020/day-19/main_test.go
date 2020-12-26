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

func TestComposeMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]string
		want   []string
	}{
		{
			name: "example",
			matrix: [][]string{
				{"a", "b"},
			},
			want: []string{"a", "b"},
		},
		{
			name: "example",
			matrix: [][]string{
				{""},
				{"c", "d"},
			},
			want: []string{"c", "d"},
		},
		{
			name: "example",
			matrix: [][]string{
				{"a", "b"},
				{"c", "d"},
			},
			want: []string{"ac", "ad", "bc", "bd"},
		},
		{
			name: "example",
			matrix: [][]string{
				{"a", "b"},
				{"c", "d"},
				{"1", "2"},
			},
			want: []string{"ac1", "ac2", "ad1", "ad2", "bc1", "bc2", "bd1", "bd2"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := composeMatrix(tc.matrix)
			assert.Equal(t, tc.want, got)
		})
	}
}
