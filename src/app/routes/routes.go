package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAllRoutes(api *gin.Engine) {

	router := api.Group("app")
	// setup all routes here
	router.GET("/hello", func(context *gin.Context) {
		context.String(http.StatusOK, "First Man Saying Ello Ello al internet")
	})

}
