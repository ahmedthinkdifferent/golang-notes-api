package middleware

import (
	"github.com/gofiber/fiber/v2"
	"noteapp/core/data/db"
	"noteapp/core/data/model"
	"noteapp/core/util"
)

func MustBeUserMiddleware(c *fiber.Ctx) error { // middleware for /api/v1
	headerToken := c.Get("Authorization") // get the auth token from the header
	token, err := util.VerifyJwtToken(headerToken)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"status": "Unauthorized",
			"data":   nil,
		})
	}
	// check if the user exists in db
	var user model.User
	db.Database.Where("id = ?", token["user_id"]).First(&user)
	if user.Id == 0 {
		return c.Status(401).JSON(fiber.Map{
			"status": "Unauthorized",
			"data":   nil,
		})
	}
	c.Locals("userId", user.Id) // set the userId in the context
	return c.Next()
}
