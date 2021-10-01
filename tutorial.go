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
========================================
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


========================================

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	for y := range(result) {
		row := make([]uint8, dx)
		for x := range(row) {
			row[x] = uint8(x ^ y)
		}
		result[y] = row
	}
	return result
}

func main() {
	pic.Show(Pic)
}

========================================
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)

	words := strings.Fields(s)

	for i := range words {

		var word string
		var currentCount int
		word = words[i]
		currentCount = result[word]
		result[word] = currentCount + 1

	}

	return result
}

func main() {
	wc.Test(WordCount)
}

========================================
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	fibA := 0
	fibB := 0
	return func(i int) int {
		if i == 0 {
			fibA = 0
			fibB = 0
		} else if i == 1 {
			fibA = 0
			fibB = 1
		} else {
			fib := (fibB) + (fibA)
			fibA = fibB
			fibB = fib
		}
		return fibB
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}

====================================
package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}


func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}