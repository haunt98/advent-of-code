package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("2020/day-17")

	bytes, err := ioutil.ReadFile("2020/day-17/input.txt")
	if err != nil {
		log.Fatal("read file error")
	}

	lines := strings.Split(string(bytes), "\n")

	m3d := parse3D(lines)
	if isDebug() {
		print3D(m3d)
	}

	part_1(m3d)
}

func part_1(m3d map[int]map[int]map[int]bool) {
	for i := 0; i < 6; i++ {
		m3d = change_1(m3d)
	}

	result := countActive(m3d)
	fmt.Println(result)
}

// init with z = 0
func parse3D(lines []string) map[int]map[int]map[int]bool {
	result := make(map[int]map[int]map[int]bool)
	result[0] = make(map[int]map[int]bool)
	i := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		result[0][i] = make(map[int]bool)
		j := 0

		for _, c := range line {
			if c == '.' {
				result[0][i][j] = false
				j++
				continue
			}

			if c == '#' {
				result[0][i][j] = true
				j++
				continue
			}

			log.Fatalf("what is %c", c)
		}

		i++
	}

	return result
}

func print3D(m map[int]map[int]map[int]bool) {
	for z, m_z := range m {
		fmt.Printf("z = %d\n", z)
		for x, m_x := range m_z {
			fmt.Printf("x = %d\n", x)
			for y, m_y := range m_x {
				fmt.Printf("y = %d ", y)
				if m_y {
					fmt.Printf("[%d]", 1)
				} else {
					fmt.Printf("[%d]", 0)
				}
				fmt.Printf(" ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func countActiveNeighbors(x, y, z int, m3d map[int]map[int]map[int]bool) int {
	count := 0

	for zz := z - 1; zz <= z+1; zz++ {
		for xx := x - 1; xx <= x+1; xx++ {
			for yy := y - 1; yy <= y+1; yy++ {
				if zz == z && xx == x && yy == y {
					continue
				}

				if _, ok := m3d[zz]; !ok {
					continue
				}

				if _, ok := m3d[zz][xx]; !ok {
					continue
				}

				if _, ok := m3d[zz][xx][yy]; !ok {
					continue
				}

				if m3d[zz][xx][yy] {
					count++
				}
			}
		}
	}

	return count
}

func countActive(m3d map[int]map[int]map[int]bool) int {
	count := 0

	for z := range m3d {
		for x := range m3d[z] {
			for y := range m3d[z][x] {
				if m3d[z][x][y] {
					count++
				}
			}
		}
	}

	return count
}

func change_1(m3d map[int]map[int]map[int]bool) map[int]map[int]map[int]bool {
	m3d = expand(m3d)

	// copy
	newM3d := make(map[int]map[int]map[int]bool)
	for z := range m3d {
		newM3d[z] = make(map[int]map[int]bool)
		for x := range m3d[z] {
			newM3d[z][x] = make(map[int]bool)
			for y := range m3d[z][x] {
				newM3d[z][x][y] = m3d[z][x][y]
			}
		}
	}

	// actually change
	for z := range m3d {
		for x := range m3d[z] {
			for y := range m3d[z][x] {
				count := countActiveNeighbors(x, y, z, m3d)
				if isDebug() {
					fmt.Println("z", z, "x", x, "y", y,
						"count", count)
				}

				if m3d[z][x][y] {
					if count == 2 || count == 3 {
						continue
					}
					newM3d[z][x][y] = false
					continue
				}

				if count == 3 {
					newM3d[z][x][y] = true
				}
			}
		}
	}

	return newM3d
}

// expand each x,y,z +- 1
func expand(m3d map[int]map[int]map[int]bool) map[int]map[int]map[int]bool {
	newM3d := make(map[int]map[int]map[int]bool)

	min_z := 0
	min_z_set := false
	max_z := 0
	max_z_set := false

	min_x := 0
	min_x_set := false
	max_x := 0
	max_x_set := false

	min_y := 0
	min_y_set := false
	max_y := 0
	max_y_set := false

	for z := range m3d {
		if !min_z_set {
			min_z = z
			min_z_set = true
		}

		if z < min_z {
			min_z = z
		}

		if !max_z_set {
			max_z = z
			max_z_set = true
		}

		if z > max_z {
			max_z = z
		}

		newM3d[z] = make(map[int]map[int]bool)
		for x := range m3d[z] {
			if !min_x_set {
				min_x = x
				min_x_set = true
			}

			if x < min_x {
				min_x = x
			}

			if !max_x_set {
				max_x = x
				max_x_set = true
			}

			if x > max_x {
				max_x = x
			}

			newM3d[z][x] = make(map[int]bool)
			for y := range m3d[z][x] {
				if !min_y_set {
					min_y = y
					min_y_set = true
				}

				if y < min_y {
					min_y = y
				}

				if !max_y_set {
					max_y = y
					max_y_set = true
				}

				if y > max_y {
					max_y = y
				}

				newM3d[z][x][y] = m3d[z][x][y]
			}
		}
	}

	if isDebug() {
		fmt.Println("z", min_z, max_z)
		fmt.Println("x", min_x, max_x)
		fmt.Println("y", min_y, max_y)
		fmt.Println()
	}

	min_z--
	max_z++
	min_x--
	max_x++
	min_y--
	max_y++

	// fill it up
	for z := min_z; z <= max_z; z++ {
		if _, ok := newM3d[z]; !ok {
			newM3d[z] = make(map[int]map[int]bool)
		}

		for x := min_x; x <= max_x; x++ {
			if _, ok := newM3d[z][x]; !ok {
				newM3d[z][x] = make(map[int]bool)
			}

			for y := min_y; y <= max_y; y++ {
				if _, ok := newM3d[z][x][y]; !ok {
					newM3d[z][x][y] = false
				}
			}
		}
	}

	return newM3d
}

func isDebug() bool {
	return os.Getenv("DEBUG") != ""
}
