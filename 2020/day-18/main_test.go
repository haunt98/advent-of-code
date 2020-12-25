package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRPN(t *testing.T) {
	tests := []struct {
		name string
		line string
		want []string
	}{
		{
			name: "example",
			line: "1 + 2 * 3 + 4 * 5 + 6",
			want: []string{"1", "2", "+", "3", "*", "4", "+", "5", "*", "6", "+"},
		},
		{
			name: "example",
			line: "1 + (2 * 3) + (4 * (5 + 6))",
			want: []string{"1", "2", "3", "*", "+", "4", "5", "6", "+", "*", "+"},
		},
	}

	for _, tc := range tests {
		expr := parseExpression(tc.line)
		got := parseRPN(expr, getPrecedence1)
		assert.Equal(t, tc.want, got)
	}
}

func TestCalculate(t *testing.T) {
	tests := []struct {
		name string
		line string
		want int
	}{
		{
			name: "example",
			line: "1 + 2 * 3 + 4 * 5 + 6",
			want: 71,
		},
		{
			name: "example",
			line: "1 + (2 * 3) + (4 * (5 + 6))",
			want: 51,
		},
		{
			name: "example",
			line: "5 + (8 * 3 + 9 + 3 * 4 * 3)",
			want: 437,
		},
		{
			name: "example",
			line: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			want: 12240,
		},
		{
			name: "example",
			line: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			want: 13632,
		},
	}

	for _, tc := range tests {
		expr := parseExpression(tc.line)
		rpn := parseRPN(expr, getPrecedence1)
		got := calculateRPN(rpn)
		assert.Equal(t, tc.want, got)
	}
}

func TestParseRPN2(t *testing.T) {
	tests := []struct {
		name string
		line string
		want []string
	}{
		{
			name: "example",
			line: "1 + 2 * 3 + 4 * 5 + 6",
			want: []string{"1", "2", "+", "3", "4", "+", "*", "5", "6", "+", "*"},
		},
	}

	for _, tc := range tests {
		expr := parseExpression(tc.line)
		got := parseRPN(expr, getPrecedence2)
		assert.Equal(t, tc.want, got)
	}
}

func TestCalculate2(t *testing.T) {
	tests := []struct {
		name string
		line string
		want int
	}{
		{
			name: "example",
			line: "1 + 2 * 3 + 4 * 5 + 6",
			want: 231,
		},
		{
			name: "example",
			line: "1 + (2 * 3) + (4 * (5 + 6))",
			want: 51,
		},
		{
			name: "example",
			line: "5 + (8 * 3 + 9 + 3 * 4 * 3)",
			want: 1445,
		},
		{
			name: "example",
			line: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			want: 669060,
		},
		{
			name: "example",
			line: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			want: 23340,
		},
	}

	for _, tc := range tests {
		expr := parseExpression(tc.line)
		rpn := parseRPN(expr, getPrecedence2)
		got := calculateRPN(rpn)
		assert.Equal(t, tc.want, got)
	}
}
