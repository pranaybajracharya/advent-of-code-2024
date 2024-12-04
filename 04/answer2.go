package main

import (
	"fmt"
	"os"
	"strings"
)

const MAS = "MAS"

func checkX(list [][]string) int {
	sum := 0
	for i := 0; i < len(list)-len(MAS)+1; i++ {
		for j := 0; j < len(list[i])-len(MAS)+1; j++ {
			word1 := list[i][j] + list[i+1][j+1] + list[i+2][j+2]
			wordReverse1 := list[i+2][j+2] + list[i+1][j+1] + list[i][j]

			word2 := list[i][j+2] + list[i+1][j+1] + list[i+2][j]
			wordReverse2 := list[i+2][j] + list[i+1][j+1] + list[i][j+2]

			if (word1 == MAS || wordReverse1 == MAS) && (word2 == MAS || wordReverse2 == MAS) {
				sum++
			}
		}
	}
	return sum
}

func answer2() {
	data, err := os.ReadFile("input.txt")
	// data, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	input := string(data)

	rows := strings.Split(input, "\n")

	list := [][]string{}

	for _, row := range rows {
		chars := strings.Split(row, "")

		list = append(list, chars)
	}

	sum := 0
	sum += checkX(list)
	fmt.Println("Answer 2:", sum)
}
