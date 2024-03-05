package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var total, nInts, nFloats int
	invalid := make([]string, 0)

	arguments := os.Args

	for _, k := range arguments[1:] {
		_, err := strconv.Atoi(k)

		if err == nil {
			total++
			nInts++
			continue
		}

		_, err = strconv.ParseFloat(k, 64)

		if err == nil {
			total++
			nFloats++
			continue
		}

		invalid = append(invalid, k)

	}

	fmt.Printf("#reed %d #ints %d #floats %d\n", total, nInts, nFloats)

	if len(invalid) > 1 {
		fmt.Printf("Too much invalid input %d\n", len(invalid))
		for _, k := range invalid {
			fmt.Printf("%s\n", k)
		}
	}

}
