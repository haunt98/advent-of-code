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

	part_1(lines)
	part_2(lines)
}

func part_1(lines []string) {
	m3d := parse3D(lines)
	for i := 0; i < 6; i++ {
		m3d = cycle3D(m3d)
	}

	result := countActive3D(m3d)
	fmt.Println(result)
}

func part_2(lines []string) {
	m4d := parse4D(lines)
	for i := 0; i < 6; i++ {
		m4d = cycle4D(m4d)
	}

	result := countActive4D(m4d)
	fmt.Println(result)
}

type (
	dimension1st map[int]bool
	dimension2th map[int]dimension1st
	dimension3th map[int]dimension2th
	dimension4th map[int]dimension3th
)

func parse2D(lines []string) dimension2th {
	m2d := make(dimension2th)
	x := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		m2d[x] = make(dimension1st)
		y := 0
		for _, c := range line {
			switch c {
			case '.':
				m2d[x][y] = false
			case '#':
				m2d[x][y] = true
			default:
				log.Fatalf("what is %c", c)
			}
			y++
		}
		x++
	}
	return m2d
}

// init with z = 0
func parse3D(lines []string) dimension3th {
	result := make(dimension3th)
	result[0] = parse2D(lines)
	return result
}

// init with z = 0, t = 0
func parse4D(lines []string) dimension4th {
	result := make(dimension4th)
	result[0] = make(dimension3th)
	result[0][0] = parse2D(lines)
	return result
}

