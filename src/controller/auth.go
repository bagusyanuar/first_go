package controller

import (
	"encoding/json"
	"first_go/database"
	"first_go/src/lib"
	"first_go/src/model"
	"first_go/src/repository"
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

func AdminSignUp(c *gin.Context) {
	email := c.PostForm("email")
	username := c.PostForm("username")
	password := c.PostForm("password")
	name := c.PostForm("name")

	roles, _ := json.Marshal([]string{"admin"})
	provider, _ := json.Marshal([]string{"app"})
	hash, errHashing := bcrypt.GenerateFromPassword([]byte(password), 13)
	if errHashing != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusBadRequest,
			"data":    nil,
			"message": "Bad Request! Failed Hashing Password",
		})
		return
	}
	var vPassword string = string(hash)

	tx := database.DATABASE.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	user := model.User{
		Email:    email,
		Username: username,
		Password: &vPassword,
		Roles:    roles,
		Provider: provider,
	}

	admin := model.Admin{
		Name: name,
	}

	userAdmin := model.UserAdmin{
		User:  user,
		Admin: admin,
	}

	if err := tx.Create(&userAdmin).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    nil,
			"message": "Error While Creating Account",
		})
		return
	}

	accessToken, errorTokenize := lib.GenerateToken(userAdmin.ID, userAdmin.Admin.ID, userAdmin.Username, userAdmin.Email, "admin")
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
		"message": "Success Sign Up",
	})
}

func MemberSignUp(c *gin.Context) {
	email := c.PostForm("email")
	username := c.PostForm("username")
	password := c.PostForm("password")
	name := c.PostForm("name")

	roles, _ := json.Marshal([]string{"member"})
	provider, _ := json.Marshal([]string{"app"})
	hash, errHashing := bcrypt.GenerateFromPassword([]byte(password), 13)
	if errHashing != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusBadRequest,
			"data":    nil,
			"message": "Bad Request! Failed Hashing Password",
		})
		return
	}
	var vPassword string = string(hash)

	tx := database.DATABASE.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
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

	if err := tx.Create(&userMember).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    nil,
			"message": "Error Insert",
		})
		return
	}

	accessToken, errorTokenize := lib.GenerateToken(userMember.ID, userMember.Member.ID, userMember.Username, userMember.Email, "member")
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
		"message": "Success Sign Up",
	})
}

func MemberSignIn(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	provider := c.PostForm("provider")
	user, err := repository.SignInMember(&model.UserMember{}, email, password, provider)
	if err != nil {
		errResponse := lib.SignInReturnErrors(err)
		c.JSON(errResponse["code"].(int), errResponse)
		return
	}

	accessToken, errTokenize := lib.GenerateToken(user.ID, user.Member.ID, user.Username, user.Email, "member")
	if errTokenize != nil {
		c.JSON(http.StatusInternalServerError, lib.BaseJsonResponse{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: "Error While Getting Access Token " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, lib.BaseJsonResponse{
		Code: http.StatusOK,
		Data: map[string]interface{}{
			"accessToken": accessToken,
		},
		Message: "Success Sign In",
	})
}

func MentorSignUp(c *gin.Context) {
	email := c.PostForm("email")
	username := c.PostForm("username")
	password := c.PostForm("password")
	name := c.PostForm("name")

	roles, _ := json.Marshal([]string{"mentor"})
	provider, _ := json.Marshal([]string{"app"})
	hash, errHashing := bcrypt.GenerateFromPassword([]byte(password), 13)
	if errHashing != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusBadRequest,
			"data":    nil,
			"message": "Bad Request! Failed Hashing Password",
		})
		return
	}
	var vPassword string = string(hash)

	tx := database.DATABASE.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	user := model.User{
		Email:    email,
		Username: username,
		Password: &vPassword,
		Roles:    roles,
		Provider: provider,
	}

	mentor := model.Mentor{
		Name: name,
	}

	userMentor := model.UserMentor{
		User:   user,
		Mentor: mentor,
	}

	if err := tx.Create(&userMentor).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    nil,
			"message": "Error While Creating Mentor Account",
		})
		return
	}

	accessToken, errorTokenize := lib.GenerateToken(userMentor.ID, userMentor.Mentor.ID, userMentor.Username, userMentor.Email, "mentor")
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
		"message": "Success Create Account",
	})
}

func MentorSignIn(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	provider := c.PostForm("provider")
	user, err := repository.SignInMentor(&model.UserMentor{}, email, password, provider)
	if err != nil {
		errResponse := lib.SignInReturnErrors(err)
		c.JSON(errResponse["code"].(int), errResponse)
		return
	}

	accessToken, errTokenize := lib.GenerateToken(user.ID, user.Mentor.ID, user.Username, user.Email, "mentor")
	if errTokenize != nil {
		c.JSON(http.StatusInternalServerError, lib.BaseJsonResponse{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: "Error While Getting Access Token " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, lib.BaseJsonResponse{
		Code: http.StatusOK,
		Data: map[string]interface{}{
			"accessToken": accessToken,
		},
		Message: "Success Sign In",
	})
}
