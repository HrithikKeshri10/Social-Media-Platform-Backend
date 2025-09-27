package server

import (
	"social-media-app/routes"

	"github.com/gofiber/fiber/v2"
)

// Default Error Handler
func errHandler(ctx *fiber.Ctx, err error) error {
	msg := err.Error()
	return ctx.Status(fiber.StatusInternalServerError).JSON(msg)
}

// Not found handler
var notFoundHandler = func(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).JSON("Requested resource not found")
}

func addRoutes(app *fiber.App) {
	baseRouter := app.Group("/socio")
	routes.Users(baseRouter)
	routes.Friendships(baseRouter)
	routes.Posts(baseRouter)
}
