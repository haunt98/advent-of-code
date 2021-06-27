package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountAtLeastOnePresent(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: ">",
			s:    ">",
			want: 2,
		},
		{
			name: "^>v<",
			s:    "^>v<",
			want: 4,
		},
		{
			name: "^v^v^v^v^v",
			s:    "^v^v^v^v^v",
			want: 2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := countAtLeastOnePresent(tc.s)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestCountAtLeastOnePresentButTakeTurn(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "^v",
			s:    "^v",
			want: 3,
		},
		{
			name: "^>v<",
			s:    "^>v<",
			want: 3,
		},
		{
			name: "^v^v^v^v^v",
			s:    "^v^v^v^v^v",
			want: 11,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := countAtLeastOnePresentButTakeTurn(tc.s)
			assert.Equal(t, tc.want, got)
		})
	}
}
