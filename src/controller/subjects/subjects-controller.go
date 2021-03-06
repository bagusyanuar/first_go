package subjects

import (
	"first_go/database"
	"first_go/src/lib"
	"first_go/src/model"
	"first_go/src/repository"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type request_body struct {
	Name string                `form:"name" validate:"required"`
	Icon *multipart.FileHeader `form:"icon"`
}

func Subjects(c *gin.Context) {

	if c.Request.Method == "POST" {
		var request request_body
		c.Bind(&request)
		
		
		validate := validator.New()
		err_validation := validate.Struct(request)
		if err_validation != nil {
			errs := err_validation.(validator.ValidationErrors)
			var ar_err []string
			for _, e := range errs {
				ar_err = append(ar_err, e.Tag())
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, lib.BaseJsonResponse{
				Code: http.StatusBadRequest,
				Data: ar_err,
				Message: "Bad Request Validation Form Data",
			})
			return
		}
		var icon_name *string
		if request.Icon != nil {
			ext := filepath.Ext(request.Icon.Filename)
			filename := "assets/icons/" + uuid.New().String() + ext
			icon_name = &filename
			if err_upload := c.SaveUploadedFile(request.Icon, filename); err_upload != nil {
				c.JSON(http.StatusInternalServerError, lib.BaseJsonResponse{
					Code:    http.StatusInternalServerError,
					Data:    nil,
					Message: "Internal Server Error. Failed while upload icon",
				})
				return
			}
		}

		subject := model.Subject{
			Name: request.Name,
			Slug: lib.MakeSlug(request.Name),
			Icon: icon_name,
		}

		if err_insert := database.DATABASE.Debug().Create(&subject).Error; err_insert != nil {
			c.JSON(http.StatusInternalServerError, lib.BaseJsonResponse{
				Code:    http.StatusInternalServerError,
				Data:    nil,
				Message: "Internal Server Error. Failed while Create New subject",
			})
			return
		}
		c.JSON(http.StatusOK, lib.BaseJsonResponse{
			Code: http.StatusOK,
			Data: subject,
			Message: "success create new subject",
		})
		return
	}
	var subjects []model.Subject

	param := c.Query("param")
	data, err := repository.FindSubjects(&subjects, param)
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
		Data:    data,
		Message: "Success",
	})
}

func GetSubjectBySlug(c *gin.Context) {
	slug := c.Param("slug")

	var subject model.Subject
	data, err := repository.FindSubjectBySlug(&subject, slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, lib.BaseJsonResponse{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: "Error While Find Subject By Slug",
		})
		return
	}
	c.JSON(http.StatusOK, lib.BaseJsonResponse{
		Code:    http.StatusOK,
		Data:    data,
		Message: "success",
	})
}
