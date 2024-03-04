package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Printf("Need one or more arguments!\n")
		return
	}

	var min, max float64
	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)

		if err != nil {
			continue
		}

		if i == 1 {
			min = n
			max = n
			continue
		}

		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	fmt.Println("min: ", min)
	fmt.Println("max: ", max)
}
