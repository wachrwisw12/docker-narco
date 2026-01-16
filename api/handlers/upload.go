package handlers

import (
	"context"
	"path/filepath"
	"time"

	"api-naco/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

func UploadMedia(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return fiber.NewError(400, "file is required")
	}

	// ตรวจ size (10MB)
	if file.Size > 10*1024*1024 {
		return fiber.NewError(400, "file too large")
	}

	ext := filepath.Ext(file.Filename)
	objectName := uuid.New().String() + ext

	src, err := file.Open()
	if err != nil {
		return fiber.NewError(500, "cannot open file")
	}
	defer src.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	info, err := storage.Minio.PutObject(
		ctx,
		"reports-media",
		objectName,
		src,
		file.Size,
		minio.PutObjectOptions{
			ContentType: file.Header.Get("Content-Type"),
		},
	)
	if err != nil {
		return fiber.NewError(500, "upload failed")
	}

	return c.JSON(fiber.Map{
		"success":  true,
		"filename": objectName,
		"size":     info.Size,
	})
}
