package server

import (
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
)

func middlewares(app *fiber.App) {
	app.Use(logger.New())
}
