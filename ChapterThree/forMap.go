package main

import "fmt"

func main() {
	aMap := make(map[string]string)
	aMap["123"] = "456"
	aMap["key"] = "value"

	for k, v := range aMap {
		fmt.Println("key:", k, "value:", v)
	}

	for _, v := range aMap {
		fmt.Print(" # ", v)
	}
	fmt.Println()
}
