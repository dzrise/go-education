package main

import "fmt"

func updateArrayInSlice(arr1 []int, arr2 []int) []int {
	arSlise := make([]int, 0)
	arSlise = append(arSlise, arr1...)
	arSlise = append(arSlise, arr2...)
	return arSlise
}

func main() {
	arr1 := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr2 := [10]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}

	fmt.Println("arr1:", arr1)
	fmt.Println("arr2", arr2)

	res := updateArrayInSlice(arr1, arr2)
	fmt.Println("result updateArrayInSlice:", res)
	res = updateArrayInArray(arr1, arr2)
	fmt.Println("result updateArrayInArray:", res)
}
