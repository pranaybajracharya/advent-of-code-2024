package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var blinks2 = 75

func answer2() {
	data, err := os.ReadFile("input.txt")
	// data, err := os.ReadFile("example.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	input := string(data)
	str := strings.Split(input, " ")

	m := make(map[int]int)

	for i := 0; i < len(str); i++ {
		num, err := strconv.Atoi(str[i])
		if err != nil {
			fmt.Println(err)
			return
		}

		m[num]++
	}

	for i := 0; i < blinks2; i++ {
		m = count(m)
	}

	sum := 0

	for _, v := range m {
		sum += v
	}

	fmt.Println("Answer 2:", sum)
}
