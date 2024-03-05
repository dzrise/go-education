package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i*i, " ")
	}

	fmt.Println()

	i := 0
	for ok := true; ok; ok = (i != 10) {
		fmt.Print(i*i, " ")
		i++
	}

	fmt.Println()

	i = 0
	for {
		if i == 10 {
			break
		}
		fmt.Print(i*i, " ")
		i++
	}

	fmt.Println()

	arSlice := []int{-1, 2, 1, -1, 2, -2}
	for i, v := range arSlice {
		fmt.Println("index: ", i, " value: ", v)
	}
}
