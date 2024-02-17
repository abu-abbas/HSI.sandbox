package api

import (
	"github.com/abu-abbas/level_5/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func (server *AppServer) routes() {
	server.app.Use(requestid.New())
	server.app.Use(logger.New(
		logger.Config{
			Format: "[${ip}]:${port} ${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
		},
	))

	server.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	middleware := basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "x",
		},
	})

	api := server.app.Group("/api", middleware)
	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})

	item := controllers.Item{}

	v1.Get("/", item.Index)
	v1.Post("/", item.Create)
	v1.Put("/:id", item.Edit)
	v1.Delete("/:id", item.Delete)
}
