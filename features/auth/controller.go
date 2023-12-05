package auth

import (
	"github.com/gofiber/fiber/v2"
	"noteapp/core/http/res"
	"noteapp/features/note"
	"strconv"
)

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(ctx *fiber.Ctx) error {
	return RegisterUser(ctx, NewUserRepository())
}

func GetUserWithNotes(ctx *fiber.Ctx) error {
	userId, _ := strconv.Atoi(ctx.Params("userId"))
	user := (&note.NoteRepo{}).UserWithNotes(userId)
	return res.Success(ctx, user)
}

func Login(ctx *fiber.Ctx) error {
	var loginDto LoginDto
	ctx.BodyParser(&loginDto)
	user, err := LoginUser(loginDto, NewUserRepository())
	if err == nil {
		return res.Success(ctx, user)
	}
	return res.Error(ctx, err.Error, err.StatusCode)
}
