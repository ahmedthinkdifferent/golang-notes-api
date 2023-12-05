package auth

import "github.com/gofiber/fiber/v2"

func RegisterAuthRoutes(app *fiber.App) {
	var v1 = app.Group("/api/v1/auth")
	v1.Post("/registration", Register)
	v1.Get("/:userId/notes", GetUserWithNotes)
	v1.Post("/login", Login)
}
