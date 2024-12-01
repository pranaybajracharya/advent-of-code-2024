package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func answer1() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	input := string(data)

	rows := strings.Split(input, "\n")

	list1 := []int{}
	list2 := []int{}
	diff := []int{}

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

	slices.Sort(list1)
	slices.Sort(list2)

	for i := 0; i < len(list1); i++ {
		diff = append(diff, int(math.Abs(float64((list1[i] - list2[i])))))
	}

	sum := 0

	for _, value := range diff {
		sum += value
	}

	fmt.Println("Answer 1:", sum)
}
