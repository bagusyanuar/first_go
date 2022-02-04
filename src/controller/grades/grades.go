package grades

import (
	"first_go/database"
	"first_go/src/lib"
	"first_go/src/model"
	"first_go/src/repository"

	"github.com/gin-gonic/gin"
)

type request struct {
	Name string `form:"name" validate:"required"`
}
func Grades(c *gin.Context) {

	if c.Request.Method == "POST" {
		var req request
		c.Bind(&req)
		_, e := lib.ValidateRequest(req)
		if e != nil {
			lib.BadRequestJsonResponse(c)
			return
		}

		grade := model.Grade{
			Name: req.Name,
			Slug: lib.MakeSlug(req.Name),
		}

		if err := database.DATABASE.Debug().Create(&grade).Error; err != nil {
			lib.ErrorJsonResponse(c, err.Error())
			return
		}
		lib.SuccessJsonResponse(c, grade)
		return
	}
	var model []model.Grade
	param := c.Query("param")
	data, err := repository.FindAllGrades(&model, param)
	if err != nil {
		lib.ErrorJsonResponse(c, err.Error())
		return
	}
	lib.SuccessJsonResponse(c, data)
}