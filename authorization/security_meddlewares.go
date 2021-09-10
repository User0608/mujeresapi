package authorization

import (
	"log"

	"github.com/labstack/echo/v4"

	"github.com/user0608/mujeresapi/authorization/roles"
)

const (
	USERNAME_KEY    = "username"
	ROLES_KEY       = "roles"
	USUARIO_ID_KEY  = "usuario_id"
	IS_APP_USER_KEY = "app_user"
)

func itemIsContainIn(item string, values []string) bool {
	for _, val := range values {
		if item == val {
			return true
		}
	}
	return false
}
func jwtValidToken(token string, c echo.Context) error {
	if token == "" {
		return echo.ErrForbidden
	}
	clain, err := ValidateToken(token)
	if err != nil {
		log.Println("El token no es valido")
		return echo.ErrForbidden
	}
	c.Set(USERNAME_KEY, clain.UserName)
	c.Set(USUARIO_ID_KEY, clain.UsuarioID)
	c.Set(ROLES_KEY, clain.Roles)
	return nil
}
func JWTMiddleware(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if err := jwtValidToken(token, c); err != nil {
			return echo.ErrForbidden
		}
		return f(c)
	}
}
func JWTMiddlewareQueryTokenParam(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.QueryParam("token")
		if err := jwtValidToken(token, c); err != nil {
			log.Println("Conexión rechazada", c.Request().Host)
			return err
		}
		return f(c)
	}
}

func RolesMiddleware(f echo.HandlerFunc, rls ...string) echo.HandlerFunc {
	return func(c echo.Context) error {
		dato := c.Get(ROLES_KEY)
		if dato == nil {
			log.Println("Roles no encontrados")
			return echo.ErrForbidden
		}
		myRoles := dato.([]string)
		var flag bool
		for _, r := range rls {
			if itemIsContainIn(r, myRoles) {
				if r == roles.APP_ROLE {
					c.Set(IS_APP_USER_KEY, "OK")
				}
				flag = true
			}
		}
		if !flag {
			log.Println("Petission no autorozada: Host", c.Request().Host)
			return echo.ErrForbidden
		}
		return f(c)
	}
}
