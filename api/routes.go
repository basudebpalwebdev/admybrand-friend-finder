package api

import (
	"github.com/basudebpalwebdev/admybrand-friend-finder/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func (server *Server) RouteSetup() {
	server.api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	server.api.Get("/users", handlers.ListUsers)
}
