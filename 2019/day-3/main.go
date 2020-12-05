package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	fmt.Println("2019/day-3")

	bytes, err := ioutil.ReadFile("2019/day-3/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Fields(string(bytes))
	if len(lines) != 2 {
		return
	}

	route_1 := strings.Split(lines[0], ",")
	route_2 := strings.Split(lines[1], ",")

	m_1 := walk(route_1)
	m_2 := walk(route_2)

	fmt.Println(part_1(m_1, m_2))
	fmt.Println(part_2(m_1, m_2))
}

func part_1(m_1, m_2 map[position]int) int {
	intersections := make([]position, 0, 1000)

	for pos, _ := range m_1 {
		if _, ok := m_2[pos]; ok {
			if pos.X == 0 && pos.Y == 0 {
				continue
			}

			intersections = append(intersections, pos)
		}
	}

	root := position{
		X: 0,
		Y: 0,
	}
	min := distance(intersections[0], root)

	for i := 1; i < len(intersections); i++ {
		dis := distance(intersections[i], root)
		if min > dis {
			min = dis
		}
	}

	return min
}

func part_2(m_1, m_2 map[position]int) int {
	intersections := make([]position, 0, 1000)

	for pos, _ := range m_1 {
		if _, ok := m_2[pos]; ok {
			if pos.X == 0 && pos.Y == 0 {
				continue
			}

			intersections = append(intersections, pos)
		}
	}

	min := m_1[intersections[0]] + m_2[intersections[0]]

	for i := 1; i < len(intersections); i++ {
		step := m_1[intersections[i]] + m_2[intersections[i]]
		if min > step {
			min = step
		}
	}

	return min
}

type position struct {
	X int
	Y int
}

func distance(pos_1, pos_2 position) int {
	x := pos_1.X - pos_2.X
	if x < 0 {
		x = -x
	}

	y := pos_1.Y - pos_2.Y
	if y < 0 {
		y = -y
	}

	return x + y
}

func walk(route []string) map[position]int {
	m := make(map[position]int)

	root := position{
		X: 0,
		Y: 0,
	}

	cur := root
	m[cur] = 0
	step := 0

	for _, raw := range route {
		var direction string
		var num int
		fmt.Sscanf(raw, "%1s%d", &direction, &num)

		switch direction {
		case "U":
			for i := 1; i <= num; i++ {
				cur = position{
					X: cur.X,
					Y: cur.Y + 1,
				}

				step++

				if _, ok := m[cur]; !ok {
					m[cur] = step
				}
			}
		case "D":
			for i := 1; i <= num; i++ {
				cur = position{
					X: cur.X,
					Y: cur.Y - 1,
				}

				step++

				if _, ok := m[cur]; !ok {
					m[cur] = step
				}
			}
		case "L":
			for i := 1; i <= num; i++ {
				cur = position{
					X: cur.X - 1,
					Y: cur.Y,
				}

				step++

				if _, ok := m[cur]; !ok {
					m[cur] = step
				}
			}
		case "R":
			for i := 1; i <= num; i++ {
				cur = position{
					X: cur.X + 1,
					Y: cur.Y,
				}

				step++

				if _, ok := m[cur]; !ok {
					m[cur] = step
				}
			}
		default:
			continue
		}
	}

	return m
}
