package util

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
	"os"
	"strings"
	"time"
)

func SaveFile(ctx *fiber.Ctx, file *multipart.FileHeader) string {
	ensureDirectoryExists("uploads")
	fileName := fmt.Sprintf("uploads/%s.%s",
		time.Now().Format("20060102150405"), strings.Split(file.Filename, ".")[1])
	err := ctx.SaveFile(file, fileName)
	if err == nil {
		return fileName
	}
	panic(err)
}

func ensureDirectoryExists(dir string) {

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0755)
	}
}
