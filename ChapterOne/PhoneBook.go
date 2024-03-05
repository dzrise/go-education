package main

import (
	"fmt"
	"os"
	"path"
)

type Entery struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entery{}

func search(key string) *Entery {
	for i, v := range data {
		if v.Surname == key {
			return &data[i]
		}
	}

	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
		return
	}
	data = append(data, Entery{"Denis", "Denisov", "06666666667"})
	data = append(data, Entery{"Alexey", "Alexeyov", "06666666668"})
	data = append(data, Entery{"Ivan", "Ivanov", "06666666666"})

	switch arguments[1] {
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}

		result := search(arguments[2])

		if result == nil {
			fmt.Println("Entery not found: ", arguments[2])
			return
		}

		fmt.Println(*result)

	case "list":

		list()

	default:
		fmt.Println("Not a valid option")
	}

}
