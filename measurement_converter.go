package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("You need to pass <values> <unit> as parameters")
		os.Exit(1)
	}
	
	originUnit := os.Args[len(os.Args)-1]
	originValues := os.Args[1 : len(os.Args) - 1]

	var destUnit string

	if originUnit == "celsius" {
		destUnit = "fahrenheit"
	} else if originUnit == "kilometers" {
		destUnit = "miles"
	} else {
		fmt.Printf("%s is not a known unit", originUnit)
		os.Exit(1)
	}

	for i, v := range originValues {
		originValue, err := strconv.ParseFloat(v, 64)

		if err != nil {
			fmt.Printf("The value %s in position %d is not a valid number!\n", v, i)
			os.Exit(1)
		}

		var destValue float64

		if originUnit == "celsius" {
			destValue = originValue * 1.8 + 32
		} else {
			destValue = originValue / 1.60934
		}

		fmt.Printf("%.2f %s = %.2f %s\n", originValue, originUnit, destValue, destUnit)
	}
}
