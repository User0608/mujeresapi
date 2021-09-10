package router

import (
	"github.com/labstack/echo/v4"
	"github.com/user0608/mujeresapi/api/injectors"
	"github.com/user0608/mujeresapi/api/router/paths"
	"github.com/user0608/mujeresapi/authorization"
	"github.com/user0608/mujeresapi/authorization/roles"
)

func personaRouterUpgrade(e *echo.Echo) {
	h := injectors.GetPersonaHandler()
	g := e.Group(paths.PERSONA)
	g.Use(authorization.JWTMiddleware)
	g.GET("", authorization.RolesMiddleware(h.FindAll, roles.ADMIN))
	g.GET("/:persona_id", authorization.RolesMiddleware(h.FindByID, roles.ADMIN))
	g.POST("", authorization.RolesMiddleware(h.Register, roles.ADMIN))
	g.PUT("", authorization.RolesMiddleware(h.Update, roles.ADMIN))
}
