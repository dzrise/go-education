package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Need more arguments")
		return
	}

	n, err := strconv.Atoi(arguments[1])

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	input := strconv.Itoa(n)
	fmt.Printf("Type: %T\n", input)
	input = strconv.FormatInt(int64(4), 10)
	fmt.Printf("Type: %T\n", input)
	input = string(n)
	fmt.Printf("Type: %T\n", input)
}
