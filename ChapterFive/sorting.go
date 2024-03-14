package main

import (
	"fmt"
	"sort"
)

type Grades struct {
	Name   string
	Suname string
	Grade  int
}

func main() {
	data := []Grades{{"J", "Levis", 10}, {"M.", "Tsou", 7}, {"N", "Tsou", 20}}

	isSorted := sort.SliceIsSorted(data, func(i, j int) bool {
		return data[i].Grade < data[j].Grade
	})

	if isSorted {
		fmt.Println("It is sorted")
	} else {
		fmt.Println("It is not sorted")
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].Grade < data[j].Grade
	})

	fmt.Println("By Grade:", data)

}
