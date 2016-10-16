package main

import (
	"fmt"
	"math"
)

// MaxIterations Setting up maximum number of iterations withing application
const MaxIterations = 99999

// Sqrt Simple square root with 5 number of iterations
func Sqrt(number float64) float64 {
	return SqrtIteration(number, 5)
}

// Step is a step function
func Step(value, number float64) float64 {
	return value - (((value * value) - number) / (2 * (value)))
}

// SqrtIteration This function is used for counting square root of number
// with custome iteration number
func SqrtIteration(number float64, iterationNumber int) float64 {
	value := float64(1.0)

	//Set Max iterationNumber, in case if someone tries to make it too large
	if iterationNumber > MaxIterations {
		iterationNumber = MaxIterations
	}

	for i := 0; i < iterationNumber; i++ {
		value = Step(value, number)
	}

	return value
}

// SqrtPrecision This function is used for counting square root of number
// with custome iteration number
func SqrtPrecision(number, precision float64) (value float64, iterationNumber int) {
	value = float64(1.0)
	realValue := math.Sqrt(number)

	for ; iterationNumber < MaxIterations && math.Abs(value-realValue) > precision; iterationNumber++ {
		value = Step(value, number)
	}

	return value, iterationNumber
}

func main() {
	number := float64(8.0)
	precision := float64(1 >> 10)
	result, iterationNumber := SqrtPrecision(number, precision)

	fmt.Println()
	fmt.Printf("math.Sqrt\t\t: %v\n", math.Sqrt(number))
	fmt.Printf("SqrtIteration result\t: %v, after: %v iterations\n", result, iterationNumber)
	fmt.Printf("SqrtIteration\t\t: %v\n", SqrtIteration(number, 5))
	fmt.Printf("Sqrt\t\t\t: %v\n\n", Sqrt(number))
	fmt.Println("Pierwszy program Go Jacka")
	fmt.Printf("Liczba do approxymacji: %v\n", number)
}
