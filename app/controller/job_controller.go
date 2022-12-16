package controller

import (
	"dans-backend-test/app/middleware"
	"dans-backend-test/app/model"
	"dans-backend-test/app/repository"
	"dans-backend-test/app/service"
	"dans-backend-test/exception"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type JobController struct {
	Service service.JobService
}

func NewJobController(JobService *service.JobService) BaseController {
	return &JobController{Service: *JobService}
}

func InitializeJobController(api *fiber.Group, DB *gorm.DB, HTTPClient *resty.Client) {
	jobRepo := repository.NewJobRepository(DB, HTTPClient)
	jobService := service.NewJobService(jobRepo.(*repository.JobRepository))
	jobController := NewJobController(jobService.(*service.JobService))
	jobController.Route(api)
}

func (controller *JobController) Route(api *fiber.Group) {
	api.Get("/job", middleware.JWTProtected(), controller.List)
	api.Get("/job/:id", middleware.JWTProtected(), controller.FindById)
}

// GetJobList is a function to get all job data
// @Security ApiKeyAuth
// @Summary      Get All Job
// @Description  get all Job Data
// @Tags         job
// @Accept       json
// @Produce      json
// @Success      200   {object}  model.WebResponse{data=[]model.JobResponse}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /v1/job [get]
func (c *JobController) List(ctx *fiber.Ctx) error {
	params := &model.JobFilter{}
	err := ctx.QueryParser(params)

	exception.PanicWhenError(err)

	responses := c.Service.List(params)

	return ctx.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}

// GeJob is a function to get Job data by id
// @Security ApiKeyAuth
// @Summary      Get Job by id
// @Description  get Job data by id
// @Tags         job
// @Accept       json
// @Produce      json
// @Param id path string true "id"
// @Success      200   {object}  model.WebResponse{data=model.JobResponse}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Failure      404   {object}  model.WebResponse{data=string}
// @Router       /v1/job/ [get]
func (c *JobController) FindById(ctx *fiber.Ctx) error {

	ID := ctx.Params("id")
	resp := c.Service.FindById(ID)

	return ctx.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   resp,
	})
}
