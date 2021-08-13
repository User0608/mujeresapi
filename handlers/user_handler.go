package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/user0608/mujeresapi/services"
	"github.com/user0608/mujeresapi/utils"
)

type UsuarioHandlear struct {
	service *services.UsuarioService
}

func NewUsuarioHandler(s *services.UsuarioService) *UsuarioHandlear {
	return &UsuarioHandlear{
		service: s,
	}
}

func (u *UsuarioHandlear) LogginUser(c echo.Echo) error {
	return nil
}

func (u *UsuarioHandlear) AllUsersHandler(c echo.Context) error {
	usuarios, err := u.service.AllUsuarios()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewBadResponse(utils.InternalMessage))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(usuarios))
}
