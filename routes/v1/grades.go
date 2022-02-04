package v1

import (
	"first_go/src/controller/grades"

	"github.com/gin-gonic/gin"
)


func V1GradeRoutes(route *gin.RouterGroup) {
	gradeGroup := route.Group("/grades")
	{
		gradeGroup.GET("/", grades.Grades)
		gradeGroup.POST("/", grades.Grades)
	}

}