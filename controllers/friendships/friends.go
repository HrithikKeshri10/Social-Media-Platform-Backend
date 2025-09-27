package friendships

import (
	"fmt"
	"social-media-app/internals/cache"
	"social-media-app/internals/dto"
	"social-media-app/internals/validator"
	"social-media-app/models/friendship"
	"social-media-app/services/friendships"
	"social-media-app/services/users"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Add(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var friend dto.FriendCreate

	if err := c.BodyParser(&friend); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("incorrect input body")
	}

	if err := validator.Payload(friend); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("incorrect input body")
	}

	us := users.New()
	us.User = &dto.User{}
	us.User.ID = friend.UserID
	if err := us.Get(ctx); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("user not found")
		}
	}

	us.User = &dto.User{}
	us.User.ID = friend.FriendID
	if err := us.Get(ctx); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("user not found")
		}
	}

	fs := friendship.New()
	fs.Friends = &dto.Friends{}
	fs.Friends.UserID = friend.UserID
	fs.Friends.FriendID = friend.FriendID

	fs.Create(ctx)

	if err := cache.Client().Del(ctx, friend.UserID.String()).Err(); err != nil {
		fmt.Printf("Error invalidating cache: %v", err)
	}

	return c.Status(fiber.StatusCreated).JSON(fs.Friends)
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

	fs := friendships.New()
	fs.UserID = userID
	fs.GetAll(ctx)

	return c.Status(fiber.StatusOK).JSON(fs.AllFriends)
}

func Delete(c *fiber.Ctx) error {
	ctx := c.UserContext()

	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("incorrect user id")
	}

	fid := c.Query("f_id")
	friendID, err := uuid.Parse(fid)
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

	us.User = &dto.User{}
	us.User.ID = friendID
	if err := us.Get(ctx); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("user not found!")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("internal server error")
	}

	fs := friendships.New()
	fs.UserID = userID
	fs.FriendID = friendID
	fs.Delete(ctx)

	if err := cache.Client().Del(ctx, userID.String()).Err(); err != nil {
		fmt.Printf("Error invalidating cache: %v", err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
