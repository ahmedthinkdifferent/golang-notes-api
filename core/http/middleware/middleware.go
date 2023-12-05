package middleware

import (
	"github.com/gofiber/fiber/v2"
	"noteapp/core/data/db"
	"noteapp/core/data/model"
)

func MustBeUserMiddleware(c *fiber.Ctx) error { // middleware for /api/v1
	userId := c.Get("Authorization") // get the auth token from the header
	// check if the user exists in db
	var user model.User
	db.Database.Where("id = ?", userId).First(&user)
	if user.Id == 0 {
		return c.Status(401).JSON(fiber.Map{
			"status": "Unauthorized",
			"data":   nil,
		})
	}
	c.Locals("userId", userId) // set the userId in the context
	return c.Next()
}
