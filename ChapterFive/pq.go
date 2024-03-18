package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgx"
)

func main() {
	arguments := os.Args
	if len(arguments) != 6 {
		fmt.Println("Please provide: hostname, port, username, password, database")
		return
	}

	host := arguments[1]
	p := arguments[2]
	user := arguments[3]
	password := arguments[4]
	database := arguments[5]

	port, err := strconv.Atoi(p)

	if err != nil {
		fmt.Println("Not a valid port:", err)
		return
	}

	conn_string := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, database)
	conn, err := pgx.Connect(context.Background(), conn_string)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), `SELECT "datname" FROM "pg_database" WHERE datistemplate = false`)
	if err != nil {
		fmt.Println("Query", err)
		return
	}

	cols, err := rows.Columns()
	if err != nil {
		panic(err)
	}
	println("pq columns:")
	fmt.Println(cols)
}
