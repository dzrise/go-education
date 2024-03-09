package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func matchNameSur(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}

func matchtel(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match(t)
}

func matchRecodrs(recodrs string) bool {
	fields := strings.Split(recodrs, ",")
	if len(fields) != 3 {
		return false
	}

	if !matchNameSur(fields[0]) {
		return false
	}

	if !matchNameSur(fields[1]) {
		return false
	}

	return matchtel(fields[2])
}

func main() {
	arguments := os.Args[1:]
	fmt.Println(arguments)
	if len(arguments) != 3 {
		fmt.Printf("Need 3 arguments!\n")
		return
	}

	res := matchRecodrs(strings.Join(arguments, ","))
	fmt.Println(res)

}
