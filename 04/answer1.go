package main

import (
	"fmt"
	"os"
	"strings"
)

const XMAS = "XMAS"

func checkHorizontal(list [][]string) int {
	sum := 0
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list[i])-len(XMAS)+1; j++ {
			word := list[i][j] + list[i][j+1] + list[i][j+2] + list[i][j+3]
			wordReverse := list[i][j+3] + list[i][j+2] + list[i][j+1] + list[i][j]
			if word == XMAS || wordReverse == XMAS {
				sum++
			}
		}
	}
	return sum
}

func checkVertical(list [][]string) int {
	sum := 0
	for i := 0; i < len(list)-len(XMAS)+1; i++ {
		for j := 0; j < len(list[i]); j++ {
			word := list[i][j] + list[i+1][j] + list[i+2][j] + list[i+3][j]
			wordReverse := list[i+3][j] + list[i+2][j] + list[i+1][j] + list[i][j]
			if word == XMAS || wordReverse == XMAS {
				sum++
			}
		}
	}
	return sum
}

func checkLRDiagonal(list [][]string) int {
	sum := 0
	for i := 0; i < len(list)-len(XMAS)+1; i++ {
		for j := 0; j < len(list[i])-len(XMAS)+1; j++ {
			word := list[i][j] + list[i+1][j+1] + list[i+2][j+2] + list[i+3][j+3]
			wordReverse := list[i+3][j+3] + list[i+2][j+2] + list[i+1][j+1] + list[i][j]
			if word == XMAS || wordReverse == XMAS {
				sum++
			}
		}
	}
	return sum
}

func checkRLDiagonal(list [][]string) int {
	sum := 0
	for i := 0; i < len(list)-len(XMAS)+1; i++ {
		for j := 3; j < len(list[i]); j++ {
			word := list[i][j] + list[i+1][j-1] + list[i+2][j-2] + list[i+3][j-3]
			wordReverse := list[i+3][j-3] + list[i+2][j-2] + list[i+1][j-1] + list[i][j]
			if word == XMAS || wordReverse == XMAS {
				sum++
			}
		}
	}
	return sum
}

func answer1() {
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
	sum += checkHorizontal(list)
	sum += checkVertical(list)
	sum += checkLRDiagonal(list)
	sum += checkRLDiagonal(list)
	fmt.Println("Answer 1:", sum)
}
