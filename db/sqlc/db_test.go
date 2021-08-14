package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	testQueries *Queries
	testDB      *sql.DB
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open("postgres", "postgresql://basu:Basudeb@2021@localhost:5432/admybrand_friend_finder?sslmode=disable")
	if err != nil {
		log.Fatal("Cannot connect to the database :", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
