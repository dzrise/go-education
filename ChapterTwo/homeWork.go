package main

import (
	"fmt"
	"reflect"
)

func TwoArrayInSlice(arr1, arr2 any) []int {
	arSlise := make([]int, 0)
	for i := 0; i < len(arr1); i++ {
		arSlise = append(arSlise, arr1[i])
	}
	for i := 0; i < len(arr2); i++ {
		arSlise = append(arSlise, arr2[i])
	}
	return arSlise
}

func TwoArrayInArray(arr1, arr2 [10]int) interface{} {
	arSlise := TwoArrayInSlice(arr1, arr2)

	s := reflect.ValueOf(arSlise)
	t := reflect.ArrayOf(s.Len(), s.Type().Elem())
	a := reflect.New(t).Elem()
	reflect.Copy(a, s)
	return a.Interface()
}

func TwoSliceInArray(arr1, arr2 []int) interface{} {
	arSlice := make([]int, len(arr1)+len(arr2))
	for i, v := range arr1 {
		arSlice[i] = v
	}

	for i, v := range arr2 {
		arSlice[i] = v
	}
	s := reflect.ValueOf(arSlice)
	t := reflect.ArrayOf(s.Len(), s.Type().Elem())
	a := reflect.New(t).Elem()
	reflect.Copy(a, s)
	return a.Interface()
}

func main() {
	arr1 := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr2 := [10]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	fmt.Println("arr1 type:", reflect.ValueOf(arr1).Type())
	fmt.Println("arr1:", arr1)
	fmt.Println("arr2", arr2)

	res := TwoArrayInSlice(arr1, arr2)
	fmt.Println("result updateArrayInSlice:", res)
	res2 := TwoArrayInArray(arr1, arr2)
	fmt.Println("result TwoArrayInArray:", res2)
	fmt.Println("res2 type:", reflect.ValueOf(res2).Type())
	res3 := TwoSliceInArray(arr1[:], arr2[:])
	fmt.Println("result TwoArrayInArray:", res3)
	fmt.Println("res3 type:", reflect.ValueOf(res3).Type())

}
