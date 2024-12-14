package main

import (
	"fmt"
	"math"
)

func solveLinear(a, b, c, d, e, f float64) (float64, float64, error) {
	// Calculate the determinant of the coefficient matrix
	determinant := a*d - b*c
	if determinant == 0 {
		return 0, 0, fmt.Errorf("the system has no unique solution")
	}

	// Solve using Cramer's Rule
	A := (e*d - b*f) / determinant
	B := (a*f - e*c) / determinant

	return A, B, nil
}

func isInteger(value float64) bool {
	return math.Floor(value) == value
}
