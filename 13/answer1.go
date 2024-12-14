package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func answer1() {
	// data, err := os.ReadFile("example.txt")
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	input := string(data)

	questions := strings.Split(input, "\n\n")

	sum := 0
	for _, question := range questions {
		qn := strings.Split(question, "\n")

		btnA := strings.Split(qn[0], ": ")
		btnB := strings.Split(qn[1], ": ")
		prize := strings.Split(qn[2], ": ")

		aVal := strings.Split(btnA[1], ", ")
		bVal := strings.Split(btnB[1], ", ")
		prizeVal := strings.Split(prize[1], ", ")

		a, _ := strconv.Atoi(aVal[0][2:])
		b, _ := strconv.Atoi(bVal[0][2:])
		e, _ := strconv.Atoi(prizeVal[0][2:])

		c, _ := strconv.Atoi(aVal[1][2:])
		d, _ := strconv.Atoi(bVal[1][2:])
		f, _ := strconv.Atoi(prizeVal[1][2:])

		A, B, _ := solveLinear(float64(a), float64(b), float64(c), float64(d), float64(e), float64(f))

		if isInteger(A) && isInteger(B) {
			sum += int(A) * 3
			sum += int(B)
		}
	}

	fmt.Println("Answer 1:", sum)
}
