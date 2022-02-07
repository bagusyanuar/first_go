package users

import (
	"errors"
	"first_go/database"
	"first_go/src/lib"
	"first_go/src/model"
	"first_go/src/response"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetMentorProfile(c *gin.Context) {
	user := c.MustGet("user").(jwt.MapClaims)
	var mentor response.MentorProfileResponse
	err := database.DATABASE.Debug().
		Preload("User", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id", "email", "username")
		}).
		Preload("Skills.Subject").
		Joins("JOIN users ON users.id = mentors.user_id").
		Where("mentors.id = ?", user["identifier"]).
		First(&mentor).Error
	if err != nil {
		eMsg := "Internal Server Error " + err.Error()
		sMsg := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			eMsg = "User Not Found!"
			sMsg = http.StatusUnauthorized
		}
		c.AbortWithStatusJSON(sMsg, lib.BaseJsonResponse{
			Code:    sMsg,
			Data:    nil,
			Message: eMsg,
		})
		return
	}
	c.JSON(http.StatusOK, lib.BaseJsonResponse{
		Code:    http.StatusOK,
		Data:    mentor,
		Message: "success",
	})
}

type request_body struct {
	SubjectID uint `form:"subject_id" validate:"required"`
}
func MentorSubjects(c *gin.Context)  {
	user := c.MustGet("user").(jwt.MapClaims)

	if c.Request.Method == "POST" {
		var request request_body
		c.Bind(&request)
		if request.SubjectID == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, lib.BaseJsonResponse{
				Code: http.StatusBadRequest,
				Data: nil,
				Message: "parameter subject id cannot be empty",
			})
			return
		}

		m_id, e := uuid.Parse(user["identifier"].(string))
		if e != nil {
			c.JSON(http.StatusInternalServerError, lib.BaseJsonResponse{
				Code:    http.StatusInternalServerError,
				Data:    nil,
				Message: "internal server error. failed to convert string id",
			})
			return
		}

		isExist := true
		if err_isExist := database.DATABASE.Debug().
		Where("mentor_id = ?", m_id).
		Where("subject_id = ?", request.SubjectID).
		First(&model.MentorSubject{}).Error; err_isExist != nil {
			if errors.Is(err_isExist, gorm.ErrRecordNotFound) {
				isExist = false
			}
		} 

		if isExist {
			c.JSON(http.StatusAccepted, lib.BaseJsonResponse{
				Code:    http.StatusAccepted,
				Data:    nil,
				Message: "subject is exist",
			})
			return
		}
		mentor_subject := model.MentorSubject{
			MentorID: m_id,
			SubjectID: uint(request.SubjectID),
		}
		if err_create := database.DATABASE.Debug().Create(&mentor_subject).Error; err_create != nil {
			c.JSON(http.StatusInternalServerError, lib.BaseJsonResponse{
				Code:    http.StatusInternalServerError,
				Data:    nil,
				Message: "internal server error. failed to create a new subject skill",
			})
			return
		}
		c.JSON(http.StatusOK, lib.BaseJsonResponse{
			Code: http.StatusOK,
			Data: mentor_subject,
			Message: "success create new subject skill",
		})
		return
	}
	
	var data []response.PreloadMentorSubjectAll
	err := database.DATABASE.Debug().
		Preload("Subject").
		Where("mentor_id = ?", user["identifier"]).
		Find(&data).Error
	if err != nil {
		eMsg := "Internal Server Error " + err.Error()
		sMsg := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			eMsg = "User Not Found!"
			sMsg = http.StatusUnauthorized
		}
		c.AbortWithStatusJSON(sMsg, lib.BaseJsonResponse{
			Code:    sMsg,
			Data:    nil,
			Message: eMsg,
		})
		return
	}
	c.JSON(http.StatusOK, lib.BaseJsonResponse{
		Code:    http.StatusOK,
		Data:    data,
		Message: "success load mentor data",
	})
}
