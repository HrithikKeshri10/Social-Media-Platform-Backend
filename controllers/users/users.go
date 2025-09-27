package users

import (
	"social-media-app/internals/dto"
	"social-media-app/internals/notifications"
	"social-media-app/internals/validator"
	"social-media-app/services/users"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Add(c *fiber.Ctx) error {

	ctx := c.UserContext()

	var user dto.UserCreate

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("incorrect input body")
	}

	if err := validator.Payload(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("incorrect input body")
	}

	us := users.New()
	us.User = &dto.User{}

	us.User.Name = user.Name
	us.User.Email = user.Email
	us.User.Password = user.Password

	us.Create(ctx)

	notifications.Register(us.User.ID)
	go notifications.ListenForNotifications(ctx, us.User.ID)

	return c.Status(fiber.StatusCreated).JSON(us.User)
}

func Get(c *fiber.Ctx) error {
	ctx := c.UserContext()

	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("incorrect user id")
	}

	us := users.New()
	us.User = &dto.User{}

	us.User.ID = userID

	if err := us.Get(ctx); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("user not found")
		}
	}

	return c.Status(fiber.StatusOK).JSON(us.User)
}

func Delete(c *fiber.Ctx) error {
	ctx := c.UserContext()

	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("incorrect user id")
	}

	us := users.New()
	us.User = &dto.User{}

	us.User.ID = userID

	if err := us.Delete(ctx); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("user not found!")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("internal server error")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func GetAll(c *fiber.Ctx) error {
	ctx := c.UserContext()

	us := users.New()
	us.Users = &dto.Users{}

	if err := us.GetAll(ctx); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("user not found!")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("internal server error")
	}

	return c.Status(fiber.StatusOK).JSON(us.Users)
}
