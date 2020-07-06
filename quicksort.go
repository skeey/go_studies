package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	inputs := os.Args[1:]
	numbers := make([]int, len(inputs))
	
	for i, n := range inputs {

		number, err := strconv.Atoi(n)

		if err != nil {
			fmt.Printf("%s is not a valid number!\n", n)
			os.Exit(1)
		}

		numbers[i] = number

	}

	fmt.Println(quicksort(numbers))

}

func quicksort(numbers []int) []int {
	
	if len(numbers) <= 1 {
		return numbers
	}

	n := make([]int, len(numbers))
	copy(n, numbers)

	pivotIndex := len(n) / 2
	pivot := n[pivotIndex]

	n = append(n[:pivotIndex], n[pivotIndex+1:]...)

	minors, more := partition(n, pivot)
	
	return append(
		append(quicksort(minors), pivot),
		quicksort(more)...
	)

}

func partition(numbers []int, pivot int) (minors []int, more []int) {
	
	for _, n := range numbers {
	
		if n <= pivot {
			minors = append(minors, n)
		} else {
			more = append(more, n)
		}

	}

	return minors, more

}
