package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: dates parse_string")
		return
	}

	dateString := os.Args[1]

	fmt.Println("dateString: ", dateString)
	now, err := time.Parse("02 January 2006 15:04 MST", dateString)
	if err != nil {
		panic(err)
	}

	loc, _ := time.LoadLocation("America/New_York")
	fmt.Printf("New York Time: %s\n", now.In(loc))
}
