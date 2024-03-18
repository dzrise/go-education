package pq

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5"
)

func main() {
	arguments := os.Args
	if len(arguments) != 6 {
		fmt.Println("Please provide: hostname, port, username, password, database")
		return
	}

	host := arguments[1]
	port := arguments[2]
	user := arguments[3]
	password := arguments[4]
	database := arguments[5]

	port, err := strconv.Atoi(port)

	if err != nil {
		fmt.Println("Not a valid port:", err)
		return
	}

	conn_string := fmt.Printf("postgres://%s:%s@%s:%s/%s", user, password, host, port, database)
	conn, err := pgx.Connect(context.Background(), conn_string)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}
