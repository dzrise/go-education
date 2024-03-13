package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

const min = 1
const max = 5

func random64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

type Shape3D interface {
	Vol() float64
}

type Cube struct {
	x float64
}

type Cuboid struct {
	x float64
	y float64
	z float64
}

type Sphere struct {
	r float64
}

func (c Cube) Vol() float64 {
	return c.x * c.x * c.x
}

func (c Cuboid) Vol() float64 {
	return c.x * c.y * c.z
}

func (c Sphere) Vol() float64 {
	return 4 * math.Pi * c.r * c.r * c.r
}

type Shapes []Shape3D

func (a Shapes) Len() int { return len(a) }

func (a Shapes) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a Shapes) Less(i, j int) bool {
	return a[i].Vol() < a[j].Vol()
}

func PrintlnShapes(a Shapes) {
	for _, s := range a {
		switch s.(type) {
		case Cube:
			fmt.Printf("Cube Volume %.2f\n", s.Vol())
		case Cuboid:
			fmt.Printf("Cuboid Volume %.2f\n", s.Vol())
		case Sphere:
			fmt.Printf("Sphere Volume %.2f\n", s.Vol())
		default:
			fmt.Printf("Unknown Shape")
		}

	}
	fmt.Println()
}

func main() {
	data := Shapes{}
	rand.Seed(time.Now().Unix())

	for i := 0; i < 3; i++ {
		cube := Cube{random64(min, max)}
		cuboid := Cuboid{random64(min, max), random64(min, max), random64(min, max)}
		shpere := Sphere{random64(min, max)}
		data = append(data, cube)
		data = append(data, cuboid)
		data = append(data, shpere)

	}

	PrintlnShapes(data)

	sort.Sort(Shapes(data))
	PrintlnShapes(data)
}
