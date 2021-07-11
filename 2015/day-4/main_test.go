package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcMD5(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "abcdef609043",
			s:    "abcdef609043",
			want: "000001dbbfa3a5c83a2d506429c7b00e",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := calcMD5(tc.s)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestIsAdventCoin5(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "abcdef609043",
			s:    "abcdef609043",
			want: true,
		},
		{
			name: "pqrstuv1048970",
			s:    "pqrstuv1048970",
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isAdventCoin5(tc.s)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestCalcAdventCoin5(t *testing.T) {
	tests := []struct {
		name string
		seed string
		want int
	}{
		{
			name: "abcdef",
			seed: "abcdef",
			want: 609043,
		},
		{
			name: "pqrstuv",
			seed: "pqrstuv",
			want: 1048970,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := calcAdventCoin(tc.seed, isAdventCoin5)
			assert.Equal(t, tc.want, got)
		})
	}
}
