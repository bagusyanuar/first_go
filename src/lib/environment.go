package lib

import "github.com/dgrijalva/jwt-go"

var JWTSigninMethod = jwt.SigningMethodHS256
var JWTSignatureKey string = "ONLYGODKNOWS"
