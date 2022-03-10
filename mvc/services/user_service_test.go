package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

	"github.com/allanmaral/go-microservice/mvc/domain"
	"github.com/allanmaral/go-microservice/mvc/utils"
)

var (
	getUserFunction func(userId uint64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDao = &userDaoMock{}
}

type userDaoMock struct{}

func (m *userDaoMock) GetUser(userId uint64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestUserService_NotFoundInDatabase(t *testing.T) {
	var input uint64 = 0
	want := &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", input),
		StatusCode: http.StatusNotFound,
		Code:       "NotFound",
	}
	getUserFunction = func(userId uint64) (*domain.User, *utils.ApplicationError) {
		assert.Equal(t, userId, input)
		return nil, want
	}

	user, got := UserService.GetUser(input)

	assert.Nil(t, user)
	assert.EqualValues(t, got, want)
}

func TestUserService_ValidUser(t *testing.T) {
	var input uint64 = 123
	want := &domain.User{
		Id:        input,
		FirstName: "AnyFirstName",
		LastName:  "AnyLastName",
		Email:     "AnyEmail@AnyProvider.com",
	}
	getUserFunction = func(userId uint64) (*domain.User, *utils.ApplicationError) {
		assert.Equal(t, userId, input)
		return want, nil
	}

	got, err := UserService.GetUser(input)

	assert.Nil(t, err)
	assert.EqualValues(t, got, want)
}
