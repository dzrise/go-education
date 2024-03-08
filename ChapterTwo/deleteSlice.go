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

	index := arguments[1]
	i, err := strconv.Atoi(index)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Usage index:", i)

	arSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Original slice:", arSlice)

	if i > len(arSlice) {
		fmt.Println("Can not delete element", i)
		return
	}

	arSlice = append(arSlice[:i], arSlice[i+1:]...)

	fmt.Println("After 1st deletion:", arSlice)

	if i > len(arSlice)-1 {
		fmt.Println("Can not delete element", i)
		return
	}

	arSlice[i] = arSlice[len(arSlice)-1]
	arSlice = arSlice[:len(arSlice)-1]
	fmt.Println("After 2st deletion:", arSlice)

}
