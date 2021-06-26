package main

import "fmt"

func main() {
	fmt.Println("2015/day-2")
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
	return 2*g.l*g.w + 2*g.w*g.h + 2*g.h*g.l + min3(g.l, g.w, g.h)
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
