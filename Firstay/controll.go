package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <int>\n", os.Args[0])
		return
	}

	argument := os.Args[1]

	switch argument {
	case "0":
		fmt.Println("Zero!")
	case "1":
		fmt.Println("One!")
	case "2", "3", "4":
		fmt.Println("2 or 3 or 4!")
	default:
		fmt.Printf("I don't understand %s.\n", argument)
	}

	value, err := strconv.Atoi(argument)

	if err != nil {
		fmt.Println("Can't convert to int:", argument)
		return
	}

	switch {
	case value == 0:
		fmt.Println("Zero!")
	case value > 0:
		fmt.Println("Positive integer!")
	case value < 0:
		fmt.Println("Negative integer!")
	default:
		fmt.Println("This should not happen:", value)
	}

}
