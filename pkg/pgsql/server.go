package main

import (
	"github.com/codenotary/immudb/pkg/pgsql/server"
	"log"
)

func main() {

	sqlServer := server.New(server.Port("5439"))

	err := sqlServer.Serve()

	if err != nil {
		log.Fatal(err)
	}
}
