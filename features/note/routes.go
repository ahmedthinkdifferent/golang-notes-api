package note

import (
	"github.com/gofiber/fiber/v2"
	"noteapp/core/http/middleware"
)

func RegisterNoteRoutes(app *fiber.App) {
	var v1 = app.Group("/api/v1/notes", middleware.MustBeUserMiddleware)
	v1.Get("/", GetUserNotes)
	v1.Post("/", CreateUserNote)
}
