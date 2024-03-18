package main

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	arguments := os.Args
	if len(os.Args) != 6 {
		fmt.Println("Please provide: hostname, port, username, password, database")
		return
	}
}
