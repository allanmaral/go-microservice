package services

import (
	"github.com/allanmaral/go-microservice/mvc/domain"
	"github.com/allanmaral/go-microservice/mvc/utils"
)

type userService struct {
}

var (
	UserService userService
)

func (u *userService) GetUser(userId uint64) (*domain.User, *utils.ApplicationError) {
	user, err := domain.UserDao.GetUser(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}
