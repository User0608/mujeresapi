package router

import (
	"github.com/labstack/echo/v4"
	"github.com/user0608/mujeresapi/api/injectors"
	"github.com/user0608/mujeresapi/api/router/paths"
	"github.com/user0608/mujeresapi/authorization"
	"github.com/user0608/mujeresapi/authorization/roles"
)

func asignacionRouterUpgrade(e *echo.Echo) {
	h := injectors.GetAsignacionHandler()
	g := e.Group(paths.ASIGNACION)
	g.Use(authorization.JWTMiddleware)

	g.POST("", authorization.RolesMiddleware(h.Create, roles.INSTITUCION_ROLE))
	g.GET("/:id", authorization.RolesMiddleware(h.GetByID, roles.INSTITUCION_ROLE))
	g.GET("/efectivo/:id", authorization.RolesMiddleware(h.GetAllByEfectivoID, roles.INSTITUCION_ROLE))
	g.GET("/notificacion/:id", authorization.RolesMiddleware(h.GetAllByNotificatioID, roles.INSTITUCION_ROLE))
	g.PUT("/:id", authorization.RolesMiddleware(h.Update, roles.INSTITUCION_ROLE))
	g.DELETE("/:id", authorization.RolesMiddleware(h.Delete, roles.INSTITUCION_ROLE))
}
