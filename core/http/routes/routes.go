package routes

import (
	"github.com/gofiber/fiber/v2"
	"noteapp/features/auth"
	"noteapp/features/note"
)

func RegisterAppRoutes(app *fiber.App) {
	auth.RegisterAuthRoutes(app)
	note.RegisterNoteRoutes(app)
}
