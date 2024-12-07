package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func answer2() {
	data, err := os.ReadFile("input.txt")
	// data, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	input := string(data)
	rows := strings.Split(input, "\n")
	var equ []Equ

	for _, row := range rows {
		chars := strings.Split(row, ": ")
		value, err := strconv.Atoi(chars[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		equs := strings.Split(chars[1], " ")
		var equations []int
		for i := 0; i < len(equs); i++ {
			num, err := strconv.Atoi(equs[i])
			if err != nil {
				fmt.Println(err)
				return
			}
			equations = append(equations, num)
		}
		equ = append(equ, Equ{value, equations})
	}

	numOfOperators := 3
	sum := findSum(equ, numOfOperators)

	fmt.Println("Answer 2:", sum)
}
