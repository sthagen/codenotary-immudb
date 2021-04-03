package server

import (
	"database/sql"
	"fmt"
	_ "github.com/lxn/go-pgsql"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestSrv_Serve(t *testing.T) {

	sqlServer := New(Port("5439"))

	go func() {
		err := sqlServer.Serve()
		require.NoError(t, err)
	}()
	time.Sleep(100 * time.Millisecond)

	db, err := sql.Open("postgres", "host=localhost port=5439 dbname=testdatabase user=testuser password=testpassword")
	require.NoError(t, err)

	defer db.Close()

	var msg string

	err = db.QueryRow("SELECT $1 || ' ' || $2;", "Hello", "SQL").Scan(&msg)
	require.NoError(t, err)

	fmt.Println(msg)

}
