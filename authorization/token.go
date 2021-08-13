package authorization

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/user0608/mujeresapi/models/authentication"
)

const USERNAME_KEY = "username"
const ROLES_KEY = "roles"

var (
	//ErrInvalidToken Error devuelto cuando el token no es valido
	ErrInvalidToken error = errors.New("token no válido")
	//ErrInvalidClams Error generado cuando los datos contenidos en el clams son incorrectos o están mal estructurados
	ErrInvalidClams error = errors.New("No se pudo obtener la data")
)

//GenerageToken Genera un toquen para un usuario o cliente, el cual se empleara para futuras peticiones
// func GenerageToken(customer model.Customer) (string, error) {
func GenerageToken(usuario authentication.Usuario) (string, error) {
	claim := Clain{
		UserName: usuario.Username,
		Roles:    usuario.GetRoles(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	signedToken, err := token.SignedString(singKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

//ValidateToken verifica la firma de un token para comprobar su valides
func ValidateToken(t string) (Clain, error) {

	token, err := jwt.ParseWithClaims(t, &Clain{}, verifyFunction)
	if err != nil {
		return Clain{}, err
	}
	if !token.Valid {
		return Clain{}, ErrInvalidToken
	}
	claim, ok := token.Claims.(*Clain)
	if !ok {
		return Clain{}, ErrInvalidClams
	}
	return *claim, nil
}
func verifyFunction(token *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}

func itemIsContainIn(item string, values []string) bool {
	for _, val := range values {
		if item == val {
			return true
		}
	}
	return false
}

func JWTMiddleware(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return echo.ErrForbidden
		}
		clain, err := ValidateToken(token)
		if err != nil {
			return echo.ErrForbidden
		}
		c.Set(USERNAME_KEY, clain.UserName)
		c.Set(ROLES_KEY, clain.Roles)
		return f(c)
	}
}

func RolesMiddleware(f echo.HandlerFunc, roles ...string) echo.HandlerFunc {
	return func(c echo.Context) error {
		dato := c.Get(ROLES_KEY)
		if dato == nil {
			return echo.ErrForbidden
		}
		myRoles := dato.([]string)
		for _, rol := range roles {
			if !itemIsContainIn(rol, myRoles) {
				return echo.ErrForbidden
			}
		}
		return f(c)
	}
}
