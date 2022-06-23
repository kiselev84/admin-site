package usecase

import "project/test_site/internal/entity"

type (
	Usecase interface {
		AddNewIp(*entity.Ipcheck) (uint8, error)
		GetAll() []*entity.Ipcheck
	}

	Repository interface {
		AddNewIp(*entity.Ipcheck) (uint8, error)
		GetAll() []*entity.Ipcheck
	}
)

type usecase struct {
	repository Repository
}

func NewUsecase(repository Repository) *usecase {
	return &usecase{
		repository: repository,
	}
}

// CreateUser return user id or error
func (u *usecase) AddNewIp(user *entity.Ipcheck) (uint8, error) {
	uid, error := u.repository.AddNewIp(user)
	return uid, error
}

// GetAll return users
func (u *usecase) GetAll() []*entity.Ipcheck {
	users := u.repository.GetAll()
	return users
}
