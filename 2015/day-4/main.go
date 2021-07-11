package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("2015/day-4")

	s := "yzbqklnj"

	part_1(s)
	part_2(s)
}

func part_1(s string) {
	fmt.Println(calcAdventCoin(s, isAdventCoin5))
}

func part_2(s string) {
	fmt.Println(calcAdventCoin(s, isAdventCoin6))
}

func calcMD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func isAdventCoin5(s string) bool {
	return strings.HasPrefix(calcMD5(s), "00000")
}

func isAdventCoin6(s string) bool {
	return strings.HasPrefix(calcMD5(s), "000000")
}

func calcAdventCoin(seed string, checkFn func(string) bool) int {
	num := 1

	for {
		coin := seed + strconv.Itoa(num)
		if checkFn(coin) {
			return num
		}

		num++
	}

	return 0
}
