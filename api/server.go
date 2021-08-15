package api

import (
	"database/sql"

	"github.com/basudebpalwebdev/admybrand-friend-finder/api/dbconn"
	db "github.com/basudebpalwebdev/admybrand-friend-finder/db/sqlc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	api   *fiber.App
	store db.Store
}

func NewServer(dbConn *sql.DB) *Server {
	store := db.NewStore(dbConn)
	server := &Server{
		api:   fiber.New(),
		store: store,
	}
	dbconn.DBConn = dbConn
	dbconn.DBQueries = db.New(dbconn.DBConn)
	server.api.Use(logger.New())
	server.RouteSetup()
	return server
}
func (server *Server) Start(address string) error {
	return server.api.Listen(address)
}
func (server *Server) Stop() error {
	return server.api.Shutdown()
}
