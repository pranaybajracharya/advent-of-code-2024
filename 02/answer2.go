package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func remove(slice []int, s int) []int {
	var n []int
	for i := 0; i < len(slice); i++ {
		if i == s {
			continue
		}
		n = append(n, slice[i])
	}

	return n
}

func answer2() {
	// data, err := os.ReadFile("example.txt")
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	input := string(data)

	rows := strings.Split(input, "\n")

	safeCount := 0
	for _, row := range rows {
		numss := strings.Split(row, " ")
		var nums []int
		for i := 0; i < len(numss); i++ {
			num, err := strconv.Atoi(numss[i])
			if err != nil {
				fmt.Println(err)
				return
			}
			nums = append(nums, num)
		}

		if safe(nums) {
			safeCount++
		} else {
			for i := 0; i < len(nums); i++ {
				newNums := remove(nums, i)
				if safe(newNums) {
					safeCount++
					break
				}
			}
		}
	}

	fmt.Println("Answer 2:", safeCount)
}
