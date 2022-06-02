package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z float64 = 1.0
	previousZ := 0.0
	for math.Abs(z-previousZ) > 0.00001 {
		previousZ = z
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Z= %v\n", z)
		fmt.Printf("Diff=%v\n", math.Abs(previousZ-z))
	}
	return z
}

func main() {
	fmt.Println(Sqrt(131393))
	fmt.Println(math.Sqrt(131393))
}
