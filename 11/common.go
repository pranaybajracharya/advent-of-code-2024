package main

import "strconv"

func trans(num int) []int {
	if num == 0 {
		return []int{1}
	}

	s := strconv.Itoa(num)
	length := len(s)
	if length%2 == 0 {
		mid := length / 2
		firstNum, _ := strconv.Atoi(s[:mid])
		secondNum, _ := strconv.Atoi(s[mid:])

		return []int{firstNum, secondNum}
	}

	return []int{num * 2024}
}

func count(counts map[int]int) map[int]int {
	m := make(map[int]int)

	for k, v := range counts {
		for _, num := range trans(k) {
			m[num] += v
		}
	}

	return m
}
