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
	g.GET("/me", authorization.RolesMiddleware(h.ME, roles.CONTROL_ROLE))
	g.GET("/:id", authorization.RolesMiddleware(h.GetByID, roles.ADMIN))
	g.GET("", authorization.RolesMiddleware(h.FindAll, roles.ADMIN, roles.CONTROL_ROLE))
	g.POST("", authorization.RolesMiddleware(h.CreateColaborador, roles.ADMIN))

}
