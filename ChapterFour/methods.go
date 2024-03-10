package main

import (
	"fmt"
	"os"
	"strconv"
)

type ar2x2 [2][2]int

func Add(a, b ar2x2) ar2x2 {
	c := ar2x2{}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = a[i][j] + b[i][j]
		}
	}
	return c
}

func (a *ar2x2) add(b ar2x2) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			a[i][j] += b[i][j]
		}
	}
}

func (a *ar2x2) Substract(b ar2x2) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			a[i][j] = a[i][j] - b[i][j]
		}
	}
}

func (a *ar2x2) Multiply(b ar2x2) {
	a[0][0] = a[0][0]*b[0][0] + a[0][1]*b[1][0]
	a[1][0] = a[1][0]*b[0][0] + a[1][1]*b[1][0]
	a[0][1] = a[0][0]*b[0][1] + a[0][1]*b[1][1]
	a[1][1] = a[1][0]*b[0][1] + a[1][1]*b[1][1]
}

func main() {
	if len(os.Args) != 9 {
		fmt.Printf("Need 8 arguments!\n")
		return
	}

	k := [8]int{}
	for index, i := range os.Args[1:] {
		v, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println(err)
			return
		}
		k[index] = v
	}
	a := ar2x2{{k[0], k[1]}, {k[2], k[3]}}
	b := ar2x2{{k[4], k[5]}, {k[6], k[7]}}

	fmt.Println("Traditional add:", Add(a, b))
	a.add(b)
	fmt.Println("a+b:", a)
	a.Substract(b)
	fmt.Println("a-b:", a)
	a.Multiply(b)
	a = ar2x2{{k[0], k[1]}, {k[2], k[3]}}
	a.Multiply(b)
	fmt.Println("a*b", a)
	a = ar2x2{{k[0], k[1]}, {k[2], k[3]}}
	b.Multiply(a)
	fmt.Println("b*a", b)
}
