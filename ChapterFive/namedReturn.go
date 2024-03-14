package main

import (
	"fmt"
	"os"
	"strconv"
)

func minMax(a, b int) (min, max int) {
	if a < b {
		min = a
		max = b
		return min, max
	}

	min = b
	max = a
	return min, max
}

func main() {
	arguments := os.Args[1:]

	if len(arguments) != 2 {
		fmt.Println("Need 2 arguments!")
		return
	}
	min, err := strconv.Atoi(arguments[0])

	if err != nil {
		fmt.Println(err)
		return
	}

	max, err := strconv.Atoi(arguments[1])

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(min, max)
	min, max = minMax(min, max)

	fmt.Println(min, max)

}
