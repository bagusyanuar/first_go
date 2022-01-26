package v1

import (
	"first_go/src/controller/subjects"

	"github.com/gin-gonic/gin"
)

func V1SubjectRoutes(route *gin.RouterGroup) {
	subjectGroup := route.Group("/subjects")
	{
		subjectGroup.GET("/", subjects.GetSubjects)
	}
}
