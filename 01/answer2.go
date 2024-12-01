package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func answer2() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	input := string(data)

	rows := strings.Split(input, "\n")

	list1 := []int{}
	list2 := []int{}

	for _, row := range rows {
		nums := strings.Split(row, "   ")

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
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	sum := 0

	for _, value := range list1 {
		count := 0
		for _, value2 := range list2 {
			if value == value2 {
				count++
			}
		}

		sum += value * count
	}

	fmt.Println("Answer 2:", sum)
}
