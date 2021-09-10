package router

import (
	"github.com/labstack/echo/v4"
	"github.com/user0608/mujeresapi/api/injectors"
	"github.com/user0608/mujeresapi/api/router/paths"
)

func rolesRouterUpgrader(e *echo.Echo) {
	h := injectors.GetRolesHandler()
	e.GET(paths.ROLES, h.Roles)
}
