package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCommand(t *testing.T) {
	tests := []struct {
		name string
		line string
		want command
	}{
		{
			name: "forward",
			line: "forward 1",
			want: command{
				unit:      1,
				direction: forward,
			},
		},
		{
			name: "up",
			line: "up 2",
			want: command{
				unit:      2,
				direction: up,
			},
		},
		{
			name: "down",
			line: "down 3",
			want: command{
				unit:      3,
				direction: down,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := parseCommand(tc.line)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPilot(t *testing.T) {
	tests := []struct {
		name  string
		start position
		cmds  []command
		want  position
	}{
		{
			name:  "example",
			start: position{},
			cmds: []command{
				{
					unit:      5,
					direction: forward,
				},
				{
					unit:      5,
					direction: down,
				},
				{
					unit:      8,
					direction: forward,
				},
				{
					unit:      3,
					direction: up,
				},
				{
					unit:      8,
					direction: down,
				},
				{
					unit:      2,
					direction: forward,
				},
			},
			want: position{
				horizontal: 15,
				depth:      10,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := pilot(tc.start, tc.cmds)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestComplicatedPilot(t *testing.T) {
	tests := []struct {
		name  string
		start complicatedPosition
		cmds  []command
		want  complicatedPosition
	}{
		{
			name:  "example",
			start: complicatedPosition{},
			cmds: []command{
				{
					unit:      5,
					direction: forward,
				},
				{
					unit:      5,
					direction: down,
				},
				{
					unit:      8,
					direction: forward,
				},
				{
					unit:      3,
					direction: up,
				},
				{
					unit:      8,
					direction: down,
				},
				{
					unit:      2,
					direction: forward,
				},
			},
			want: complicatedPosition{
				position: position{
					horizontal: 15,
					depth:      60,
				},
				aim: 10,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := compilcatedPilot(tc.start, tc.cmds)
			assert.Equal(t, tc.want, got)
		})
	}
}
