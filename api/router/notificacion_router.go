package router

import (
	"github.com/labstack/echo/v4"
	"github.com/user0608/mujeresapi/api/injectors"
	"github.com/user0608/mujeresapi/api/router/paths"
	"github.com/user0608/mujeresapi/authorization"
	"github.com/user0608/mujeresapi/authorization/roles"
)

func notificacionUpgrader(e *echo.Echo) {
	h := injectors.GetNotificationHandler()
	g := e.Group(paths.NOTIFICACION)
	g.Use(authorization.JWTMiddleware)
	g.POST("", authorization.RolesMiddleware(h.Create, roles.CONTROL_ROLE))
	g.PUT("", authorization.RolesMiddleware(h.Update, roles.CONTROL_ROLE))
	g.GET("", authorization.RolesMiddleware(h.FindAll, roles.ADMIN, roles.CONTROL_ROLE))
	g.GET("/:id", authorization.RolesMiddleware(h.FindByID, roles.ADMIN, roles.CONTROL_ROLE, roles.INSTITUCION_ROLE))
	g.GET("/institucion/:institucion_id", authorization.RolesMiddleware(h.FindByInstitucionID, roles.ADMIN, roles.CONTROL_ROLE))
	g.GET("/colaborador/:colaborador_id", authorization.RolesMiddleware(h.FindByColaboradorID, roles.ADMIN, roles.CONTROL_ROLE))
	//path reservado, solo para institucion
	g.GET("/institucion", authorization.RolesMiddleware(h.FindNotificationByInstitucion, roles.INSTITUCION_ROLE))

}
