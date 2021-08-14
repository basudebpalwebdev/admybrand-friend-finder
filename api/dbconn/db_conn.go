package dbconn

import (
	"database/sql"

	db "github.com/basudebpalwebdev/admybrand-friend-finder/db/sqlc"
)

var (
	DBQueries *db.Queries
	DBConn    *sql.DB
)
