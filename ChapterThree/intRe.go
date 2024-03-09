package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchInt(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match(t)
}

func main() {
	arguments := os.Args[1:]
	fmt.Println(arguments)
	if len(arguments) == 0 {
		fmt.Printf("Need one or more arguments!\n")
		return
	}

	res := matchInt(arguments[0])
	fmt.Println(res)

}
