package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func check(a, b int) error {
	if a == 0 && b == 0 {
		return errors.New("this custom error message")
	}

	return nil
}

func formattedError(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("a %d and b %d, UserId: %d", a, b, os.Geteuid())
	}

	return nil
}

func main() {
	err := check(0, 10)
	if err == nil {
		fmt.Println("check() ended normally")
	} else {
		fmt.Println(err)
	}

	err = check(0, 0)
	if err.Error() == "this custom error message" {
		fmt.Println("custom error detected")
	}

	err = formattedError(0, 0)

	if err != nil {
		fmt.Println(err)
	}

	i, err := strconv.Atoi("-124")
	if err == nil {
		fmt.Println("Int value is", i)
	}

	i, err = strconv.Atoi("Y124")
	if err != nil {
		fmt.Println(err)
	}
}
