package router

import "github.com/labstack/echo/v4"

func Upgrade(e *echo.Echo) {
	usuarioRouterUpgrade(e)
	alertaRouterUpgrade(e)
	multimediaRouterUpgrade(e)
	personaRouterUpgrade(e)
	rolesRouterUpgrader(e)
	colaboradorRouterUpgrader(e)
	institucionRouterUpgrade(e)
	notificacionUpgrader(e)
	efectivoRouterUpgrade(e)
	asignacionRouterUpgrade(e)
	reporteRouterUpgrader(e)
}
