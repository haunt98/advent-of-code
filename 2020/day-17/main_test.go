package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChange_1(t *testing.T) {
	lines := []string{
		".#.",
		"..#",
		"###",
	}

	m3d := parse3D(lines)
	wantM3d := map[int]map[int]map[int]bool{
		0: map[int]map[int]bool{
			0: map[int]bool{
				0: false,
				1: true,
				2: false,
			},
			1: map[int]bool{
				0: false,
				1: false,
				2: true,
			},
			2: map[int]bool{
				0: true,
				1: true,
				2: true,
			},
		},
	}
	assert.Equal(t, wantM3d, m3d)
}
