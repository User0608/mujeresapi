package router

import (
	"github.com/labstack/echo/v4"
	"github.com/user0608/mujeresapi/authorization"
	"github.com/user0608/mujeresapi/cmd/api/router/paths"
)

func alertaRouterUpgrade(e *echo.Echo) {
	g := e.Group(paths.ALERTA)
	g.Use(authorization.JWTMiddleware)

}
