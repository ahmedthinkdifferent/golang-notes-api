package routes

import (
	"github.com/gofiber/fiber/v2"
	"noteapp/features/auth"
	"noteapp/features/note"
	"noteapp/features/uploads"
)

func RegisterAppRoutes(app *fiber.App) {
	auth.RegisterAuthRoutes(app)
	note.RegisterNoteRoutes(app)
	uploads.RegisterUploadsRoutes(app)
}
