package main

import (
	"fmt"
	"math"
)

type Shape2D interface {
	Perimeter() float64
}

type Circle struct {
	R float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}

func main() {
	c := Circle{R: 1.5}
	fmt.Printf("R %.2f -> Perimeter %.3f \n", c.R, c.Perimeter())
	_, ok := interface{}(c).(Shape2D)

	if ok {
		fmt.Println("Shape2D")
	} else {
		fmt.Println("Not Shape2D")
	}
}
