package middleware

import (
	"first_go/src/lib"
	"net/http"
	"strings"

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
