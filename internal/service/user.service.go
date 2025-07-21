package service

import (
	"github.com/3xhvy/go-backend/internal/repo"
)

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetInfoUser() string {
	return us.userRepo.GetInfoUser()
}
