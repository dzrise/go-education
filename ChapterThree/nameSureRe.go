package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchNameSur(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}

func main() {
	arguments := os.Args[1:]
	fmt.Println(arguments)
	if len(arguments) == 0 {
		fmt.Printf("Need one or more arguments!\n")
		return
	}

	res := matchNameSur(arguments[0])
	fmt.Println(res)

}
