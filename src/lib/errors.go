package lib

import (
	"errors"
	"net/http"

	"gorm.io/gorm"
)

var ErrorBearerType = errors.New("invalid bearer type")
var ErrorSignInMethod = errors.New("invalid signin method")
var ErrorJWTClaims = errors.New("invalid jwt claim")
var ErrorJWTParse = errors.New("invalid parse jwt")
var ErrorNoAuthorization = errors.New("invalid unauthorized")
var ErrorInvalidPassword = errors.New("password did not match")

func SignInReturnErrors(err error) map[string]interface{} {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return map[string]interface{}{
			"code":    http.StatusUnauthorized,
			"data":    nil,
			"message": "User Not Found!",
		}
	} else if errors.Is(err, ErrorInvalidPassword) {
		return map[string]interface{}{
			"code":    http.StatusUnauthorized,
			"data":    nil,
			"message": "Password Did Not Match",
		}
	} else {
		return map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"data":    nil,
			"message": "Error While Sign In " + err.Error(),
		}
	}
}
