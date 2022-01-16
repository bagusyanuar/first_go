package v1

import (
	"first_go/src/controller/users"

	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.RouterGroup) {
	userGroup := route.Group("/users")
	{
		userGroup.GET("/", users.GetUsers)
		userGroup.POST("/", users.GetUsers)
	}
}
