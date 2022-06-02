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
