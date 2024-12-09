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
	diskmap := strings.Split(input, "")

	var dm []int
	idx := 0
	for i, disk := range diskmap {
		count, err := strconv.Atoi(disk)
		if err != nil {
			fmt.Println(err)
			return
		}
		if i%2 == 0 {
			for i := 0; i < count; i++ {
				dm = append(dm, idx)
			}
			idx++
		} else {
			for i := 0; i < count; i++ {
				dm = append(dm, -1)
			}
		}
	}

	ptr1 := 0
	ptr2 := len(dm) - 1
	for {
		if ptr1 >= ptr2 {
			break
		}

		if dm[ptr1] != -1 {
			ptr1++
			continue
		}

		if dm[ptr2] == -1 {
			ptr2--
			continue
		}

		dm[ptr1] = dm[ptr2]
		dm[ptr2] = -1
		ptr1++
		ptr2--
	}

	sum := 0

	for i, char := range dm {
		if char == -1 {
			break
		}
		sum += char * i
	}

	fmt.Println("Answer 1:", sum)
}
