package main

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"
)

type Entery struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entery{}
var MIN = 0
var MAX = 26

func search(key string) *Entery {
	for i, v := range data {
		if v.Tel == key {
			return &data[i]
		}
	}
	return nil
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(l int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == l {
			break
		}
		i++
	}
	return temp
}

func populate(n int, s []Entery) {
	for i := 0; i < n; i++ {
		name := getString(4)
		suname := getString(5)
		n := strconv.Itoa(random(100, 199))

		data = append(data, Entery{name, suname, n})
	}
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

	SEED := time.Now().Unix()
	rand.Seed(SEED)

	// How many records to insert
	n := 100
	populate(n, data)
	fmt.Printf("Data has %d entries.\n", len(data))

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
