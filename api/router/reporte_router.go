package router

import (
	"github.com/labstack/echo/v4"
	"github.com/user0608/mujeresapi/api/injectors"
	"github.com/user0608/mujeresapi/api/router/paths"
	"github.com/user0608/mujeresapi/authorization"
	"github.com/user0608/mujeresapi/authorization/roles"
)

func reporteRouterUpgrader(e *echo.Echo) {
	h := injectors.GetReporteHandler()
	g := e.Group(paths.REPORTE)
	g.Use(authorization.JWTMiddleware)
	g.POST("", authorization.RolesMiddleware(h.Create, roles.INSTITUCION_ROLE))
	g.PUT("/:id", authorization.RolesMiddleware(h.Update, roles.INSTITUCION_ROLE))
	g.DELETE("/:id", authorization.RolesMiddleware(h.Delete, roles.INSTITUCION_ROLE))
	g.GET("", authorization.RolesMiddleware(h.FindAll, roles.INSTITUCION_ROLE, roles.ADMIN, roles.CONTROL_ROLE))
	g.GET("/:id", authorization.RolesMiddleware(h.FindByID, roles.INSTITUCION_ROLE, roles.ADMIN, roles.CONTROL_ROLE))
	g.GET("/notificacion/:id", authorization.RolesMiddleware(h.FindByNotificaconID, roles.INSTITUCION_ROLE, roles.ADMIN, roles.CONTROL_ROLE))
}
