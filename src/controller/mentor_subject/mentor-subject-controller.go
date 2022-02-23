package mentor_subject

import (
	"errors"
	"first_go/database"
	"first_go/src/lib"
	"first_go/src/model"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type MentorSubjectRequest struct {
	SubjectID uint           `json:"subject_id"`
	GradeID   uint           `json:"grade_id"`
	Method    datatypes.JSON `json:"method"`
}

type response struct {
	model.MentorSubject
	Mentor struct {
		ID      uuid.UUID  `json:"id"`
		UserID  uuid.UUID  `json:"user_id"`
		Name    string     `json:"name"`
		Phone   string     `json:"phone"`
		Avatar  lib.Avatar `json:"avatar"`
		Address string     `json:"address"`
	} `gorm:"foreignKey:MentorID" json:"mentor"`
	Subject struct {
		ID   uint    `json:"id"`
		Name string  `json:"name"`
		Slug string  `json:"slug"`
		Icon *string `json:"icon"`
	} `gorm:"foreignKey:SubjectID" json:"subject"`
	Grade model.GradeSimple `gorm:"foreignKey:GradeID" json:"grade"`
}

func CreateMentorSubject(c *gin.Context) {
	if c.Request.Method == "POST" {
		user := c.MustGet("user").(jwt.MapClaims)
		var request MentorSubjectRequest
		c.BindJSON(&request)

		if request.SubjectID == 0 || request.GradeID == 0 || request.Method == nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, lib.BaseJsonResponse{
				Code:    http.StatusBadRequest,
				Data:    nil,
				Message: "fill all field",
			})
			return
		}
		m_id, e := uuid.Parse(user["identifier"].(string))
		if e != nil {
			c.JSON(http.StatusUnauthorized, lib.BaseJsonResponse{
				Code:    http.StatusUnauthorized,
				Data:    nil,
				Message: "Internal Server Error. Failed To Get User Authorize",
			})
			return
		}

		isExist := true
		if err_isExist := database.DATABASE.Debug().
			Where("mentor_id = ?", m_id).
			Where("subject_id = ?", request.SubjectID).
			Where("grade_id = ?", request.GradeID).
			First(&model.MentorSubject{}).Error; err_isExist != nil {
			if errors.Is(err_isExist, gorm.ErrRecordNotFound) {
				isExist = false
			}
		}

		if isExist {
			c.JSON(http.StatusBadRequest, lib.BaseJsonResponse{
				Code:    http.StatusBadRequest,
				Data:    nil,
				Message: "Skill Already Exist",
			})
			return
		}
		model := model.MentorSubject{
			MentorID:  m_id,
			SubjectID: request.SubjectID,
			GradeID:   request.GradeID,
			Method:    request.Method,
		}
		if err_create := database.DATABASE.Debug().Create(&model).Error; err_create != nil {
			c.JSON(http.StatusInternalServerError, lib.BaseJsonResponse{
				Code:    http.StatusInternalServerError,
				Data:    nil,
				Message: "internal server error. failed to create a new subject skill",
			})
			return
		}
		c.JSON(http.StatusOK, lib.BaseJsonResponse{
			Code:    http.StatusOK,
			Data:    model,
			Message: "success create new subject skill",
		})
		return
	}
	subject_id := c.Query("subject")
	grade_id := c.Query("grade")
	method_id := c.Query("method")

	var data []response
	query := database.DATABASE.Debug().
		Preload("Mentor", func(tx *gorm.DB) *gorm.DB {
			return tx.Table("mentors")
		}).Joins("JOIN mentors ON mentors.id = mentor_subject.mentor_id").
		Preload("Subject", func(tx *gorm.DB) *gorm.DB {
			return tx.Table("subjects")
		}).
		Preload("Grade")

	if subject_id != "" {
		query.Where("subject_id = ?", subject_id)
	}

	if grade_id != "" {
		query.Where("grade_id = ?", grade_id)
	}

	if method_id != "" {
		query.Where("JSON_CONTAINS(method, " + method_id + ", '$')")
	}
	err := query.Find(&data).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, lib.BaseJsonResponse{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: "Failed To Fetch Data " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, lib.BaseJsonResponse{
		Code:    http.StatusOK,
		Data:    data,
		Message: "success load mentor data",
	})
}
