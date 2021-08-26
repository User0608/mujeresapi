package authorization

import "github.com/dgrijalva/jwt-go"

type Clain struct {
	UsuarioID int      `json:"usuario_id"`
	UserName  string   `json:"username"`
	Roles     []string `json:"roles"`
	jwt.StandardClaims
}
