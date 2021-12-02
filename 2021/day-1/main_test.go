package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountIncrease(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			name: "example",
			arr: []int{
				199,
				200,
				208,
				210,
				200,
				207,
				240,
				269,
				260,
				263,
			},
			want: 7,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := countIncrease(tc.arr)

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestCountIncrease3Windows(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			name: "example",
			arr: []int{
				199,
				200,
				208,
				210,
				200,
				207,
				240,
				269,
				260,
				263,
			},
			want: 5,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := countIncrease3Windows(tc.arr)

			assert.Equal(t, tc.want, got)
		})
	}
}
