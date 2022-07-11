package usecase

import "project/test_site/internal/entity"

type (
	Usecase interface {
		AddNewIp(*entity.Ipcheck) (uint8, error)
		GetAll() []*entity.Ipcheck
		GetLogSsh() []*entity.SshLog
		GetLogCheckNet() []*entity.CheckNetLog
		GetLogCheckNetCity(string) []*entity.CheckNetLog
	}

	Repository interface {
		AddNewIp(*entity.Ipcheck) (uint8, error)
		GetAll() []*entity.Ipcheck
		GetLogSsh() []*entity.SshLog
		GetLogCheckNet() []*entity.CheckNetLog
		GetLogCheckNetCity(string) []*entity.CheckNetLog
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

// GetAll return check_ip
func (u *usecase) GetAll() []*entity.Ipcheck {
	users := u.repository.GetAll()
	return users
}

// GetLogSsh return log_ssh
func (u *usecase) GetLogSsh() []*entity.SshLog {
	users := u.repository.GetLogSsh()
	return users
}

// GetLogCheckNet return log_check_net
func (u *usecase) GetLogCheckNet() []*entity.CheckNetLog {
	users := u.repository.GetLogCheckNet()
	return users
}

// GetLogCheckNetCity return log_check_net_City
func (u *usecase) GetLogCheckNetCity(city string) []*entity.CheckNetLog {
	users := u.repository.GetLogCheckNetCity(city)
	return users
}
