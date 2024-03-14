package main

import "fmt"

func funRet(i int) func(int) int {
	if i < 0 {
		return func(x int) int {
			x = -x
			return x + x
		}
	}

	return func(x int) int {
		return x * x
	}
}

func main() {
	n := 10
	i := funRet(n)
	j := funRet(-4)
	fmt.Println("%T\n", i)
	fmt.Println("%T %v\n", j, j)
	fmt.Println("j", j, j(-5))
}
