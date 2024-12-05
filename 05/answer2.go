package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sort(rules [][]string, list []string) []string {
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			for _, rule := range rules {
				if rule[1] == list[i] && rule[0] == list[j] {
					list[i], list[j] = list[j], list[i]
				}
			}
		}
	}
	return list
}

func answer2() {
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

	var invalidList [][]string
	for _, list := range formattedList {
		if !checkValid(formattedRules, list) {
			invalidList = append(invalidList, list)
		}
	}

	for i := 0; i < len(invalidList); i++ {
		sort(formattedRules, invalidList[i])
	}

	sum := 0
	for _, list := range invalidList {
		mid := ((len(list) - 1) / 2)
		num, err := strconv.Atoi(list[mid])
		if err != nil {
			fmt.Println(err)
			return
		}
		sum += num
	}

	fmt.Println("Answer 2:", sum)
}
