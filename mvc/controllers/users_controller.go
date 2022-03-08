package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/allanmaral/go-microservice/mvc/services"
	"github.com/allanmaral/go-microservice/mvc/utils"
)

func GetUser(response http.ResponseWriter, request *http.Request) {
	userId, err := strconv.ParseInt(request.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "BadRequest",
		}
		jsonValue, _ := json.Marshal(apiErr)
		response.WriteHeader(apiErr.StatusCode)
		_, _ = response.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		response.WriteHeader(apiErr.StatusCode)
		_, _ = response.Write(jsonValue)
		return
	}

	// return thr user to client
	jsonValue, _ := json.Marshal(user)
	_, _ = response.Write(jsonValue)
}
