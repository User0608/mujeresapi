package authorization

import "github.com/dgrijalva/jwt-go"

type Clain struct {
	UserName string
	Roles    []string
	jwt.StandardClaims
}
