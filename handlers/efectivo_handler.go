package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/user0608/mujeresapi/authorization"
	"github.com/user0608/mujeresapi/models/auxiliares"
	"github.com/user0608/mujeresapi/services"
	"github.com/user0608/mujeresapi/utils"
)

type EfectivoHandler struct {
	binder  echo.DefaultBinder
	service *services.EfectivoService
}

func NewEfectivoHandler(s *services.EfectivoService) *EfectivoHandler {
	return &EfectivoHandler{service: s, binder: echo.DefaultBinder{}}
}

func (h *EfectivoHandler) Create(c echo.Context) error {
	efectivo := &auxiliares.Efectivo{}
	if err := h.binder.BindBody(c, efectivo); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse("Estructura JSON Invalida"))
	}
	e, err := h.service.Create(efectivo)
	if err != nil {
		if err == utils.ErrDataBaseError {
			return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
		}
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(e))
}

func (h *EfectivoHandler) FindAll(c echo.Context) error {
	efectivos, err := h.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(efectivos))
}
func (h *EfectivoHandler) FindByInstitucionID(c echo.Context) error {
	institucionID, err := strconv.Atoi(c.Param("institucion_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse(""))
	}
	efectivos, err := h.service.GetAllByInstitucionID(institucionID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(efectivos))
}

func (h *EfectivoHandler) FindEfectivosIntitucion(c echo.Context) error {
	usuarioID, ok := c.Get(authorization.USUARIO_ID_KEY).(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	efectivos, err := h.service.GetAllByUsuarioID(usuarioID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(efectivos))
}
