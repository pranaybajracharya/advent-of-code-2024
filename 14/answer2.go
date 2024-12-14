package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func answer2() {
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

	for it := 0; it < 100000; it++ {
		for i, velocity := range robotVelocity {
			robotPositions[i] = formatPosition(robotPositions[i], velocity)
		}

		grid := [][]string{}
		for y := 0; y < tileY; y++ {
			row := []string{}
			for x := 0; x < tileX; x++ {
				row = append(row, " ")
			}
			grid = append(grid, row)
		}
		for _, rp := range robotPositions {
			if grid[rp.y][rp.x] == " " {
				grid[rp.y][rp.x] = "X"
			}
		}
		fmt.Println("*************************************** ", it, " ***************************************************")
		for _, y := range grid {
			fmt.Println(y)
		}
		time.Sleep(time.Second / 10)
	}
}
