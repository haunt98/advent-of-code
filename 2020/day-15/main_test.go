package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart_1(t *testing.T) {
	tests := []struct {
		name  string
		inits []int
		want  int
	}{
		{
			name:  "example",
			inits: []int{0, 3, 6},
			want:  436,
		},
		{
			name:  "example",
			inits: []int{1, 3, 2},
			want:  1,
		},
		{
			name:  "example",
			inits: []int{2, 1, 3},
			want:  10,
		},
		{
			name:  "example",
			inits: []int{1, 2, 3},
			want:  27,
		},
		{
			name:  "example",
			inits: []int{2, 3, 1},
			want:  78,
		},
		{
			name:  "example",
			inits: []int{3, 2, 1},
			want:  438,
		},
		{
			name:  "example",
			inits: []int{3, 1, 2},
			want:  1836,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := part_1(tc.inits)
			assert.Equal(t, tc.want, got)
		})
	}
}
