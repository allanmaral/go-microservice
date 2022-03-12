package app

import (
	"github.com/allanmaral/go-microservice/mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
