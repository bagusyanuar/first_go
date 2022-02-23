package lib

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type BaseJsonResponse struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"msg"`
}

func SuccessJsonResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, BaseJsonResponse{
		Code:    http.StatusOK,
		Data:    data,
		Message: "success",
	})
}

func ErrorJsonResponse(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, BaseJsonResponse{
		Code:    http.StatusInternalServerError,
		Data:    nil,
		Message: "Internal Server Error (" + msg + ")",
	})
}

func BadRequestJsonResponse(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusBadRequest, BaseJsonResponse{
		Code:    http.StatusBadRequest,
		Data:    nil,
		Message: "Bad Request Validation Form Data",
	})
}
func IsPasswordValid(plainPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

func MakeSlug(text string) string {
	str := []byte(strings.ToLower(text))

	regE := regexp.MustCompile("[[:space:]]")
	str = regE.ReplaceAll(str, []byte("-"))

	regE = regexp.MustCompile("[[:blank:]]")
	str = regE.ReplaceAll(str, []byte(""))

	return string(str)
}

func ValidateRequest(s interface{}) (res []string, e error) {
	validate := validator.New()
	err := validate.Struct(s)
	var errs []string
	if err != nil {
		e_validate := err.(validator.ValidationErrors)

		for _, e := range e_validate {
			errs = append(errs, e.Tag())
		}
		return errs, err
	}
	return errs, nil
}

type Avatar string

func (a *Avatar) Scan(value interface{}) error {
	var result string
	switch v := value.(type) {
	case []byte:
		if string(v) == "" {
			result = "-"
		} else {
			result = "http://localhost:8002/" + string(v)
		}
	}
	
	*a = Avatar(fmt.Sprintf("%v", result))
	return nil
}
