package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/user0608/mujeresapi/authorization"
	"github.com/user0608/mujeresapi/models/application"
	"github.com/user0608/mujeresapi/services"
	"github.com/user0608/mujeresapi/utils"
)

type AlertaHandlear struct {
	service *services.AlertaService
}

func NewAlertaHandler(s *services.AlertaService) *AlertaHandlear {
	return &AlertaHandlear{
		service: s,
	}
}
func IfRequestIsAppUserGetUserID(c echo.Context) (int, bool) {
	isAppUser := false
	userId := 0
	if r := c.Get(authorization.IS_APP_USER_KEY); r != nil {
		if r.(string) == "OK" {
			isAppUser = true
		}
	}
	if r := c.Get(authorization.USUARIO_ID_KEY); r != nil {
		userId = r.(int)
	}
	return userId, isAppUser
}
func (h *AlertaHandlear) FindByUsuarioID(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse(""))
	}
	alertas, err := h.service.GatAllWithUserID(ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(alertas))
}
func (h *AlertaHandlear) FindAllHandler(c echo.Context) error {
	userId, isAppUser := IfRequestIsAppUserGetUserID(c)
	var alertas []application.Alerta
	var err error
	if isAppUser {
		alertas, err = h.service.GatAllWithUserID(userId)
	} else {
		alertas, err = h.service.GatAll()
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(alertas))
}
func (h *AlertaHandlear) FindByIDHandler(c echo.Context) error {
	alertaId := c.Param("alerta_id")
	userId, isAppUser := IfRequestIsAppUserGetUserID(c)
	alerta, err := h.service.GatByID(alertaId)
	if err != nil {
		if err == utils.ErrNothingFind {
			return c.JSON(http.StatusBadRequest, utils.NewNoFindResponse(""))
		}
	}
	if isAppUser && alerta.UsuarioID != userId {
		return c.JSON(http.StatusForbidden, utils.NewForbiddenResponse(""))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(alerta))
}
func (h *AlertaHandlear) CreateHandler(c echo.Context) error {
	userID := c.Get(authorization.USUARIO_ID_KEY).(int)
	alerta := &application.Alerta{UsuarioID: userID, Estado: "enviado"}
	binder := echo.DefaultBinder{}
	if err := binder.BindBody(c, alerta); err != nil {
		return echo.ErrBadRequest
	}
	if err := h.service.Create(alerta); err != nil {
		if err == utils.ErrInvalidData {
			log.Println("CreateHandler:", err.Error())
			return c.JSON(
				http.StatusBadRequest,
				utils.NewBadResponse("Tipo o formato de datos incorrecto."),
			)
		} else {
			return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
		}
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(alerta))
}
