package router

import (
	"github.com/labstack/echo/v4"
	"github.com/user0608/mujeresapi/api/injectors"
	"github.com/user0608/mujeresapi/api/router/paths"
	"github.com/user0608/mujeresapi/authorization"
	"github.com/user0608/mujeresapi/authorization/roles"
)

func institucionRouterUpgrade(e *echo.Echo) {
	h := injectors.GetInstitucionHandler()
	g := e.Group(paths.INSTUTUCION)
	g.Use(authorization.JWTMiddleware)
	g.POST("", authorization.RolesMiddleware(h.Create, roles.ADMIN, roles.CONTROL_ROLE))
	g.PUT("", authorization.RolesMiddleware(h.Update, roles.ADMIN, roles.CONTROL_ROLE))
	g.GET("", authorization.RolesMiddleware(h.Find, roles.ADMIN, roles.CONTROL_ROLE))
	g.GET("/:id", authorization.RolesMiddleware(h.FindByID, roles.ADMIN, roles.CONTROL_ROLE))
	g.GET("/username/:username", authorization.RolesMiddleware(h.FindByUserName, roles.ADMIN))
	g.GET("/me", authorization.RolesMiddleware(h.FindMyInfo, roles.INSTITUCION_ROLE))
}
