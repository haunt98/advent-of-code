package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("2015/day-2")

	gifts := prepare()
	part_1(gifts)
	part_2(gifts)
}

func prepare() []gift {
	bytes, err := os.ReadFile("2015/day-2/input.txt")
	if err != nil {
		return nil
	}

	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	gifts := make([]gift, 0, len(lines))

	for _, line := range lines {
		var l, w, h int
		fmt.Sscanf(line, "%dx%dx%d", &l, &w, &h)
		gifts = append(gifts, newGift(l, w, h))
	}

	return gifts
}

func part_1(gifts []gift) {
	total := 0

	for _, g := range gifts {
		total += getWrappingPaper(g)
	}

	fmt.Println(total)
}

func part_2(gifts []gift) {
	total := 0

	for _, g := range gifts {
		total += getRibbon(g)
	}

	fmt.Println(total)
}

type gift struct {
	l int
	w int
	h int
}

func newGift(l, w, h int) gift {
	return gift{
		l: l,
		w: w,
		h: h,
	}
}

func getWrappingPaper(g gift) int {
	// s is area
	sLW := g.l * g.w
	sWH := g.w * g.h
	sHL := g.h * g.l
	return 2*sLW + 2*sWH + 2*sHL + min3(sLW, sWH, sHL)
}

func getRibbon(g gift) int {
	// p is perimeter
	pLW := 2 * (g.l + g.w)
	pWH := 2 * (g.w + g.h)
	pHL := 2 * (g.h + g.l)
	return min3(pLW, pWH, pHL) + g.l*g.w*g.h
}

func min2(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func min3(a, b, c int) int {
	return min2(min2(a, b), c)
}
