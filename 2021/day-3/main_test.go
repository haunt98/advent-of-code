package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCommonBits(t *testing.T) {
	tests := []struct {
		name      string
		lines     []string
		wantMost  int64
		wantLeast int64
	}{
		{
			name: "example",
			lines: []string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			},
			wantMost:  22,
			wantLeast: 9,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotMost, gotLeast := getCommonBits(tc.lines)
			assert.Equal(t, tc.wantMost, gotMost)
			assert.Equal(t, tc.wantLeast, gotLeast)
		})
	}
}

func TestGetCommonBitsThenFilter(t *testing.T) {
	tests := []struct {
		name      string
		lines     []string
		wantMost  int64
		wantLeast int64
	}{
		{
			name: "example",
			lines: []string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			},
			wantMost:  23,
			wantLeast: 10,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotMost, gotLeast := getCommonBitsThenFilter(tc.lines)
			assert.Equal(t, tc.wantMost, gotMost)
			assert.Equal(t, tc.wantLeast, gotLeast)
		})
	}
}

func TestGetCommonBitsAtPosition(t *testing.T) {
	tests := []struct {
		name           string
		lines          []string
		position       int
		wantMostLines  []string
		wantLeastLines []string
	}{
		{
			name: "example",
			lines: []string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			},
			position:       0,
			wantMostLines:  []string{"11110", "10110", "10111", "10101", "11100", "10000", "11001"},
			wantLeastLines: []string{"00100", "01111", "00111", "00010", "01010"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotMostLines, gotLeastLines := getCommonBitsAtPosition(tc.lines, tc.position)
			assert.Equal(t, tc.wantMostLines, gotMostLines)
			assert.Equal(t, tc.wantLeastLines, gotLeastLines)
		})
	}
}
