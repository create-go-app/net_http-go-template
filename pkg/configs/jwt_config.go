package configs

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/golang-jwt/jwt"
)

// JWTConfig func for
func JWTConfig() jwtmiddleware.Options {
	return jwtmiddleware.Options{
		UserProperty:  "jwt",
		SigningMethod: jwt.SigningMethodHS256,
	}
}
