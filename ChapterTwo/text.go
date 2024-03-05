package main

import (
	"fmt"
)

func main() {
	aString := "Hello, World! €"
	fmt.Println("First character of aString: ", string(aString[0]))
	r := '€'
	fmt.Println("As an int32 value: ", r)
	fmt.Printf("As a string: %s and as a character: %c\n", r, r)

	for _, c := range aString {
		fmt.Printf("%x ", c)
	}
	fmt.Println()
	for _, c := range aString {
		fmt.Printf("%c", c)
	}

	fmt.Println()
}
