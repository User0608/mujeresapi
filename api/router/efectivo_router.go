package router

import (
	"github.com/labstack/echo/v4"
	"github.com/user0608/mujeresapi/api/injectors"
	"github.com/user0608/mujeresapi/api/router/paths"
	"github.com/user0608/mujeresapi/authorization"
	"github.com/user0608/mujeresapi/authorization/roles"
)

func efectivoRouterUpgrade(e *echo.Echo) {
	h := injectors.GetEfectivoHandler()
	g := e.Group(paths.EFECTIVO)
	g.Use(authorization.JWTMiddleware)
	g.POST("", authorization.RolesMiddleware(h.Create, roles.INSTITUCION_ROLE))
	g.GET("", authorization.RolesMiddleware(h.FindAll, roles.ADMIN, roles.CONTROL_ROLE))
	g.GET("/institucion/:institucion_id", authorization.RolesMiddleware(h.FindByInstitucionID, roles.ADMIN, roles.CONTROL_ROLE))
	g.GET("/institucion", authorization.RolesMiddleware(h.FindEfectivosIntitucion, roles.INSTITUCION_ROLE))
}
