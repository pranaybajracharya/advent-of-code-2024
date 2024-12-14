package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

type Velocity struct {
	x int
	y int
}

var tileX = 101
var tileY = 103

var iteration = 100

func formatPosition(position Position, velocity Velocity) Position {
	newX := position.x + velocity.x
	newY := position.y + velocity.y

	newX = (newX + tileX) % tileX
	newY = (newY + tileY) % tileY

	return Position{x: newX, y: newY}
}

func answer1() {
	// data, err := os.ReadFile("example.txt")
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	input := string(data)

	robots := strings.Split(input, "\n")

	var robotPositions []Position
	var robotVelocity []Velocity
	for _, robot := range robots {
		qn := strings.Split(robot, " ")

		pp := qn[0][2:]
		vv := qn[1][2:]
		p := strings.Split(pp, ",")
		v := strings.Split(vv, ",")

		px, _ := strconv.Atoi(p[0])
		py, _ := strconv.Atoi(p[1])
		vx, _ := strconv.Atoi(v[0])
		vy, _ := strconv.Atoi(v[1])

		robotPositions = append(robotPositions, Position{x: px, y: py})
		robotVelocity = append(robotVelocity, Velocity{x: vx, y: vy})
	}

	grid := [][]string{}
	for y := 0; y < tileY; y++ {
		row := []string{}
		for x := 0; x < tileX; x++ {
			row = append(row, ".")
		}
		grid = append(grid, row)
	}

	for it := 0; it < iteration; it++ {
		for i, velocity := range robotVelocity {
			robotPositions[i] = formatPosition(robotPositions[i], velocity)
		}
	}

	for _, rp := range robotPositions {
		if grid[rp.y][rp.x] == "." {
			grid[rp.y][rp.x] = "1"
		} else {
			v, _ := strconv.Atoi(grid[rp.y][rp.x])
			v++
			grid[rp.y][rp.x] = strconv.FormatInt(int64(v), 10)
		}
	}

	midX := tileX / 2
	midY := tileY / 2

	quad1 := 0
	for x := 0; x < midX; x++ {
		for y := 0; y < midY; y++ {
			if grid[y][x] == "." {
				continue
			} else {
				c, _ := strconv.Atoi(grid[y][x])
				quad1 += c
			}
		}
	}

	quad2 := 0
	for x := midX + 1; x < tileX; x++ {
		for y := 0; y < midY; y++ {
			if grid[y][x] == "." {
				continue
			} else {
				c, _ := strconv.Atoi(grid[y][x])
				quad2 += c
			}
		}
	}

	quad3 := 0
	for x := 0; x < midX; x++ {
		for y := midY + 1; y < tileY; y++ {
			if grid[y][x] == "." {
				continue
			} else {
				c, _ := strconv.Atoi(grid[y][x])
				quad3 += c
			}
		}
	}

	quad4 := 0
	for x := midX + 1; x < tileX; x++ {
		for y := midY + 1; y < tileY; y++ {
			if grid[y][x] == "." {
				continue
			} else {
				c, _ := strconv.Atoi(grid[y][x])
				quad4 += c
			}
		}
	}

	fmt.Println("Answer 1:", quad1*quad2*quad3*quad4)
}
