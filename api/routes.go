package api

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/basudebpalwebdev/admybrand-friend-finder/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func (server *Server) RouteSetup() {
	server.api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	server.api.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "/swagger/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
	}))
	server.api.Get("/users", handlers.ListUsers)
	server.api.Get("/users/:id", handlers.GetUser)
	server.api.Get("/users/:username", handlers.FindUserByUsername)
	server.api.Post("/users", handlers.CreateNewUser)
	server.api.Put("/users/:id", handlers.UpdateUser)
	server.api.Delete("users/:id", handlers.DeleteUser)
}
