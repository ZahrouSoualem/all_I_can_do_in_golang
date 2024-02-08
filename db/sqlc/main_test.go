package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	DB_DRIVER = "postgres"
	DB_SOURCE = "postgres://root:secret@localhost:5432/testerdb?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	//confg, err := util.LoadConfig("./.")
	conx, err := sql.Open(DB_DRIVER, DB_SOURCE)

	if err != nil {
		log.Fatal("cannot connect to db ", err)
	}

	testQueries = New(conx)

	os.Exit(m.Run())

}