func countActiveNeighbors3D(x, y, z int, m3d dimension3th) int {
	count := 0
	for zz := z - 1; zz <= z+1; zz++ {
		if _, ok := m3d[zz]; !ok {
			continue
		}

		for xx := x - 1; xx <= x+1; xx++ {
			if _, ok := m3d[zz][xx]; !ok {
				continue
			}

			for yy := y - 1; yy <= y+1; yy++ {
				if _, ok := m3d[zz][xx][yy]; !ok {
					continue
				}

				if zz == z && xx == x && yy == y {
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

func countActiveNeighbors4D(x, y, z, t int, m4d dimension4th) int {
	count := 0
	for tt := t - 1; tt <= t+1; tt++ {
		if _, ok := m4d[tt]; !ok {
			continue
		}

		for zz := z - 1; zz <= z+1; zz++ {
			if _, ok := m4d[tt][zz]; !ok {
				continue
			}

			for xx := x - 1; xx <= x+1; xx++ {
				if _, ok := m4d[tt][zz][xx]; !ok {
					continue
				}

				for yy := y - 1; yy <= y+1; yy++ {
					if _, ok := m4d[tt][zz][xx][yy]; !ok {
						continue
					}

					if tt == t && zz == z && xx == x && yy == y {
						continue
					}

					if m4d[tt][zz][xx][yy] {
						count++
					}
				}
			}
		}
	}
	return count
}

func countActive1D(m1d dimension1st) int {
	count := 0
	for y := range m1d {
		if m1d[y] {
			count++
		}
	}
	return count
}

func countActive2D(m2d dimension2th) int {
	count := 0
	for x := range m2d {
		count += countActive1D(m2d[x])
	}
	return count
}

func countActive3D(m3d dimension3th) int {
	count := 0
	for z := range m3d {
		count += countActive2D(m3d[z])
	}
	return count
}

func countActive4D(m4d dimension4th) int {
	count := 0
	for t := range m4d {
		count += countActive3D(m4d[t])
	}
	return count
}

func cycle3D(m3d dimension3th) dimension3th {
	m3d, _, _, _, _, _, _ = expand3D(m3d)

	// copy
	new_m3d := make(dimension3th)
	for z := range m3d {
		new_m3d[z] = make(dimension2th)
		for x := range m3d[z] {
			new_m3d[z][x] = make(dimension1st)
			for y := range m3d[z][x] {
				new_m3d[z][x][y] = m3d[z][x][y]
			}
		}
	}

	// actually change
	for z := range m3d {
		for x := range m3d[z] {
			for y := range m3d[z][x] {
				count := countActiveNeighbors3D(x, y, z, m3d)
				if m3d[z][x][y] {
					if count == 2 || count == 3 {
						continue
					}
					new_m3d[z][x][y] = false
					continue
				} else {
					if count == 3 {
						new_m3d[z][x][y] = true
					}
				}

			}
		}
	}

	return new_m3d
}

func cycle4D(m4d dimension4th) dimension4th {
	m4d, _, _, _, _, _, _, _, _ = expand4D(m4d)

	// copy
	new_m4d := make(dimension4th)
	for t := range m4d {
		new_m4d[t] = make(dimension3th)
		for z := range m4d[t] {
			new_m4d[t][z] = make(dimension2th)
			for x := range m4d[t][z] {
				new_m4d[t][z][x] = make(dimension1st)
				for y := range m4d[t][z][x] {
					new_m4d[t][z][x][y] = m4d[t][z][x][y]
				}
			}
		}
	}

	// actually change
	for t := range m4d {
		for z := range m4d[t] {
			for x := range m4d[t][z] {
				for y := range m4d[t][z][x] {
					count := countActiveNeighbors4D(x, y, z, t, m4d)
					if m4d[t][z][x][y] {
						if count == 2 || count == 3 {
							continue
						}
						new_m4d[t][z][x][y] = false
					} else {
						if count == 3 {
							new_m4d[t][z][x][y] = true
						}
					}

				}
			}
		}
	}

	return new_m4d
}

func expand1D(m1d dimension1st) (dimension1st, int, int) {
	var min_y, max_y int
	var min_y_set, max_y_set bool

	for y := range m1d {
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

		m1d[y] = m1d[y]
	}

	if !min_y_set || !max_y_set {
		log.Fatal("why y tho")
	}

	// y +- 1
	min_y--
	max_y++
	m1d[min_y] = false
	m1d[max_y] = false

	return m1d, min_y, max_y
}

func expand2D(m2d dimension2th) (dimension2th, int, int, int, int) {
	var min_x, max_x int
	var min_x_set, max_x_set bool

	var min_y, max_y int

	for x := range m2d {
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

		m2d[x], min_y, max_y = expand1D(m2d[x])
	}

	// x +- 1
	min_x--
	max_x++
	m2d[min_x] = init1D(min_y, max_y)
	m2d[max_x] = init1D(min_y, max_y)

	return m2d, min_x, max_x, min_y, max_y
}

func expand3D(m3d dimension3th) (dimension3th, int, int, int, int, int, int) {
	var min_z, max_z int
	var min_z_set, max_z_set bool

	var min_x, max_x int
	var min_y, max_y int

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

		m3d[z], min_x, max_x, min_y, max_y = expand2D(m3d[z])
	}

	// z +- 1
	min_z--
	max_z++
	m3d[min_z] = init2D(min_x, max_x, min_y, max_y)
	m3d[max_z] = init2D(min_x, max_x, min_y, max_y)

	return m3d, min_z, max_z, min_x, max_x, min_y, max_y
}

func expand4D(m4d dimension4th) (dimension4th,
	int, int, int, int, int, int, int, int) {
	var min_t, max_t int
	var min_t_set, max_t_set bool

	var min_z, max_z int
	var min_x, max_x int
	var min_y, max_y int

	for t := range m4d {
		if !min_t_set {
			min_t = t
			min_t_set = true
		}

		if t < min_t {
			min_t = t
		}

		if !max_t_set {
			max_t = t
			max_t_set = true
		}

		if t > max_t {
			max_t = t
		}

		m4d[t], min_z, max_z, min_x, max_x, min_y, max_y = expand3D(m4d[t])
	}

	// t +- 1
	min_t--
	max_t++
	m4d[min_t] = init3D(min_z, max_z, min_x, max_x, min_y, max_y)
	m4d[max_t] = init3D(min_z, max_z, min_x, max_x, min_y, max_y)

	return m4d, min_t, max_t, min_z, max_z, min_x, max_x, min_y, max_y
}

func init1D(min_y, max_y int) dimension1st {
	m1d := make(dimension1st)
	for i := min_y; i <= max_y; i++ {
		m1d[i] = false
	}
	return m1d
}

func init2D(min_x, max_x, min_y, max_y int) dimension2th {
	m2d := make(dimension2th)
	for i := min_x; i <= max_x; i++ {
		m2d[i] = init1D(min_y, max_y)
	}
	return m2d
}

func init3D(min_z, max_z, min_x, max_x, min_y, max_y int) dimension3th {
	m3d := make(dimension3th)
	for i := min_z; i <= max_z; i++ {
		m3d[i] = init2D(min_x, max_x, min_y, max_y)
	}
	return m3d
}

func isDebug() bool {
	return os.Getenv("DEBUG") != ""
}
