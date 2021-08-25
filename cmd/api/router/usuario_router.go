package router

import (
	"github.com/labstack/echo/v4"
	"github.com/user0608/mujeresapi/authorization"
	"github.com/user0608/mujeresapi/cmd/api/injectors"
	"github.com/user0608/mujeresapi/cmd/api/router/paths"
)

func usuarioRouterUpgrade(e *echo.Echo) {
	handlers := injectors.GetUsuarioHandler()

	uGroup := e.Group(paths.APP_USER)
	uGroup.Use(authorization.JWTMiddleware)

	e.POST("/v1/login", handlers.LogginUser)
	e.POST("/v1/registrar", handlers.RegistrarAppUser)

	uGroup.GET("", authorization.RolesMiddleware(handlers.AllUsersHandler, authorization.ADMIN))
	uGroup.POST("/detalle", authorization.RolesMiddleware(handlers.CrateUpdateUser(true), authorization.ADMIN, authorization.APP_ROLE))
	uGroup.PUT("/detalle", authorization.RolesMiddleware(handlers.CrateUpdateUser(false), authorization.ADMIN, authorization.APP_ROLE))
	uGroup.GET("/detalle/:usuario_id", authorization.RolesMiddleware(handlers.GetUserAppById, authorization.ADMIN, authorization.APP_ROLE))

}
