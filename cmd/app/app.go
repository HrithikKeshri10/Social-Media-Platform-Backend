package app

import (
	"log"
	"social-media-app/internals/cache"
	"social-media-app/internals/database"
	"social-media-app/internals/notifications"
	"social-media-app/internals/server"
)

func Setup() {
	database.Connect()
	cache.Connect()

	notifications.InitNotificationsSystem()
	notifications.Hydrate()

	server.Setup()
	app := server.New()

	if err := app.Listen(":3015"); err != nil {
		log.Fatalf("Error starting server %v\n", err)
	}
}
