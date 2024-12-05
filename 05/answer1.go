package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkValid(rules [][]string, list []string) bool {
	for i := 0; i < len(list)-1; i++ {
		for _, rule := range rules {
			if rule[0] == list[i+1] && rule[1] == list[i] {
				return false
			}
		}
	}
	return true
}

func answer1() {
	// data, err := os.ReadFile("example.txt")
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	input := string(data)

	rows := strings.Split(input, "\n\n")

	rules := rows[0]
	list := rows[1]

	var formattedRules [][]string
	for _, row := range strings.Split(rules, "\n") {
		formattedRules = append(formattedRules, strings.Split(row, "|"))
	}

	var formattedList [][]string
	for _, row := range strings.Split(list, "\n") {
		formattedList = append(formattedList, strings.Split(row, ","))
	}

	var validList [][]string
	for _, list := range formattedList {
		if checkValid(formattedRules, list) {
			validList = append(validList, list)
		}
	}

	sum := 0
	for _, list := range validList {
		mid := ((len(list) - 1) / 2)
		num, err := strconv.Atoi(list[mid])
		if err != nil {
			fmt.Println(err)
			return
		}
		sum += num
	}

	fmt.Println("Answer 1:", sum)
}
