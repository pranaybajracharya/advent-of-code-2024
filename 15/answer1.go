package main

import (
	"fmt"
	"os"
	"strings"
)

type Position struct {
	x int
	y int
}

func nextTick1(grid [][]string, currentPosition Position, move string) ([][]string, Position) {
	if move == "<" {
		dotX := -1
		for i := currentPosition.x; i > 0; i-- {
			if grid[currentPosition.y][i] == "." {
				dotX = i
				break
			} else if grid[currentPosition.y][i] == "#" {
				return grid, currentPosition
			}
		}
		if dotX == -1 {
			return grid, currentPosition
		}
		for i := dotX; i < currentPosition.x; i++ {
			grid[currentPosition.y][i] = grid[currentPosition.y][i+1]
		}
		grid[currentPosition.y][currentPosition.x] = "."
		currentPosition.x--
	} else if move == ">" {
		dotX := -1
		for i := currentPosition.x; i < len(grid[0]); i++ {
			if grid[currentPosition.y][i] == "." {
				dotX = i
				break
			} else if grid[currentPosition.y][i] == "#" {
				return grid, currentPosition
			}
		}
		if dotX == -1 {
			return grid, currentPosition
		}
		for i := dotX; i > currentPosition.x; i-- {
			grid[currentPosition.y][i] = grid[currentPosition.y][i-1]
		}
		grid[currentPosition.y][currentPosition.x] = "."
		currentPosition.x++
	} else if move == "^" {
		dotY := -1
		for i := currentPosition.y; i > 0; i-- {
			if grid[i][currentPosition.x] == "." {
				dotY = i
				break
			} else if grid[i][currentPosition.x] == "#" {
				return grid, currentPosition
			}
		}
		if dotY == -1 {
			return grid, currentPosition
		}
		for i := dotY; i < currentPosition.y; i++ {
			grid[i][currentPosition.x] = grid[i+1][currentPosition.x]
		}
		grid[currentPosition.y][currentPosition.x] = "."
		currentPosition.y--
	} else if move == "v" {
		dotY := -1
		for i := currentPosition.y; i < len(grid); i++ {
			if grid[i][currentPosition.x] == "." {
				dotY = i
				break
			} else if grid[i][currentPosition.x] == "#" {
				return grid, currentPosition
			}
		}
		if dotY == -1 {
			return grid, currentPosition
		}
		for i := dotY; i > currentPosition.y; i-- {
			grid[i][currentPosition.x] = grid[i-1][currentPosition.x]
		}
		grid[currentPosition.y][currentPosition.x] = "."
		currentPosition.y++
	}

	return grid, currentPosition
}

func answer1() {
	// data, err := os.ReadFile("example.txt")
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	input := string(data)
	qn := strings.Split(input, "\n\n")

	rows := strings.Split(qn[0], "\n")
	movesZ := strings.Split(qn[1], "\n")
	moves := ""

	for _, move := range movesZ {
		moves = moves + string(move)
	}

	grid := [][]string{}
	for _, row := range rows {
		chars := strings.Split(row, "")
		grid = append(grid, chars)
	}

	currentPosition := Position{x: 0, y: 0}

	for y, row := range grid {
		for x, char := range row {
			if char == "@" {
				currentPosition = Position{x: x, y: y}
			}
		}
	}

	for _, moveRune := range moves {
		move := string(moveRune)
		grid, currentPosition = nextTick1(grid, currentPosition, move)
	}

	sum := 0

	for y, row := range grid {
		for x, char := range row {
			if char == "O" {
				sum += 100*y + x
			}
		}
	}

	fmt.Println("Answer 1:", sum)
}
