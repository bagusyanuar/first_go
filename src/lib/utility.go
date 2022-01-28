package lib

import (
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type BaseJsonResponse struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"msg"`
}

func IsPasswordValid(plainPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

func MakeSlug(text string) string  {
	str := []byte(strings.ToLower(text))

	regE := regexp.MustCompile("[[:space:]]")
	str = regE.ReplaceAll(str, []byte("-"))

	regE = regexp.MustCompile("[[:blank:]]")
	str = regE.ReplaceAll(str, []byte(""))

	return string(str)
}
