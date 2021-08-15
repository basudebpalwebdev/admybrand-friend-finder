package api

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/basudebpalwebdev/admybrand-friend-finder/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func (server *Server) RouteSetup() {
	server.App.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	server.App.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "/swagger/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
	}))
	server.App.Get("/users", handlers.ListUsers)
	server.App.Get("/users/:id", handlers.GetUser)
	server.App.Get("/users/find/:username", handlers.FindUserByUsername)
	server.App.Post("/users", handlers.CreateNewUser)
	server.App.Put("/users/:id", handlers.UpdateUser)
	server.App.Delete("users/:id", handlers.DeleteUser)
}
