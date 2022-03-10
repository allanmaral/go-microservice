package domain

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

	"github.com/allanmaral/go-microservice/mvc/utils"
)

func TestGetUserNoUserFound(t *testing.T) {
	var input uint64 = 0
	want := &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", input),
		StatusCode: http.StatusNotFound,
		Code:       "NotFound",
	}

	usr, got := UserDao.GetUser(input)

	assert.Nil(t, usr)
	assert.EqualValues(t, got, want)
}

func TestGetUserValidId(t *testing.T) {
	var input uint64 = 123
	want := &User{Id: 123, FirstName: "Allan", LastName: "Ribeiro", Email: "email@email.com"}

	got, err := UserDao.GetUser(input)

	assert.Nil(t, err)
	assert.EqualValues(t, got, want)
}
