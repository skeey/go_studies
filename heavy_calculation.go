package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func calc(base float64, control *sync.WaitGroup) {
	defer control.Done()
	n := 0.0
	
	for i := 0; i < 100000000; i++ {
		n += base / math.Pi * math.Sin(2)
	}

	fmt.Println(n)
}

func main() {
	start := time.Now()
	var control sync.WaitGroup
	control.Add(3)

	go calc(9.37, &control)
	go calc(6.94, &control)
	go calc(42.57, &control)

	control.Wait()
	fmt.Printf("Finished in %s.\n", time.Since(start))
}
