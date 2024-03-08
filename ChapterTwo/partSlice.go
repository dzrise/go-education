package main

import "fmt"

func main() {
	arSlise := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(arSlise)
	i := len(arSlise)
	fmt.Println(i)
	fmt.Println(arSlise[0:5])
	fmt.Println(arSlise[:5])
	fmt.Println(arSlise[5:])

	fmt.Println(arSlise[i-2 : i])
	fmt.Println(arSlise[i-2:])

	t := arSlise[0:5:10]
	fmt.Println(len(t), cap(t))
	fmt.Println(t)

	t = arSlise[:5:6]
	fmt.Println(len(t), cap(t))
	fmt.Println(t)

}
