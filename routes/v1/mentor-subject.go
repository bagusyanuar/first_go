package v1

import (
	"first_go/src/controller/mentor_subject"
	"first_go/src/middleware"

	"github.com/gin-gonic/gin"
)

func MentorSubjectRoutes(route *gin.RouterGroup) {
	skillGroup := route.Group("/mentor-subject")
	{
		skillGroup.POST("/", middleware.Auth, mentor_subject.CreateMentorSubject)
	}
}
