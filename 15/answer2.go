package main

import (
	"fmt"
	"os"
	"strings"
)

type Box struct {
	left  Position
	right Position
}

func copySlice(original []Box) []Box {
	copied := make([]Box, len(original))
	copy(copied, original)
	return copied
}

func arePositionsEqual(a, b Position) bool {
	return a.x == b.x && a.y == b.y
}

func areBoxesEqual(a, b Box) bool {
	return arePositionsEqual(a.left, b.left) && arePositionsEqual(a.right, b.right)
}

func containsBox(boxes []Box, box Box) bool {
	for _, b := range boxes {
		if areBoxesEqual(b, box) {
			return true
		}
	}
	return false
}

func moveBoxesUp(grid [][]string, currentPosition Position) ([][]string, Position) {
	listOfBoxes := []Box{}

	if grid[currentPosition.y-1][currentPosition.x] == "[" {
		listOfBoxes = append(listOfBoxes, Box{left: Position{x: currentPosition.x, y: currentPosition.y - 1}, right: Position{x: currentPosition.x + 1, y: currentPosition.y - 1}})
	} else {
		listOfBoxes = append(listOfBoxes, Box{left: Position{x: currentPosition.x - 1, y: currentPosition.y - 1}, right: Position{x: currentPosition.x, y: currentPosition.y - 1}})
	}

	for {
		prevBoxes := copySlice(listOfBoxes)
		for _, box := range listOfBoxes {
			if grid[box.left.y-1][box.left.x] == "[" {
				newBox := Box{left: Position{x: box.left.x, y: box.left.y - 1}, right: Position{x: box.right.x, y: box.right.y - 1}}
				if !containsBox(listOfBoxes, newBox) {
					listOfBoxes = append(listOfBoxes, newBox)
				}
			}
			if grid[box.right.y-1][box.right.x] == "]" {
				newBox := Box{left: Position{x: box.left.x, y: box.left.y - 1}, right: Position{x: box.right.x, y: box.right.y - 1}}
				if !containsBox(listOfBoxes, newBox) {
					listOfBoxes = append(listOfBoxes, newBox)
				}
			}
			if grid[box.left.y-1][box.left.x] == "]" {
				newBox := Box{left: Position{x: box.left.x - 1, y: box.left.y - 1}, right: Position{x: box.right.x - 1, y: box.right.y - 1}}
				if !containsBox(listOfBoxes, newBox) {
					listOfBoxes = append(listOfBoxes, newBox)
				}
			}
			if grid[box.right.y-1][box.right.x] == "[" {
				newBox := Box{left: Position{x: box.left.x + 1, y: box.left.y - 1}, right: Position{x: box.right.x + 1, y: box.right.y - 1}}
				if !containsBox(listOfBoxes, newBox) {
					listOfBoxes = append(listOfBoxes, newBox)
				}
			}
		}

		if len(prevBoxes) == len(listOfBoxes) {
			break
		}
	}
	allBoxesMoveable := true
	for _, box := range listOfBoxes {
		if grid[box.left.y-1][box.left.x] == "#" || grid[box.right.y-1][box.right.x] == "#" {
			allBoxesMoveable = false
			break
		}
	}

	if !allBoxesMoveable {
		return grid, currentPosition
	}

	reverseListOfBoxes := []Box{}
	for i := len(listOfBoxes) - 1; i >= 0; i-- {
		reverseListOfBoxes = append(reverseListOfBoxes, listOfBoxes[i])
	}

	for _, box := range reverseListOfBoxes {
		grid[box.left.y-1][box.left.x] = grid[box.left.y][box.left.x]
		grid[box.left.y][box.left.x] = "."
		grid[box.right.y-1][box.right.x] = grid[box.right.y][box.right.x]
		grid[box.right.y][box.right.x] = "."
	}
	grid[currentPosition.y-1][currentPosition.x] = "@"
	grid[currentPosition.y][currentPosition.x] = "."
	currentPosition.y--
	return grid, currentPosition
}

