package main

import (
	"fmt"
	"sort"
)

func main() {
	sInts := []int{1, 0, 2, -3, 4, -20}
	sFloats := []float64{1.0, 0.2, 0.22, -3, 4.1, -0.1}
	sSrings := []string{"aa", "a", "A", "Aa", "aab", "AAa"}

	fmt.Println("Original sInts:", sInts)
	sort.Ints(sInts)
	fmt.Println("sInts:", sInts)
	sort.Sort(sort.Reverse(sort.IntSlice(sInts)))
	fmt.Println("Reverse:", sInts)

	fmt.Println("sFloats original:", sFloats)
	sort.Float64s(sFloats)
	fmt.Println("sFloats new:", sFloats)
	sort.Sort(sort.Reverse(sort.Float64Slice(sFloats)))
	fmt.Println("Reverse:", sFloats)

	fmt.Println("sStrings original:", sSrings)
	sort.Strings(sSrings)
	fmt.Println("sStrings", sSrings)
	sort.Sort(sort.Reverse(sort.StringSlice(sSrings)))
	fmt.Println("Reverse", sSrings)

}
