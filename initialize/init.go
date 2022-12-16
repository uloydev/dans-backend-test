package initialize

import (
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RunInitFunctions(app *fiber.Group, DB *gorm.DB, HTTPClient *resty.Client) {
	for _, initFunction := range InitFunctions {
		initFunction(app, DB, HTTPClient)
	}
}
