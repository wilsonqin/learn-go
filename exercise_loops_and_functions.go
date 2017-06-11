package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	last_z := 0.0
	
	for math.Abs(z - last_z) >= 0000.1 {
		last_z = z
		z = last_z - ((math.Pow(last_z,2.) - x) / (2. * last_z))
	}

	return z
}

func RunSqrt() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v Squared: Guess %v,", i, Sqrt(float64(i)))
		fmt.Printf(" Actual %v\n", math.Sqrt(float64(i)))
	}
}

