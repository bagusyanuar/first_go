package controller

import (
	"encoding/json"
	"first_go/database"
	"first_go/src/lib"
	"first_go/src/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func SuperAdminSeeder(c *gin.Context) {

	roles, _ := json.Marshal([]string{"superadmin"})
	provider, _ := json.Marshal([]string{"app"})
	hash, errHashing := bcrypt.GenerateFromPassword([]byte("superadmin2"), 13)
	if errHashing != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusBadRequest,
			"data":    nil,
			"message": "Bad Request! Failed Hashing Password",
		})
		return
	}
	password := string(hash)

	tx := database.DATABASE.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	superadmin := model.UserAdmin{
		User: model.User{
			ID:       uuid.New(),
			Email:    "superadmin2@gmail.com",
			Password: &password,
			Provider: provider,
			Roles:    roles,
		},
		Admin: model.Admin{
			ID:   uuid.New(),
			Name: "Super Admin",
		},
	}

	if err := tx.Create(&superadmin).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    nil,
			"message": "Error Insert",
		})
		return
	}

	accessToken, errorTokenize := lib.GenerateToken(superadmin.ID, superadmin.Admin.ID, "superadmin", superadmin.Email, "superadmin")
	if errorTokenize != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    nil,
			"message": "Error While Generate Token " + errorTokenize.Error(),
		})
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": map[string]interface{}{
			"accessToken": accessToken,
		},
		"message": "Super Admin Created",
	})
}
