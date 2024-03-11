package main

import (
	"fmt"
)

func returnNumber() interface{} {
	return 12
}

func isBool(x interface{}) bool {
	switch x.(type) {
	case bool:
		return true
	default:
		return false
	}
}

func main() {
	antInt := returnNumber()
	number := antInt.(int)
	number++
	fmt.Println(number)
	value, ok := antInt.(int64)
	if ok {
		fmt.Println("Type association successful:", value)
	} else {
		fmt.Println("Type association failed")
	}

	i := antInt.(int)
	fmt.Println("i: ", i)
	test := isBool(antInt)
	fmt.Println(test)
	if test {
		_ = antInt.(bool)
	}

}
