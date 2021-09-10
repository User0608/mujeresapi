package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/user0608/mujeresapi/services"
	"github.com/user0608/mujeresapi/utils"
)

type RolesHandlear struct {
	service *services.RoleService
}

func NewRolesHandler(s *services.RoleService) *RolesHandlear {
	return &RolesHandlear{service: s}
}

func (h *RolesHandlear) Roles(c echo.Context) error {
	roles, err := h.service.FindRoles()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(roles))
}
