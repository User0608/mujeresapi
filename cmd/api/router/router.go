package router

import (
	"github.com/labstack/echo/v4"
	"github.com/user0608/mujeresapi/cmd/api/injectors"
)

func RoutesUsuario(e *echo.Echo) {
	uGroup := e.Group("/v1/usuario")
	handlers := injectors.GetNewUsuarioHandler()
	uGroup.GET("", handlers.AllUsersHandler)
}
