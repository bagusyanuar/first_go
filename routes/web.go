package routes

import (
	v1 "first_go/routes/v1"
	"first_go/src/controller"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine {

	route := gin.Default()
	apiV1 := route.Group("/api/v1")
	{
		v1.UserRoutes(apiV1)
		v1.V1SubjectRoutes(apiV1)
	}
	route.POST("/test", controller.SuperAdminSeeder)
	return route
}
