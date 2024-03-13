package main

import (
	"fmt"
	"sort"
)

type S1 struct {
	Name string
}

type IS2 interface {
	getName() string
}

type sliceS2 []IS2

func (s S1) getName() string {
	return s.Name
}

func (s sliceS2) Len() int {
	return len(s)
}

func (s sliceS2) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sliceS2) Less(i, j int) bool {
	return s[i].getName() < s[j].getName()
}

func main() {
	ss := sliceS2{S1{Name: "qq"}, S1{Name: "lk"}}
	fmt.Println(ss)
	sort.Sort(ss)
	fmt.Println(ss)
	sort.Sort(sort.Reverse(ss))
	fmt.Println(ss)
}
