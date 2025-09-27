package routes

import (
	"social-media-app/controllers/friendships"

	"github.com/gofiber/fiber/v2"
)

func Friendships(r fiber.Router) {
	friendRoutes := r.Group("/friends")

	friendRoutes.Post("/", friendships.Add)
	friendRoutes.Get("/:id", friendships.Get)
	friendRoutes.Delete("/:id", friendships.Delete)
}
