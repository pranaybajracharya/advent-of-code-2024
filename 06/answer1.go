package main

import (
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x    int
	y    int
	char string
}

func startPos(area [][]string) Pos {
	for i, row := range area {
		for j, char := range row {
			if char == "^" || char == ">" || char == "v" || char == "<" {
				return Pos{x: i, y: j, char: char}
			}
		}
	}
	return Pos{x: -1, y: -1}
}

func nextMove(pos Pos) Pos {
	if pos.char == "^" {
		return Pos{x: pos.x - 1, y: pos.y, char: "^"}
	} else if pos.char == ">" {
		return Pos{x: pos.x, y: pos.y + 1, char: ">"}
	} else if pos.char == "v" {
		return Pos{x: pos.x + 1, y: pos.y, char: "v"}
	} else if pos.char == "<" {
		return Pos{x: pos.x, y: pos.y - 1, char: "<"}
	}
	return Pos{x: -1, y: -1}
}

func rotate(pos Pos) Pos {
	if pos.char == "^" {
		return Pos{x: pos.x, y: pos.y, char: ">"}
	} else if pos.char == ">" {
		return Pos{x: pos.x, y: pos.y, char: "v"}
	} else if pos.char == "v" {
		return Pos{x: pos.x, y: pos.y, char: "<"}
	} else if pos.char == "<" {
		return Pos{x: pos.x, y: pos.y, char: "^"}
	}
	return Pos{x: -1, y: -1}
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
	area := [][]string{}

	for _, row := range rows {
		chars := strings.Split(row, "")
		area = append(area, chars)
	}

	pos := startPos(area)

	for {
		next := nextMove(pos)
		if next.x < 0 || next.y < 0 || next.x >= len(area) || next.y >= len(area[0]) {
			area[pos.x][pos.y] = "X"
			break
		}
		if area[next.x][next.y] == "#" {
			pos = rotate(pos)
		} else {
			area[pos.x][pos.y] = "X"
			pos = next
		}
	}

	sum := 0

	for _, row := range area {
		for _, char := range row {
			if char == "X" {
				sum++
			}
		}
	}

	fmt.Println("Answer 1:", sum)
}
