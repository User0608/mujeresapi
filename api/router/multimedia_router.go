package router

import (
	"github.com/labstack/echo/v4"
	"github.com/user0608/mujeresapi/api/injectors"
	"github.com/user0608/mujeresapi/api/router/paths"
)

func multimediaRouterUpgrade(e *echo.Echo) {
	h := injectors.GetMultimediaHandler()
	g := e.Group(paths.MULTIMEDIA)
	g.POST("/files", h.UploadMultimedia)
	//g.GET("/ws", authorization.RolesMiddleware(h.LoadFileHandler, roles.APP_ROLE), authorization.JWTMiddlewareQueryTokenParam)
}
