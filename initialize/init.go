package initialize

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RunInitFunctions(app *fiber.Group, DB *gorm.DB) {
	for _, initFunction := range InitFunctions {
		initFunction(app, DB)
	}
}
