package main

import (
	"fmt"
	"math"
)

func main() {
	// If we don't use "var" to create variables, we can't declare a type for the variable.
	// var userHeight float64 = 1.8
	// var userKg float64 = 100

	// userHeight, userKg := 1.8, 100
	userHeight := 180
	userKg := 100

	const IMTPower float64 = 2 // We can't reassign the value because it is a constant.

	IMT := float64(userKg) / math.Pow(float64(userHeight) / 100, IMTPower)

	fmt.Print(IMT)
}