package note

import (
	"github.com/gofiber/fiber/v2"
	"noteapp/core/http/res"
	"noteapp/core/util"
	"strconv"
)

type CreateUserNoteDto struct {
	Title   string   `json:"title" validate:"required,min=2"`
	Content string   `json:"content" validate:"required,min=2"`
	Files   []string `json:"files"`
}

func GetUserNotes(ctx *fiber.Ctx) error {
	userId, _ := strconv.Atoi(ctx.Get("userId"))
	notes := UserNotes(userId, &Repo{})
	return res.Success(ctx, notes)
}

func CreateUserNote(ctx *fiber.Ctx) error {
	userId, _ := ctx.Locals("userId").(int)
	var dto = new(CreateUserNoteDto)
	err := ctx.BodyParser(&dto)
	if err != nil {
		return res.ValidationError(ctx, map[string]any{"error": "Invalid request body"})
	}
	if validationError := util.Validate(dto); validationError.HasErrors() {
		return res.ValidationError(ctx, validationError.Errors)
	}
	userNote := SaveUserNote(userId, dto.Title, dto.Content, dto.Files, &Repo{})
	return res.Success(ctx, userNote)
}
