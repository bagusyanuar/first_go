package lib

import "errors"

var ErrorBearerType = errors.New("invalid bearer")
var ErrorSignInMethod = errors.New("invalid signin method")
var ErrorJWTClaims = errors.New("invalid jwt claim")
var ErrorJWTParse = errors.New("invalid parse jwt")
var ErrorNoAuthorization = errors.New("invalid Unauthorized")
