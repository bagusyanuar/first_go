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
		admin := authGroup.Group("/admin")
		{
			admin.POST("/sign-up", controller.AdminSignUp)
		}
		member := authGroup.Group("/member")
		{
			member.POST("/sign-up", controller.MemberSignUp)
			member.POST("/sign-in", controller.MemberSignIn)
		}

		mentor := authGroup.Group("/mentor")
		{
			mentor.POST("/sign-up", controller.MentorSignUp)
			mentor.POST("/sign-in", controller.MentorSignIn)
		}
	}
}
