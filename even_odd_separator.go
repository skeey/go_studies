package main

import (
	"fmt"
)

func main() {
	evenChan, oddChan := make(chan int), make(chan int)
	finishedChan := make(chan bool)
	numbers := []int{1, 23, 42, 5, 8, 6, 7, 4, 99, 100}

	go splitEvenOdd(numbers, evenChan, oddChan, finishedChan)

	var evenNumbers, oddNumbers []int
	finished := false

	for !finished {
		select {
			case n := <-evenChan:
				evenNumbers = append(evenNumbers, n)
			case n := <-oddChan:
				oddNumbers = append(oddNumbers, n)
			case finished = <-finishedChan:
		}
	}

	fmt.Printf("Even numbers: %v | Odd numbers: %v\n", evenNumbers, oddNumbers)
}

func splitEvenOdd(numbers []int, evenChan, oddChan chan<- int, finishedChan chan<- bool) {
	for _, number := range numbers {
		if number % 2 == 0 {
			evenChan <- number
		} else {
			oddChan <- number
		}
	}
	finishedChan <- true
}
