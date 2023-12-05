package note

import (
	"github.com/gofiber/fiber/v2"
	"noteapp/core/http/res"
	"noteapp/core/util"
	"strconv"
)

type CreateUserNoteDto struct {
	Title   string `json:"title" validate:"required,min=2"`
	Content string `json:"content" validate:"required,min=2"`
}

func GetUserNotes(ctx *fiber.Ctx) error {
	userId, _ := strconv.Atoi(ctx.Get("userId"))
	notes := UserNotes(userId, &NoteRepo{})
	return res.Success(ctx, notes)
}

func CreateUserNote(ctx *fiber.Ctx) error {
	userId, _ := strconv.Atoi(ctx.Locals("userId").(string))
	var dto = new(CreateUserNoteDto)
	err := ctx.BodyParser(&dto)
	if err != nil {
		return res.ValidationError(ctx, map[string]any{"error": "Invalid request body"})
	}
	if validationError := util.Validate(dto); validationError.HasErrors() {
		return res.ValidationError(ctx, validationError.Errors)
	}
	userNote := SaveUserNote(userId, dto.Title, dto.Content, &NoteRepo{})
	return res.Success(ctx, userNote)
}
