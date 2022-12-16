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

func (repo *JobRepository) FindAll() (jobs *[]model.JobResponse) {
	resp, err := repo.HTTPClient.R().
		//   SetQueryParams(map[string]string{
		//       "page_no": "1",
		//       "limit": "20",
		//       "sort":"name",
		//   }).
		SetHeader("Accept", "application/json").
		SetResult(&[]model.JobResponse{}).
		Get("http://dev3.dansmultipro.co.id/api/recruitment/positions.json")

	exception.PanicWhenError(err)

	jobs = resp.Result().(*[]model.JobResponse)

	return jobs
}

func (repo *JobRepository) FindById(ID string) (job *model.JobResponse) {
	resp, err := repo.HTTPClient.R().
		//   SetQueryParams(map[string]string{
		//       "page_no": "1",
		//       "limit": "20",
		//       "sort":"name",
		//   }).
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
