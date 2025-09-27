package config

import (
	"social-media-app/internals/database"
	"social-media-app/models/friendship"
	"social-media-app/models/posts"
	"social-media-app/models/users"
)

func Automigration() {
	database.Client().AutoMigrate(&users.Users{}, &friendship.Friendships{}, &posts.Posts{})
}
