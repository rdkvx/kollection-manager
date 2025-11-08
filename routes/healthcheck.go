package routes

import (
	"github.com/gofiber/fiber/v2"
	"kollectionmanager/m/utils"
)

func Healthcheck(app *fiber.App) {
	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Get("/readyz", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": utils.ServerStatus,
		})
	})
}