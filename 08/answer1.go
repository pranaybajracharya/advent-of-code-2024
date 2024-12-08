package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func plantAntiNodes(area [][]string, node [][]int, pattern string) {
	if string(pattern[0]) == string(pattern[1]) {
		return
	}
	idx1, err := strconv.Atoi(string(pattern[0]))
	if err != nil {
		fmt.Println(err)
		return
	}
	idx2, err := strconv.Atoi(string(pattern[1]))
	if err != nil {
		fmt.Println(err)
		return
	}
	vertex1 := node[idx1]
	vertex2 := node[idx2]

	diffX := vertex2[0] - vertex1[0]
	diffY := vertex2[1] - vertex1[1]

	antiNode1 := []int{vertex1[0] - diffX, vertex1[1] - diffY}
	antiNode2 := []int{vertex2[0] + diffX, vertex2[1] + diffY}

	drawAntiNode(area, antiNode1)
	drawAntiNode(area, antiNode2)
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

	nodes := make(map[string][][]int)
	for i := 0; i < len(area); i++ {
		for j := 0; j < len(area[0]); j++ {
			char := area[i][j]
			if char != "." {
				nodes[char] = append(nodes[char], []int{i, j})
			}
		}
	}

	newArea := copySlice(area)
	for _, n := range nodes {
		totalCheck := int(math.Pow(float64(len(n)), float64(2)) - 1)
		patterns := []string{}
		for i := 0; i <= totalCheck; i++ {
			pattern := strconv.FormatInt(int64(i), len(n))
			patterns = append(patterns, leftPad(pattern, 2))
		}

		for _, p := range patterns {
			plantAntiNodes(newArea, n, p)
		}
	}

	sum := 0

	for i := 0; i < len(newArea); i++ {
		for j := 0; j < len(newArea[0]); j++ {
			if newArea[i][j] == "#" {
				sum++
			}
		}
	}

	fmt.Println("Answer 1:", sum)
}
