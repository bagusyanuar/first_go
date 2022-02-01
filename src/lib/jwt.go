package lib

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type JWTClaims struct {
	jwt.StandardClaims
	Unique     uuid.UUID `json:"unique"`
	Identifier uuid.UUID `json:"identifier"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Role       string    `json:"roles"`
}

func GenerateToken(unique uuid.UUID, identifier uuid.UUID, username string, email string, role string) (string, error) {
	claims := JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer: "JONI_APP",
		},
		Unique:     unique,
		Username:   username,
		Email:      email,
		Role:       role,
		Identifier: identifier,
	}

	token := jwt.NewWithClaims(JWTSigninMethod, claims)
	signedToken, err := token.SignedString([]byte(JWTSignatureKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ClaimToken(auth string) (interface{}, error) {
	if auth == "" {
		return nil, ErrorNoAuthorization
	}
	bearer := string(auth[0:7])
	token := string(auth[7:])

	if bearer != "Bearer " {
		return nil, ErrorBearerType
	}

	vToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrorSignInMethod
		} else if method != JWTSigninMethod {
			return nil, ErrorSignInMethod
		}
		return []byte(JWTSignatureKey), nil
	})

	if err != nil {
		return nil, ErrorJWTParse
	}

	claim, ok := vToken.Claims.(jwt.MapClaims)
	if !ok || !vToken.Valid {
		return nil, ErrorJWTClaims
	}
	return claim, nil
}
