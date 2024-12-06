package main

import (
	"fmt"
	"os"
	"strings"
)

func copySlice(original [][]string) [][]string {
	copied := make([][]string, len(original))

	for i := range original {
		copied[i] = make([]string, len(original[i]))
		copy(copied[i], original[i])
	}

	return copied
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
	originalArea := [][]string{}

	for _, row := range rows {
		chars := strings.Split(row, "")
		originalArea = append(originalArea, chars)
	}

	sum := 0

	for i := 0; i < len(originalArea); i++ {
		for j := 0; j < len(originalArea[0]); j++ {
			area := copySlice(originalArea)
			hits := 0
			if area[i][j] == "#" || area[i][j] == "^" || area[i][j] == ">" || area[i][j] == "v" || area[i][j] == "<" {
				continue
			}
			area[i][j] = "O"
			pos := startPos(originalArea)

			for {
				next := nextMove(pos)
				if next.x < 0 || next.y < 0 || next.x >= len(area) || next.y >= len(area[0]) {
					area[pos.x][pos.y] = "X"
					break
				}
				if hits > 1000 { // Random hit limit but it does the job :)
					sum++
					break
				}
				if area[next.x][next.y] == "O" || area[next.x][next.y] == "#" {
					hits++
					pos = rotate(pos)
				} else {
					area[pos.x][pos.y] = "X"
					pos = next
				}
			}
		}
	}

	fmt.Println("Answer 2:", sum)
}
