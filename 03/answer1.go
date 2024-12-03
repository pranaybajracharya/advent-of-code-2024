package main

import (
	"fmt"
	"os"
	"regexp"
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

	r, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)

	b := r.FindAllString(input, -1)

	sum := 0
	for i := 0; i < len(b); i++ {
		substring := b[i][4:]
		final := substring[:len(substring)-1]

		nums := strings.Split(final, ",")

		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			fmt.Println(err)
			return
		}

		sum += num1 * num2
	}

	fmt.Println("Answer 1:", sum)
}
