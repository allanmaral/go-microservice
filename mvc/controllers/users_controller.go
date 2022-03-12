package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"github.com/allanmaral/go-microservice/mvc/services"
	"github.com/allanmaral/go-microservice/mvc/utils"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "BadRequest",
		}

		RespondWithError(c, apiErr)
		return
	}

	user, apiErr := services.UserService.GetUser(userId)
	if apiErr != nil {
		RespondWithError(c, apiErr)
		return
	}

	// return thr user to client
	Respond(c, http.StatusOK, user)
}
