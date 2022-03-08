package services

import (
	"github.com/allanmaral/go-microservice/mvc/domain"
	"github.com/allanmaral/go-microservice/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
