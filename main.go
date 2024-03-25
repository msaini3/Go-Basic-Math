package main

import (
	"fmt"
	"mathops"
)

func main() {
	// Define two operands
	x := 10.0
	y := 5.0

	// Create channels to receive results
	addResult := make(chan float64)
	subResult := make(chan float64)
	mulResult := make(chan float64)
	divResult := make(chan float64)

	// Perform addition
	go mathops.Calculate(mathops.Add{}, x, y, addResult)

	// Perform subtraction
	go mathops.Calculate(mathops.Subtract{}, x, y, subResult)

	// Perform multiplication
	go mathops.Calculate(mathops.Multiply{}, x, y, mulResult)

	// Perform division
	go mathops.Calculate(mathops.Divide{}, x, y, divResult)

	// Receive and print results
	fmt.Printf("Addition Result: %.2f\n", <-addResult)
	fmt.Printf("Subtraction Result: %.2f\n", <-subResult)
	fmt.Printf("Multiplication Result: %.2f\n", <-mulResult)
	fmt.Printf("Division Result: %.2f\n", <-divResult)
}
