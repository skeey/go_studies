package main

import (
	"fmt"
	"time"
)

func main() {
	passedTime(generateFibonacci(8))
	passedTime(generateFibonacci(48))
	passedTime(generateFibonacci(88))
}

func passedTime(function func()) {
	start := time.Now()

	function()

	fmt.Printf("\nRuntime: %s\n\n", time.Since(start))
}

func generateFibonacci(n int) func() {
	return func() {
		a, b := 0, 1

		fib := func() int {
			a, b = b, a+b

			return a
		}

		for i := 0; i < n; i++ {
			fmt.Printf("%d ", fib())
		}
	}
}
