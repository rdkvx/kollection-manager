package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Router(app *fiber.App, db *gorm.DB) error {

	ConsoleRoutes(app, db)
	DeveloperRoutes(app, db)
	ManufacturerRoutes(app, db)
	GameRoutes(app, db)
	Healthcheck(app)
	return nil
}
