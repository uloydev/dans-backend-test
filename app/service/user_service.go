package service

import (
	"dans-backend-test/app/entity"
	"dans-backend-test/app/model"
	"dans-backend-test/app/repository"
	"dans-backend-test/app/validation"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) BaseService[model.UserRequest, model.UserResponse] {
	return &UserService{
		Repo: repo,
	}
}

func (service *UserService) Create(request model.UserRequest) (response model.UserResponse) {
	validation.ValidateUser(request)

	user := entity.User{
		Username: request.Username,
		Password: request.Password,
	}

	user = service.Repo.Insert(user)

	response = model.UserResponse{
		BasicData: model.BasicData{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		},
		Username: user.Username,
	}

	return response
}

func (service *UserService) List() (responses []model.UserResponse) {
	users := service.Repo.FindAll()

	for _, user := range users {
		responses = append(responses, model.UserResponse{
			BasicData: model.BasicData{
				ID:        user.ID,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
				DeletedAt: user.DeletedAt,
			},
			Username: user.Username,
		})
	}

	return responses
}

func (service *UserService) FindById(user_id uint) (response model.UserResponse) {
	user := service.Repo.FindById(user_id)

	response = model.UserResponse{
		BasicData: model.BasicData{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		},
		Username: user.Username,
	}

	return response
}

func (service *UserService) FindByUsername(email string) (response model.UserResponse) {
	user := service.Repo.FindByUsername(email)

	response = model.UserResponse{
		BasicData: model.BasicData{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		},
		Username: user.Username,
	}
	return response
}

func (service *UserService) FindByAuth(request model.AuthRequest) (response model.AuthResponse) {
	validation.ValidateAuth(request)
	user := service.Repo.FindByUsername(request.Username)

	response = model.AuthResponse{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
	}

	return response
}
