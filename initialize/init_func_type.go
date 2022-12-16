package initialize

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type InitFunc = func(*fiber.Group, *gorm.DB)
