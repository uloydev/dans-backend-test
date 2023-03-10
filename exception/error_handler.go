package exception

import (
	"dans-backend-test/app/model"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	if _, ok := err.(ValidationError); ok {
		return ctx.JSON(model.WebResponse{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	if _, ok := err.(NotFoundError); ok {
		return ctx.JSON(model.WebResponse{
			Code:   404,
			Status: "NOT_FOUND",
			Data:   err.Error(),
		})
	}

	return ctx.JSON(model.WebResponse{
		Code:   500,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err.Error(),
	})
}

func JwtError(ctx *fiber.Ctx, err error) error {

	if err.Error() == "Missing or malformed JWT" {
		return ctx.JSON(model.WebResponse{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}
	return ctx.JSON(model.WebResponse{
		Code:   401,
		Status: "UNAUTHORIZED",
		Data:   err.Error(),
	})
}
