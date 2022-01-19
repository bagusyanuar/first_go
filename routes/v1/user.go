package v1

import (
	"first_go/src/controller"
	"first_go/src/controller/users"

	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.RouterGroup) {
	userGroup := route.Group("/users")
	{
		userGroup.GET("/", users.GetUsers)
		userGroup.POST("/", users.GetUsers)
	}

	authGroup := route.Group("/auth")
	{
		member := authGroup.Group("/member")
		{
			member.POST("/sign-up", controller.MemberSignUp)
			member.POST("/sign-in", controller.MemberSignUp)
		}
	}
}
