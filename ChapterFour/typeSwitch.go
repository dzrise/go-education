package main

import (
	"fmt"
)

type Secret struct {
	SecretValue string
}

type Entry struct {
	F1 int
	F2 string
	F3 Secret
}

func testStruct(x interface{}) {
	switch T := x.(type) {
	case Secret:
		fmt.Println("secret type")
	case Entry:
		fmt.Println("entry type")
	default:
		fmt.Printf("Not supported type: %T\n", T)
	}
}

func learn(x interface{}) {
	switch T := x.(type) {
	default:
		fmt.Printf("Data type: %T\n", T)
	}
}

func main() {
	A := Entry{F1: 100, F2: "F2", F3: Secret{"password"}}
	testStruct(A)
	testStruct(A.F3)
	testStruct("A string")
	learn(12)
	learn("€")
}
