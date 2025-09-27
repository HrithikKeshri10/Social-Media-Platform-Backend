package routes

import (
	"social-media-app/controllers/users"

	"github.com/gofiber/fiber/v2"
)

func Users(r fiber.Router) {
	u := r.Group("/users")

	u.Post("/", users.Add)
	u.Get("/", users.GetAll)

	u.Get("/:id", users.Get)
	u.Delete("/:id", users.Delete)
}
