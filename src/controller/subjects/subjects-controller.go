package subjects

import (
	"first_go/src/lib"
	"first_go/src/model"
	"first_go/src/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSubjects(c *gin.Context) {
	var subjects []model.Subject

	param := c.Query("param")
	results, err := repository.FindSubjects(&subjects, param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, lib.BaseJsonResponse{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: "Internal Server Error.. " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, lib.BaseJsonResponse{
		Code:    200,
		Data:    results,
		Message: "Success",
	})
}
