package main

import "fmt"

func main() {
	a := make([]int, 4)
	fmt.Println("L:", len(a), "C:", cap(a))

	b := []int{0, 1, 2, 3, 4}
	fmt.Println("L:", len(b), "C:", cap(b))

	arSlise := make([]int, 4, 4)
	fmt.Println(arSlise)
	arSlise = append(arSlise, 5)
	fmt.Println(arSlise)
	fmt.Println("L:", len(arSlise), "C:", cap(arSlise))

	arSlise = append(arSlise, []int{-1, -2, -3, -4}...)
	fmt.Println(arSlise)
	fmt.Println("L:", len(arSlise), "C:", cap(arSlise))
}
