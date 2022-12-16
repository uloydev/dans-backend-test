package repository

import (
	"dans-backend-test/app/entity"
	"dans-backend-test/app/model"
	"dans-backend-test/exception"
	"errors"

	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
)

type JobRepository struct {
	DB         *gorm.DB
	HTTPClient *resty.Client
}

func NewJobRepository(db *gorm.DB, HTTPClient *resty.Client) BaseRepository[entity.Job] {
	return &JobRepository{
		DB:         db,
		HTTPClient: HTTPClient,
	}
}

func (repo *JobRepository) FindAll(filter *model.JobFilter) (jobs *[]model.JobResponse) {
	query := map[string]string{}

	if filter.Page != "" {
		query["page"] = filter.Page
	}

	if filter.FullTime != "" {
		query["full_time"] = filter.FullTime
	}

	if filter.Description != "" {
		query["description"] = filter.Description
	}

	if filter.Location != "" {
		query["location"] = filter.Location
	}

	resp, err := repo.HTTPClient.R().
		SetQueryParams(query).
		SetHeader("Accept", "application/json").
		SetResult(&[]model.JobResponse{}).
		Get("http://dev3.dansmultipro.co.id/api/recruitment/positions.json")

	exception.PanicWhenError(err)

	jobs = resp.Result().(*[]model.JobResponse)

	return jobs
}

func (repo *JobRepository) FindById(ID string) (job *model.JobResponse) {
	resp, err := repo.HTTPClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&model.JobResponse{}).
		Get("http://dev3.dansmultipro.co.id/api/recruitment/positions/" + ID)

	exception.PanicWhenError(err)

	job = resp.Result().(*model.JobResponse)
	if *job == (model.JobResponse{}) {
		exception.NotFoundWhenError(errors.New("not found"))
	}

	return job
}
