package router

import (
	"github.com/labstack/echo/v4"
	"github.com/user0608/mujeresapi/api/injectors"
	"github.com/user0608/mujeresapi/api/router/paths"
	"github.com/user0608/mujeresapi/authorization"
	"github.com/user0608/mujeresapi/authorization/roles"
)

func alertaRouterUpgrade(e *echo.Echo) {
	alerta := injectors.GetAlertaHandler()
	g := e.Group(paths.ALERTA)
	g.Use(authorization.JWTMiddleware)
	g.POST("", authorization.RolesMiddleware(alerta.CreateHandler, roles.APP_ROLE))
	g.GET("/usuario/:id", authorization.RolesMiddleware(alerta.FindByUsuarioID, roles.ADMIN, roles.CONTROL_ROLE))
	g.GET("/:alerta_id", authorization.RolesMiddleware(alerta.FindByIDHandler, roles.ADMIN, roles.CONTROL_ROLE, roles.APP_ROLE))
	g.GET("", authorization.RolesMiddleware(alerta.FindAllHandler, roles.ADMIN, roles.CONTROL_ROLE, roles.APP_ROLE))
}
