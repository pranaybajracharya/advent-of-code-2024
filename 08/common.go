package main

import (
	"strings"
)

func leftPad(s string, count int) string {
	z := len(s)
	return strings.Repeat("0", count-z) + s
}

func copySlice(original [][]string) [][]string {
	copied := make([][]string, len(original))

	for i := range original {
		copied[i] = make([]string, len(original[i]))
		copy(copied[i], original[i])
	}

	return copied
}

func drawAntiNode(area [][]string, node []int) {
	if node[0] < 0 || node[0] >= len(area) || node[1] < 0 || node[1] >= len(area[0]) {
		return
	}

	area[node[0]][node[1]] = "#"
}
