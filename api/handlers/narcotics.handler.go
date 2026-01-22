package handlers

import (
	"context"
	"time"

	"api-naco/db"
	"api-naco/models"

	"github.com/gofiber/fiber/v2"
)

func ReportInit(c *fiber.Ctx) error {
	println("test innit")
	return nil
}

type SendReportRequest struct {
	Details string `json:"details"`
}

func SendReport(c *fiber.Ctx) error {
	req := SendReportRequest{
		Details: c.FormValue("details"),
	}
	if req.Details == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing required fields")
	}
	// filse, err := c.FormFile("file")
	// if err != nil {
	// 	return fiber.NewError(fiber.StatusBadRequest, "file required")
	// }
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

	err := db.DB.QueryRow(ctx, query, req.Details).Scan(
		&report.ID,
		&report.Details,
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
		//"files":   filse,
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
			//&r.Details,
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
