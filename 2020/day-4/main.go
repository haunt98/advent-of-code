package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	byr = iota + 1
	iyr
	eyr
	hgt
	hcl
	ecl
	pid
)

const (
	defaultLen = 1000
)

type ruleFn func(data string) bool

var (
	ruleFns = map[int]ruleFn{
		byr: ruleBYR,
		iyr: ruleIYR,
		eyr: ruleEYR,
		hgt: ruleHGT,
		hcl: ruleHCL,
		ecl: ruleECL,
		pid: rulePID,
	}
)

func main() {
	fmt.Println("2020/day-4")

	bytes, err := ioutil.ReadFile("2020/day-4/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := string(bytes)
	passportsRaw := strings.Split(data, "\n\n")

	passports := make([]map[int]string, 0, defaultLen)

	for _, passportRaw := range passportsRaw {
		fields := strings.Fields(passportRaw)

		m := make(map[int]string)

		for _, field := range fields {
			var fName, fData string
			fmt.Sscanf(field, "%3s:%s", &fName, &fData)

			fID := convertField(fName)
			if fID > 0 {
				m[fID] = fData
			}
		}

		passports = append(passports, m)
	}

	fmt.Println(part_1(passports))
	fmt.Println(part_2(passports))
}

func part_1(passports []map[int]string) int {
	result := 0

	for _, passport := range passports {
		valid := true

		for i := byr; i <= pid; i++ {
			if _, ok := passport[i]; !ok {
				valid = false
				break
			}
		}

		if valid {
			result++
		}
	}

	return result
}

func part_2(passports []map[int]string) int {
	result := 0

	for _, passport := range passports {
		valid := true

		for fID, fn := range ruleFns {
			if _, ok := passport[fID]; !ok {
				valid = false
				break
			}

			if !fn(passport[fID]) {
				valid = false
				break
			}
		}

		if valid {
			result++
		}
	}

	return result
}

func convertField(f string) int {
	switch f {
	case "byr":
		return byr
	case "iyr":
		return iyr
	case "eyr":
		return eyr
	case "hgt":
		return hgt
	case "hcl":
		return hcl
	case "ecl":
		return ecl
	case "pid":
		return pid
	default:
		return 0
	}
}

func ruleBYR(data string) bool {
	year, err := strconv.Atoi(data)
	if err != nil {
		return false
	}

	return year >= 1920 && year <= 2002
}

func ruleIYR(data string) bool {
	year, err := strconv.Atoi(data)
	if err != nil {
		return false
	}

	return year >= 2010 && year <= 2020
}

func ruleEYR(data string) bool {
	year, err := strconv.Atoi(data)
	if err != nil {
		return false
	}

	return year >= 2020 && year <= 2030
}

func ruleHGT(data string) bool {
	if strings.HasSuffix(data, "cm") {
		data = strings.TrimSuffix(data, "cm")

		num, err := strconv.Atoi(data)
		if err != nil {
			return false
		}

		return num >= 150 && num <= 193
	}

	if strings.HasSuffix(data, "in") {
		data = strings.TrimSuffix(data, "in")

		num, err := strconv.Atoi(data)
		if err != nil {
			return false
		}

		return num >= 59 && num <= 76
	}

	return false
}

func ruleHCL(data string) bool {
	if !strings.HasPrefix(data, "#") {
		return false
	}

	if len(data[1:]) != 6 {
		return false
	}

	for i := 1; i < len(data); i++ {
		if data[i] >= '0' && data[i] <= '9' {
			continue
		}

		if data[i] >= 'a' && data[i] <= 'f' {
			continue
		}

		return false
	}

	return true
}

func ruleECL(data string) bool {
	switch data {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	default:
		return false
	}
}

func rulePID(data string) bool {
	if len(data) != 9 {
		return false
	}

	for i := 0; i < len(data); i++ {
		if data[i] >= '0' && data[i] <= '9' {
			continue
		}

		return false
	}

	return true
}
