package handlers

import (
	"context"
	"time"

	"api-naco/db"
	"api-naco/models"

	"github.com/gofiber/fiber/v2"
)

type SendReportRequest struct {
	Message string `json:"content_detail"`
}

func SendReport(c *fiber.Ctx) error {
	var req SendReportRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid body")
	}

	// if len(req.Message) < 10 {
	// 	return fiber.NewError(
	// 		fiber.StatusBadRequest,
	// 		"message must be at least 10 characters",
	// 	)
	// }

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var report models.NacorticsReport

	query := `
	INSERT INTO incident_reports (details)
	VALUES ($1)
	RETURNING id,  details ,tracking_code
	`

	err := db.DB.QueryRow(ctx, query, req.Message).Scan(
		&report.ID,
		&report.ContentDetail,
		&report.TrackingCode,
	)
	if err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			err.Error(),
		)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    report,
	})
}

func Test(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
	SELECT id, content_detail
	FROM incident_reports
	ORDER BY id DESC
	LIMIT 5
	`

	rows, err := db.DB.Query(ctx, query)
	if err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			err.Error(),
		)
	}
	defer rows.Close()

	var reports []models.NacorticsReport

	for rows.Next() {
		var r models.NacorticsReport
		if err := rows.Scan(
			&r.ID,
			&r.ContentDetail,
		); err != nil {
			return fiber.NewError(
				fiber.StatusInternalServerError,
				"cannot scan report",
			)
		}
		reports = append(reports, r)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"count":   len(reports),
		"data":    reports,
	})
}
