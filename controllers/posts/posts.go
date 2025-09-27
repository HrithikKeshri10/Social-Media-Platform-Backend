package posts

import (
	"fmt"
	"social-media-app/internals/dto"
	"social-media-app/internals/notifications"
	"social-media-app/internals/validator"
	"social-media-app/services/posts"
	"social-media-app/services/users"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Add(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var post dto.PostCreate

	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("incorrect user id")
	}

	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("incorrect input body")
	}

	if err := validator.Payload(post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("incorrect input body")
	}

	us := users.New()
	us.User = &dto.User{}
	us.User.ID = userID
	if err := us.Get(ctx); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("user not found!")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("internal server error")
	}

	ps := posts.New()
	ps.Post = &dto.Post{}

	ps.Post.UserID = userID
	ps.Post.Content = post.Content

	ps.Create(ctx)

	msg := fmt.Sprintf("Hello, your friend %v has created a new post", us.User.Name)
	notifications.NotifyUsers(ctx, userID, msg)
	
	return c.Status(fiber.StatusCreated).JSON(ps.Post)
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
			return c.Status(fiber.StatusNotFound).JSON("user not found!")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("internal server error")
	}

	ps := posts.New()
	ps.Posts = &dto.Posts{}
	ps.UserID = userID

	ps.GetAll(ctx)

	return c.Status(fiber.StatusOK).JSON(ps.Posts)
}

func Delete(c *fiber.Ctx) error {
	ctx := c.UserContext()

	userID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("incorrect user id")
	}

	postID, err := uuid.Parse(c.Params("post_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("incorrect user id")
	}

	us := users.New()
	us.User = &dto.User{}
	us.User.ID = userID
	if err := us.Get(ctx); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("user not found!")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("internal server error")
	}

	ps := posts.New()
	ps.Posts = &dto.Posts{}

	ps.UserID = userID
	ps.ID = postID

	ps.Delete(ctx)

	return c.SendStatus(fiber.StatusNoContent)
}
