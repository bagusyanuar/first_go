package users

import (
	"encoding/json"
	"first_go/database"
	"first_go/src/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *gin.Context) {
	var users []model.User
	if c.Request.Method == "POST" {
		username := c.PostForm("username")
		password := c.PostForm("password")
		email := c.PostForm("email")
		name := c.PostForm("name")
		roles, _ := json.Marshal([]string{"ROLE_MEMBER"})
		provider, _ := json.Marshal([]string{"app"})

		hashedPassword, errHashed := bcrypt.GenerateFromPassword([]byte(password), 13)
		if errHashed != nil {
			c.AbortWithError(http.StatusInternalServerError, errHashed)
		}
		var vPassword string = string(hashedPassword)
		user := model.User{
			Email:    email,
			Username: username,
			Password: &vPassword,
			Roles:    roles,
			Provider: provider,
		}

		member := model.Member{
			Name: name,
		}

		userMember := model.UserMember{
			User:   user,
			Member: member,
		}

		if err := database.DATABASE.Create(&userMember).Error; err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"data":   userMember,
		})
		return
	}
	if err := database.DATABASE.Find(&users).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   users,
	})
}
