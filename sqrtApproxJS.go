package main

import (
	"fmt"
	"math"
)

//  Sqrt x
func Sqrt(number float64) float64 {
	value := float64(1.0)

	for i := 0; i < 5; i++ {
		value = value - (((value * value) - number) / (2 * (value)))
	}

	return value
}

func main() {
	number := float64(8.0)

	fmt.Println(math.Sqrt(number))
	fmt.Println(Sqrt(number))
	fmt.Println("Pierwszy program Go Jacka")
	fmt.Printf("Liczba do approxymacji: %v", number)
}
