package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWrappingPaper(t *testing.T) {
	tests := []struct {
		name string
		l    int
		w    int
		h    int
		want int
	}{
		{
			name: "2x3x4",
			l:    2,
			w:    3,
			h:    4,
			want: 58,
		},
		{
			name: "1x1x10",
			l:    1,
			w:    1,
			h:    10,
			want: 43,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := getWrappingPaper(newGift(tc.l, tc.w, tc.h))
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestGetRibbon(t *testing.T) {
	tests := []struct {
		name string
		l    int
		w    int
		h    int
		want int
	}{
		{
			name: "2x3x4",
			l:    2,
			w:    3,
			h:    4,
			want: 34,
		},
		{
			name: "1x1x10",
			l:    1,
			w:    1,
			h:    10,
			want: 14,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := getRibbon(newGift(tc.l, tc.w, tc.h))
			assert.Equal(t, tc.want, got)
		})
	}
}
