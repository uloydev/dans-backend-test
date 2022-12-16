package service

import (
	"dans-backend-test/app/model"
	"dans-backend-test/app/repository"
)

type JobService struct {
	Repo *repository.JobRepository
}

func NewJobService(repo *repository.JobRepository) BaseService[model.JobResponse, model.JobResponse] {
	return &JobService{
		Repo: repo,
	}
}

func (service *JobService) List() (responses *[]model.JobResponse) {
	responses = service.Repo.FindAll()
	return responses
}

func (service *JobService) FindById(ID string) (response *model.JobResponse) {
	response = service.Repo.FindById(ID)

	return response
}
