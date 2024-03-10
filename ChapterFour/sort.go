package main

import (
	"fmt"
	"sort"
)

type S1 struct {
	F1 int
	F2 string
	F3 int
}

type S2 struct {
	F1 int
	F2 string
	F3 S1
}

type S2slice []S2

func (s S2slice) Len() int {
	return len(s)
}

func (s S2slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s S2slice) Less(i, j int) bool {
	return s[i].F1 < s[j].F1
}

func main() {
	data := []S2{
		S2{1, "One", S1{1, "S1_1", 10}},
		S2{-1, "Two", S1{2, "S1_2", 20}},
		S2{3, "Three", S1{-1, "S1_3", -20}},
	}

	fmt.Println("Before: ", data)
	sort.Sort(S2slice(data))
	fmt.Println("After: ", data)
	sort.Sort(sort.Reverse(S2slice(data)))
	fmt.Println("After Reverse: ", data)
}
