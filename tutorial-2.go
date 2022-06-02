package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	var errorVar ErrNegativeSqrt
	var z float64 = 1.0

	if x < 0 {
		return z, errorVar
	}

	previousZ := 0.0
	for math.Abs(z-previousZ) > 0.00001 {
		previousZ = z
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Z= %v\n", z)
		fmt.Printf("Diff=%v\n", math.Abs(previousZ-z))
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