func moveBoxesDown(grid [][]string, currentPosition Position) ([][]string, Position) {
	listOfBoxes := []Box{}

	if grid[currentPosition.y+1][currentPosition.x] == "[" {
		listOfBoxes = append(listOfBoxes, Box{left: Position{x: currentPosition.x, y: currentPosition.y + 1}, right: Position{x: currentPosition.x + 1, y: currentPosition.y + 1}})
	} else {
		listOfBoxes = append(listOfBoxes, Box{left: Position{x: currentPosition.x - 1, y: currentPosition.y + 1}, right: Position{x: currentPosition.x, y: currentPosition.y + 1}})
	}

	for {
		prevBoxes := copySlice(listOfBoxes)
		for _, box := range listOfBoxes {
			if grid[box.left.y+1][box.left.x] == "[" {
				newBox := Box{left: Position{x: box.left.x, y: box.left.y + 1}, right: Position{x: box.right.x, y: box.right.y + 1}}
				if !containsBox(listOfBoxes, newBox) {
					listOfBoxes = append(listOfBoxes, newBox)
				}
			}
			if grid[box.right.y+1][box.right.x] == "]" {
				newBox := Box{left: Position{x: box.left.x, y: box.left.y + 1}, right: Position{x: box.right.x, y: box.right.y + 1}}
				if !containsBox(listOfBoxes, newBox) {
					listOfBoxes = append(listOfBoxes, newBox)
				}
			}
			if grid[box.left.y+1][box.left.x] == "]" {
				newBox := Box{left: Position{x: box.left.x - 1, y: box.left.y + 1}, right: Position{x: box.right.x - 1, y: box.right.y + 1}}
				if !containsBox(listOfBoxes, newBox) {
					listOfBoxes = append(listOfBoxes, newBox)
				}
			}
			if grid[box.right.y+1][box.right.x] == "[" {
				newBox := Box{left: Position{x: box.left.x + 1, y: box.left.y + 1}, right: Position{x: box.right.x + 1, y: box.right.y + 1}}
				if !containsBox(listOfBoxes, newBox) {
					listOfBoxes = append(listOfBoxes, newBox)
				}
			}
		}

		if len(prevBoxes) == len(listOfBoxes) {
			break
		}
	}

	allBoxesMoveable := true
	for _, box := range listOfBoxes {
		if grid[box.left.y+1][box.left.x] == "#" || grid[box.right.y+1][box.right.x] == "#" {
			allBoxesMoveable = false
			break
		}
	}

	if !allBoxesMoveable {
		return grid, currentPosition
	}

	reverseListOfBoxes := []Box{}
	for i := len(listOfBoxes) - 1; i >= 0; i-- {
		reverseListOfBoxes = append(reverseListOfBoxes, listOfBoxes[i])
	}

	for _, box := range reverseListOfBoxes {
		grid[box.left.y+1][box.left.x] = grid[box.left.y][box.left.x]
		grid[box.left.y][box.left.x] = "."
		grid[box.right.y+1][box.right.x] = grid[box.right.y][box.right.x]
		grid[box.right.y][box.right.x] = "."
	}
	grid[currentPosition.y+1][currentPosition.x] = "@"
	grid[currentPosition.y][currentPosition.x] = "."
	currentPosition.y++
	return grid, currentPosition
}

func nextTick2(grid [][]string, currentPosition Position, move string) ([][]string, Position) {
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
		if grid[currentPosition.y-1][currentPosition.x] == "#" {
			return grid, currentPosition
		}

		if grid[currentPosition.y-1][currentPosition.x] == "." {
			grid[currentPosition.y-1][currentPosition.x] = "@"
			grid[currentPosition.y][currentPosition.x] = "."
			currentPosition.y--

			return grid, currentPosition
		}

		return moveBoxesUp(grid, currentPosition)

	} else if move == "v" {
		if grid[currentPosition.y+1][currentPosition.x] == "#" {
			return grid, currentPosition
		}

		if grid[currentPosition.y+1][currentPosition.x] == "." {
			grid[currentPosition.y+1][currentPosition.x] = "@"
			grid[currentPosition.y][currentPosition.x] = "."
			currentPosition.y++

			return grid, currentPosition
		}

		return moveBoxesDown(grid, currentPosition)
	}

	return grid, currentPosition
}

func answer2() {
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

	smallGrid := [][]string{}
	for _, row := range rows {
		chars := strings.Split(row, "")
		smallGrid = append(smallGrid, chars)
	}

	grid := [][]string{}
	for _, row := range smallGrid {
		newRow := []string{}
		for _, char := range row {
			if char == "." {
				newRow = append(newRow, ".", ".")
			} else if char == "#" {
				newRow = append(newRow, "#", "#")
			} else if char == "O" {
				newRow = append(newRow, "[", "]")
			} else if char == "@" {
				newRow = append(newRow, "@", ".")
			}
		}
		grid = append(grid, newRow)
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
		grid, currentPosition = nextTick2(grid, currentPosition, move)
	}

	sum := 0

	for y, row := range grid {
		for x, char := range row {
			if char == "[" {
				sum += 100*y + x
			}
		}
	}

	fmt.Println("Answer 2:", sum)
}
