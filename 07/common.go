package main

import (
	"math"
	"strconv"
	"strings"
)

type Equ struct {
	value     int
	equations []int
}

func leftPad(s string, count int) string {
	z := len(s)
	return strings.Repeat("0", count-z) + s
}

func compute(pattern string, equ Equ) bool {
	value := equ.equations[0]
	for i := 0; i < len(equ.equations)-1; i++ {
		if pattern[i] == '0' {
			value += equ.equations[i+1]
		} else if pattern[i] == '1' {
			value *= equ.equations[i+1]
		} else if pattern[i] == '2' {
			current := strconv.FormatInt(int64(value), 10)
			next := strconv.FormatInt(int64(equ.equations[i+1]), 10)
			// Concat accumulator with next digit
			newValue, err := strconv.Atoi(strings.Join([]string{current, next}, ""))
			if err != nil {
				return false
			}
			value = newValue
		}
	}

	if value == equ.value {
		return true
	} else {
		return false
	}
}

func findSum(equ []Equ, numOfOperators int) int {
	sum := 0
	for _, e := range equ {
		totalCheck := int(math.Pow(float64(numOfOperators), float64(len(e.equations)-1)) - 1)
		patterns := []string{}
		for i := 0; i <= totalCheck; i++ {
			pattern := strconv.FormatInt(int64(i), numOfOperators)
			patterns = append(patterns, leftPad(pattern, len(e.equations)-1))
		}

		for _, p := range patterns {
			if compute(p, e) {
				sum += e.value
				break
			}
		}
	}

	return sum
}
