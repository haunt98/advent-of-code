package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("2020/day-12")

	bytes, err := ioutil.ReadFile("2020/day-12/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	part_1(lines)
	part_2(lines)
}

func part_1(lines []string) {
	instrs := parseLines(lines)

	pos := position{
		direction: 0,
		x:         0,
		y:         0,
	}

	for _, instr := range instrs {
		pos = move(pos, instr)
	}

	x := pos.x
	if x < 0 {
		x = -x
	}

	y := pos.y
	if y < 0 {
		y = -y
	}

	fmt.Println(x + y)
}

func part_2(lines []string) {
	instrs := parseLines(lines)

	pos := position{
		x: 0,
		y: 0,
	}

	waypoint := position{
		x: 10,
		y: -1,
	}

	for _, instr := range instrs {
		pos, waypoint = move_2(pos, waypoint, instr)
		if isDebug() {
			fmt.Println(pos, waypoint)
		}
	}

	x := pos.x
	if x < 0 {
		x = -x
	}

	y := pos.y
	if y < 0 {
		y = -y
	}

	fmt.Println(x + y)
}

type position struct {
	//   3
	// 2 X 0
	//   1
	// 0 facing E
	// 1 facing S
	// 2 facing W
	// 3 facing N
	direction int

	x int
	y int
}

type instruction struct {
	side string
	step int
}

func parseLines(lines []string) []instruction {
	instrs := make([]instruction, 0, 1000)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		instrs = append(instrs, parseLine(line))
	}

	return instrs
}

func parseLine(line string) instruction {
	var side string
	var step int

	fmt.Sscanf(line, "%1s%d", &side, &step)

	return instruction{
		side: side,
		step: step,
	}
}

func move(pos position, instr instruction) position {
	direction := pos.direction
	x := pos.x
	y := pos.y

	step := instr.step

	switch instr.side {
	case "N":
		x -= step
	case "S":
		x += step
	case "E":
		y += step
	case "W":
		y -= step
	case "L":
		step = step % 360
		switch step {
		case 90:
			direction = (direction + 3) % 4
		case 180:
			direction = (direction + 2) % 4
		case 270:
			direction = (direction + 1) % 4
		}
	case "R":
		step = step % 360
		switch step {
		case 90:
			direction = (direction + 1) % 4
		case 180:
			direction = (direction + 2) % 4
		case 270:
			direction = (direction + 3) % 4
		}
	case "F":
		switch direction {
		case 0:
			y += step
		case 1:
			x += step
		case 2:
			y -= step
		case 3:
			x -= step
		}
	}

	return position{
		direction: direction,
		x:         x,
		y:         y,
	}
}

func move_2(ship, waypoint position, instr instruction) (newShip, newWaypoint position) {
	newShip = ship
	newWaypoint = waypoint

	switch instr.side {
	case "N":
		newWaypoint.y -= instr.step
	case "S":
		newWaypoint.y += instr.step
	case "E":
		newWaypoint.x += instr.step
	case "W":
		newWaypoint.x -= instr.step
	case "L":
		step := instr.step % 360
		switch step {
		case 90:
			newWaypoint.x = waypoint.y
			newWaypoint.y = -waypoint.x
		case 180:
			newWaypoint.x = -waypoint.x
			newWaypoint.y = -waypoint.y
		case 270:
			newWaypoint.x = -waypoint.y
			newWaypoint.y = waypoint.x
		}
	case "R":
		step := instr.step % 360
		switch step {
		case 90:
			newWaypoint.x = -waypoint.y
			newWaypoint.y = waypoint.x
		case 180:
			newWaypoint.x = -waypoint.x
			newWaypoint.y = -waypoint.y
		case 270:
			newWaypoint.x = waypoint.y
			newWaypoint.y = -waypoint.x
		}
	case "F":
		newShip.x += newWaypoint.x * instr.step
		newShip.y += newWaypoint.y * instr.step
	}

	return
}

func isDebug() bool {
	return os.Getenv("DEBUG") != ""
}
