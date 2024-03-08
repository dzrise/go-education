package main

import "fmt"

func main() {
	arSlise := []float64{}

	fmt.Println(arSlise, len(arSlise), cap(arSlise))

	arSlise = append(arSlise, 1234.56)
	arSlise = append(arSlise, -34.0)
	fmt.Println(arSlise, "with length", len(arSlise))

	t := make([]int, 4)
	t[0] = -1
	t[1] = -2
	t[2] = -3
	t[3] = -4

	t = append(t, -5)
	fmt.Println(t)

	twoD := [][]int{{1, 2, 3}, {4, 5}}
	fmt.Println(twoD, "with length", len(twoD))
	for _, i := range twoD {
		for _, k := range i {
			fmt.Print(k, " ")
		}
		fmt.Println()
	}

	make2D := make([][]int, 2)
	fmt.Println(make2D)
	make2D[0] = []int{1, 2, 3, 4}
	make2D[1] = []int{-1, -2, -3, -4}
	fmt.Println(make2D)

}
