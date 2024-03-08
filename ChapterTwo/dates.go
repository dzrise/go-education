package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	startDate := time.Now()

	if len(os.Args) != 2 {
		fmt.Println("Usage: dates parse_string")
		return
	}

	dateString := os.Args[1]
	d, err := time.Parse("02 January 2006", dateString)

	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Timeear:", d.Day(), d.Month(), d.Year())
	}

	d, err = time.Parse("02 Junuary 2006 15:04", dateString)

	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Date:", d.Day(), d.Month(), d.Year())
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	d, err = time.Parse("02-01-2006 15:04", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Date:", d.Day(), d.Month(), d.Year())
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	d, err = time.Parse("15:04", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	t := time.Now().Unix()
	fmt.Println("Epoch Time:", t)

	d = time.Unix(t, 0)
	fmt.Println("Date:", d.Year(), d.Month(), d.Day())
	fmt.Printf("Time: %d:%d\n", d.Hour(), d.Minute())
	duration := time.Since(startDate)
	fmt.Println("Duration:", duration)

}
