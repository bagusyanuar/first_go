package middleware

import (
	"first_go/src/lib"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")

	claim, err := lib.ClaimToken(auth)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, lib.BaseJsonResponse{
			Code:    http.StatusUnauthorized,
			Data:    nil,
			Message: strings.Title(err.Error()),
		})
	}
	c.Set("user", claim)
	c.Next()
}

func Mentor(c *gin.Context)  {
	user := c.MustGet("user").(jwt.MapClaims)
	if user["roles"] != "mentor" {
		c.AbortWithStatusJSON(http.StatusForbidden, lib.BaseJsonResponse{
			Code:    http.StatusForbidden,
			Data:    nil,
			Message: "Forbidden Access",
		})
		return
	}
	c.Next()
}

func Member(c *gin.Context)  {
	user := c.MustGet("user").(jwt.MapClaims)
	if user["roles"] != "member" {
		c.AbortWithStatusJSON(http.StatusForbidden, lib.BaseJsonResponse{
			Code:    http.StatusForbidden,
			Data:    nil,
			Message: "Forbidden Access",
		})
		return
	}
	c.Next()
}
