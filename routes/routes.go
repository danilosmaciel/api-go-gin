package routes

import (
	"github.com/danilosmaciel/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() *gin.Engine {
	r := gin.Default()
	r.POST("/api/v1/login", controllers.Authenticate)
	r.GET("/api/v1/state/import", controllers.StateImport)
	r.GET("/api/v1/state", controllers.StateHandler)
	r.GET("/api/v1/city/import", controllers.CityImport)
	return r
}
