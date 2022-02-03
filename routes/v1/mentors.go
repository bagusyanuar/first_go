package v1

import (
	"first_go/src/controller/users"
	"first_go/src/middleware"

	"github.com/gin-gonic/gin"
)

func MentorRoutes(route *gin.RouterGroup) {
	mentorGroup := route.Group("/mentor")
	{
		mentorGroup.GET("/", middleware.Auth, users.GetMentorProfile)

		meGroup := mentorGroup.Group("/me")
		{
			meGroup.GET("/", middleware.Auth, users.GetMentorProfile)
			meGroup.GET("/subjects", middleware.Auth, users.MentorSubjects)
			meGroup.POST("/subjects", middleware.Auth, users.MentorSubjects)
		}
	}
}
