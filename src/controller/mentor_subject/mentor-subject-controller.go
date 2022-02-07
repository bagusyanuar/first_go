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

func CreateMentorSubject(c *gin.Context) {
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
}
