package controller

import (
	"dans-backend-test/app/middleware"
	"dans-backend-test/app/model"
	"dans-backend-test/app/repository"
	"dans-backend-test/app/service"
	"dans-backend-test/exception"
	"dans-backend-test/utils"

	"github.com/go-gomail/gomail"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserController struct {
	Service  service.UserService
	MailConn *gomail.Dialer
}

func NewUserController(UserService *service.UserService) BaseController {
	return &UserController{Service: *UserService}
}

func InitializeUserController(api *fiber.Group, DB *gorm.DB) {
	userRepo := repository.NewUserRepository(DB)
	userService := service.NewUserService(userRepo.(*repository.UserRepository))
	userController := NewUserController(userService.(*service.UserService))
	userController.Route(api)
}

func (controller *UserController) Route(api *fiber.Group) {
	api.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(model.WebResponse{
			Status: "Success",
			Code:   200,
		})
	})
	api.Post("/user", controller.Create)
	api.Post("/user/auth", controller.GetNewAccessToken)
	api.Get("/test-jwt", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		return c.JSON(model.WebResponse{
			Status: "Success",
			Code:   200,
		})
	})
}

// CreateUser is a function to insert user to database
// @Summary      Create User
// @Description  Create New User / Register User
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user  body      model.UserRequest  true  "Register user"
// @Success      200   {object}  model.WebResponse{data=model.UserResponse}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /v1/user [post]
func (controller *UserController) Create(c *fiber.Ctx) error {
	var request model.UserRequest

	err := c.BodyParser(&request)
	exception.PanicWhenError(err)

	request.Password = utils.GenerateHash(request.Password)

	response := controller.Service.Create(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

// AuthUser is a function to authenticate user
// @Summary      Auth User
// @Description  authenticate user / login User
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        auth  body      model.AuthRequest  true  "Auth user"
// @Success      200   {object}  model.WebResponse{data=model.AuthResponse}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /v1/user/auth [post]
func (controller *UserController) GetNewAccessToken(c *fiber.Ctx) error {
	// Generate a new Access token.
	var request model.AuthRequest
	err := c.BodyParser(&request)
	exception.PanicWhenError(err)

	auth := controller.Service.FindByAuth(request)

	utils.ValidatePassword(request.Password, auth.Password)

	jwtToken := utils.GenerateNewToken(&auth, false)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "success",
		Data:   jwtToken,
	})
}
