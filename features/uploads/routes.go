package uploads

import "github.com/gofiber/fiber/v2"

func RegisterUploadsRoutes(app *fiber.App) {
	app.Post("/api/v1/uploads", UploadFile)
}
