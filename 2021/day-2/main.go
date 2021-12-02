package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/make-go-great/panda-go/parser"
)

func main() {
	fmt.Println("2021/day-2")

	lines, err := parser.ParseLines("2021/day-2/input.txt")
	if err != nil {
		log.Fatalln("Failed to parse int lines", err)
	}

	cmds := parseCommands(lines)

	fmt.Println("part 1:", part_1(cmds))
	fmt.Println("part 2:", part_2(cmds))
}

type direction string

const (
	forward direction = "forward"
	down    direction = "down"
	up      direction = "up"
)

type command struct {
	unit      int
	direction direction
}

type position struct {
	horizontal, depth int
}

type complicatedPosition struct {
	position

	aim int
}

func parseCommand(line string) command {
	cmd := command{}
	fmt.Sscanf(line, "%s %d", &cmd.direction, &cmd.unit)
	return cmd
}

func parseCommands(lines []string) []command {
	cmds := make([]command, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		cmd := parseCommand(line)
		cmds = append(cmds, cmd)
	}
	return cmds
}

func part_1(cmds []command) int {
	start := position{}
	end := pilot(start, cmds)
	return end.depth * end.horizontal
}

func part_2(cmds []command) int {
	start := complicatedPosition{}
	end := compilcatedPilot(start, cmds)
	return end.depth * end.horizontal
}

func pilot(start position, cmds []command) position {
	for _, cmd := range cmds {
		if cmd.direction == forward {
			start.horizontal += cmd.unit
		} else if cmd.direction == down {
			start.depth += cmd.unit
		} else if cmd.direction == up {
			start.depth -= cmd.unit
		}
	}
	return start
}

func compilcatedPilot(start complicatedPosition, cmds []command) complicatedPosition {
	for _, cmd := range cmds {
		if cmd.direction == down {
			start.aim += cmd.unit
		} else if cmd.direction == up {
			start.aim -= cmd.unit
		} else if cmd.direction == forward {
			start.horizontal += cmd.unit
			start.depth += start.aim * cmd.unit
		}
	}
	return start
}
