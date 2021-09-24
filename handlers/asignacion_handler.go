package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/user0608/mujeresapi/models/control"
	"github.com/user0608/mujeresapi/services"
	"github.com/user0608/mujeresapi/utils"
)

type AsignacionHandler struct {
	service *services.AsignacionService
	binder  echo.DefaultBinder
}

func NewAsinacinoHandler(s *services.AsignacionService) *AsignacionHandler {
	return &AsignacionHandler{service: s, binder: echo.DefaultBinder{}}
}

func (h *AsignacionHandler) Create(c echo.Context) error {
	a := control.Asignacion{}
	if err := h.binder.BindBody(c, &a); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse("JSON Invalido"))
	}
	if asignacion, err := h.service.Create(a); err != nil {
		if err == utils.ErrDataBaseError {
			return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
		}
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse(err.Error()))
	} else {
		return c.JSON(http.StatusOK, utils.NewOkResponse(asignacion))
	}
}
func (h *AsignacionHandler) Update(c echo.Context) error {
	a := control.Asignacion{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewBadResponse("ID invalido"))
	}
	if err := h.binder.BindBody(c, &a); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse("JSON Invalido"))
	}
	if asignacion, err := h.service.Update(id, a); err != nil {
		if err == utils.ErrDataBaseError {
			return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
		}
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse(err.Error()))
	} else {
		return c.JSON(http.StatusOK, utils.NewOkResponse(asignacion))
	}
}

func (h *AsignacionHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewBadResponse("ID invalido"))
	}
	if err := h.service.Delete(id); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	return c.JSON(http.StatusOK, utils.NewOkMesage("Eliminado correctamente"))
}

func (h *AsignacionHandler) GetAllByEfectivoID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewBadResponse("ID invalido"))
	}
	if asignacion, err := h.service.GetAllByEfectivoID(id); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	} else {
		return c.JSON(http.StatusOK, utils.NewOkResponse(asignacion))
	}
}
func (h *AsignacionHandler) GetAllByNotificatioID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewBadResponse("ID invalido"))
	}
	if asignacion, err := h.service.GetAllByNotificatioID(id); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	} else {
		return c.JSON(http.StatusOK, utils.NewOkResponse(asignacion))
	}
}
func (h *AsignacionHandler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewBadResponse("ID invalido"))
	}
	if asignacion, err := h.service.GetByID(id); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	} else {
		return c.JSON(http.StatusOK, utils.NewOkResponse(asignacion))
	}
}
