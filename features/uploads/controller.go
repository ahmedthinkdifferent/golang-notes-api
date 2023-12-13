package uploads

import (
	"github.com/gofiber/fiber/v2"
	"noteapp/core/http/res"
	"noteapp/di"
)

func UploadFile(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return res.BadRequestError(ctx, "file is required")
	}
	var uploadService UploadService
	_ = di.Container.Invoke(func(service UploadService) {
		uploadService = service
	})
	fileName := uploadService.UploadFile(ctx, file)
	return res.Success(ctx, fiber.Map{
		"path": fileName,
	})
}
