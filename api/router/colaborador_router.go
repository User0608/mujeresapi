package router

import (
	"github.com/labstack/echo/v4"
	"github.com/user0608/mujeresapi/api/injectors"
	"github.com/user0608/mujeresapi/api/router/paths"
	"github.com/user0608/mujeresapi/authorization"
	"github.com/user0608/mujeresapi/authorization/roles"
)

func colaboradorRouterUpgrader(e *echo.Echo) {
	h := injectors.GetColaboradorHandler()
	g := e.Group(paths.COLABORADOR)
	g.Use(authorization.JWTMiddleware)
	g.GET("/:id", authorization.RolesMiddleware(h.GetByID, roles.ADMIN))
	g.POST("", authorization.RolesMiddleware(h.CreateColaborador, roles.ADMIN))
}
