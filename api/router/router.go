package router

import "github.com/labstack/echo/v4"

func Upgrade(e *echo.Echo) {
	usuarioRouterUpgrade(e)
	alertaRouterUpgrade(e)
	multimediaRouterUpgrade(e)
}
