package main

import "fmt"

type IntA interface {
	foo()
}

type IntB interface {
	boo()
}

type IntC interface {
	IntA
	IntB
}

func processA(a IntA) {
	fmt.Printf("%T\n", a)
}

type a struct {
	XX int
	YY int
}

func (varC c) foo() {
	fmt.Println("Foo proccessing", varC)
}

func (varC c) boo() {
	fmt.Println("Boo proccessing", varC)
}

type b struct {
	AA string
	XX int
}

type c struct {
	A a
	B b
}

type compose struct {
	field1 int
	field2 a
}

func (A a) A() {
	fmt.Println("Function A() for A")
}

func (B b) B() {
	fmt.Println("Function B() for B")
}

func main() {
	var iC = c{a{120, 12}, b{"-12", -12}}
	iC.A.A()
	iC.B.B()

	iComp := compose{123, a{456, 789}}
	fmt.Println(iComp.field2.XX, iComp.field2.YY, iComp.field1)
	iC.boo()
	processA(iC)

}
