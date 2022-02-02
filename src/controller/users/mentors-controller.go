package users

import (
	"errors"
	"first_go/database"
	"first_go/src/lib"
	"first_go/src/response"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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
