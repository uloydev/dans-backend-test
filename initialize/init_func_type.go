package initialize

import (
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type InitFunc = func(*fiber.Group, *gorm.DB, *resty.Client)
