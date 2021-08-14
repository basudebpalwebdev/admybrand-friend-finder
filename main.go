package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/basudebpalwebdev/admybrand-friend-finder/api"
)

func main() {
	dbConn, err := sql.Open("postgres", "postgresql://basu:Basudeb@2021@localhost:5432/admybrand_friend_finder?sslmode=disable")
	if err != nil {
		log.Fatal("Cannot connect to the database :", err)
	}
	webServer := api.NewServer(dbConn)
	webServer.Start("0.0.0.0:9999")
}
