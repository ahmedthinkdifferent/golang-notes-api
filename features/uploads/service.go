package uploads

import (
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
	"noteapp/core/util"
)

type UploadService interface {
	UploadFile(ctx *fiber.Ctx, header *multipart.FileHeader) string
}

type uploadService struct {
}

func (receiver *uploadService) UploadFile(ctx *fiber.Ctx, file *multipart.FileHeader) string {
	fileName := util.SaveFile(ctx, file)
	return fileName
}

func NewUploadService() UploadService {
	return &uploadService{}
}
