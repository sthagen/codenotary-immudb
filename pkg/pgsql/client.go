package main

import (
	"database/sql"
	"fmt"
	"log"
)

import (
	_ "github.com/lxn/go-pgsql"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5439 dbname=testdatabase user=testuser password=testpassword")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var msg string

	err = db.QueryRow("SELECT $1 || ' ' || $2;", "Hello", "SQL").Scan(&msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)
}
