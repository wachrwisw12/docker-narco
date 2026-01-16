package routers

import (
	handlers "api-naco/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupNarcoticsReport(nacoticRoute fiber.Router) {
	nacoticRoute.Post("/sendreport", handlers.SendReport)
	nacoticRoute.Get("/test", handlers.Test)
}

func SetupAuth(auth fiber.Router) {
	auth.Post("/singin", handlers.Authhandler)
	auth.Post("/register", handlers.Registerhandler)
}
