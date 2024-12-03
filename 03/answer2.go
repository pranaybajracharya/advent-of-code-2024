package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func answer2() {
	// data, err := os.ReadFile("example2.txt")
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	input := string(data)

	pattern := `mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`

	re := regexp.MustCompile(pattern)

	b := re.FindAllString(input, -1)

	do := true
	sum := 0
	for i := 0; i < len(b); i++ {
		if b[i] == "do()" {
			do = true
			continue
		} else if b[i] == "don't()" {
			do = false
			continue
		} else if do {

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
	}

	fmt.Println("Answer 2:", sum)
}
