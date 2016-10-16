package main

import (
	"fmt"
	"math"
)

// MaxIterations Setting up maximum number of iterations withing application
const MaxIterations int = 99999

// Sqrt x
func Sqrt(number float64) float64 {
	value := float64(1.0)

	for i := 0; i < 5; i++ {
		value = value - (((value * value) - number) / (2 * (value)))
	}

	return value
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
		value = value - (((value * value) - number) / (2 * (value)))
	}

	return value
}

func main() {
	number := float64(8.0)

	fmt.Println()
	fmt.Printf("math.Sqrt\t\t: %v\n", math.Sqrt(number))
	fmt.Printf("SqrtIteration\t\t: %v\n", SqrtIteration(number, 47))
	fmt.Printf("SqrtIteration\t\t: %v\n", SqrtIteration(number, 45))
	fmt.Printf("SqrtIteration\t\t: %v\n", SqrtIteration(number, 43))
	fmt.Printf("SqrtIteration\t\t: %v\n", SqrtIteration(number, 40))
	fmt.Printf("SqrtIteration\t\t: %v\n", SqrtIteration(number, 30))
	fmt.Printf("SqrtIteration\t\t: %v\n", SqrtIteration(number, 20))
	fmt.Printf("SqrtIteration\t\t: %v\n", SqrtIteration(number, 10))
	fmt.Printf("SqrtIteration\t\t: %v\n", SqrtIteration(number, 8))
	fmt.Printf("SqrtIteration\t\t: %v\n", SqrtIteration(number, 5))
	fmt.Printf("SqrtIteration\t\t: %v\n", SqrtIteration(number, 3))
	fmt.Printf("Sqrt\t\t\t: %v\n\n", Sqrt(number))
	fmt.Println("Pierwszy program Go Jacka")
	fmt.Printf("Liczba do approxymacji: %v\n", number)
}
