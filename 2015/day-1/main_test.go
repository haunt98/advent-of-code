package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoFloor(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "(())",
			s:    "(())",
			want: 0,
		},
		{
			name: "()()",
			s:    "()()",
			want: 0,
		},
		{
			name: "(((",
			s:    "(((",
			want: 3,
		},
		{
			name: "(()(()(",
			s:    "(()(()(",
			want: 3,
		},
		{
			name: "))(((((",
			s:    "))(((((",
			want: 3,
		},
		{
			name: "())",
			s:    "())",
			want: -1,
		},
		{
			name: ")))",
			s:    ")))",
			want: -3,
		},
		{
			name: ")())())",
			s:    ")())())",
			want: -3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := goFloor(tc.s)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestGoIndexFloorUntil(t *testing.T) {
	tests := []struct {
		name      string
		s         string
		stopFloor int
		want      int
	}{
		{
			name:      ")",
			s:         ")",
			stopFloor: -1,
			want:      1,
		},
		{
			name:      "()())",
			s:         "()())",
			stopFloor: -1,
			want:      5,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := goIndexFloorUntil(tc.s, tc.stopFloor)
			assert.Equal(t, tc.want, got)
		})
	}
}
