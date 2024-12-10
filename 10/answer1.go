package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func answer1() {
	data, err := os.ReadFile("input.txt")
	// data, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	input := string(data)
	rows := strings.Split(input, "\n")
	area := [][]int{}
	for _, row := range rows {
		chars := strings.Split(row, "")
		nums := []int{}
		for _, char := range chars {
			num, err := strconv.Atoi(char)
			if err != nil {
				fmt.Println(err)
				return
			}
			nums = append(nums, num)
		}
		area = append(area, nums)
	}

	sum := 0

	for i := 0; i < len(area); i++ {
		for j := 0; j < len(area[0]); j++ {
			pos := Pos{x: i, y: j, value: area[i][j]}

			if pos.value == 0 {
				destination := []Pos{}
				found := findPaths(area, pos, destination)
				sum += len(uniquePositions(found))
			}
		}
	}

	fmt.Println("Answer 1:", sum)
}
