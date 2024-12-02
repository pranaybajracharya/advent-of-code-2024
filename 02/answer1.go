package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func answer1() {
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
		}
	}

	fmt.Println("Answer 1:", safeCount)
}
