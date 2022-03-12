package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/allanmaral/go-microservice/mvc/utils"
)

func Respond(c *gin.Context, status int, body interface{}) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(status, body)
		return
	}
	c.JSON(status, body)
}

func RespondWithError(c *gin.Context, err *utils.ApplicationError) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(err.StatusCode, err)
		return
	}
	c.JSON(err.StatusCode, err)
}
