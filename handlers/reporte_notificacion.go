package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/user0608/mujeresapi/models/control"
	"github.com/user0608/mujeresapi/services"
	"github.com/user0608/mujeresapi/utils"
)

type ReporteHandler struct {
	service *services.ReporteService
	binder  echo.DefaultBinder
}

func NewReporteHandler(s *services.ReporteService) *ReporteHandler {
	return &ReporteHandler{service: s, binder: echo.DefaultBinder{}}
}
func (h *ReporteHandler) HandlerErrorMessage(c echo.Context, err error) error {
	if err == utils.ErrDataBaseError {
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	if err == utils.ErrNothingFind {
		return c.JSON(http.StatusNotFound, utils.NewNoFindResponse("/<ID> invalido, no se encontro el registro"))
	}
	return c.JSON(http.StatusBadRequest, utils.NewBadResponse(err.Error()))
}
func (h *ReporteHandler) Create(c echo.Context) error {
	r := control.Reporte{}
	if err := h.binder.BindBody(c, &r); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse("Estructura JSON invalida"))
	}
	if rr, err := h.service.Create(r); err != nil {
		return h.HandlerErrorMessage(c, err)
	} else {
		return c.JSON(http.StatusOK, utils.NewOkResponse(rr))
	}
}
func (h *ReporteHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewNoFindResponse("/<ID> invalido, no se encontro el registro"))
	}
	r := control.Reporte{}
	if err := h.binder.BindBody(c, &r); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse("Estructura JSON invalida"))
	}
	if rr, err := h.service.Update(id, r); err != nil {
		return h.HandlerErrorMessage(c, err)
	} else {
		return c.JSON(http.StatusOK, utils.NewOkResponse(rr))
	}
}

func (h *ReporteHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewNoFindResponse("/<ID> invalido, no se encontro el registro"))
	}
	if err := h.service.Delete(id); err != nil {
		return h.HandlerErrorMessage(c, err)
	}
	return c.JSON(http.StatusOK, utils.NewOkMesage("Registro eliminado"))

}
func (h *ReporteHandler) FindAll(c echo.Context) error {
	if repos, err := h.service.FindAll(); err != nil {
		return h.HandlerErrorMessage(c, err)
	} else {
		return c.JSON(http.StatusOK, utils.NewOkResponse(repos))
	}
}
func (h *ReporteHandler) FindByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewNoFindResponse("/<ID> invalido, no se encontro el registro"))
	}
	if repo, err := h.service.FindByID(id); err != nil {
		return h.HandlerErrorMessage(c, err)
	} else {
		return c.JSON(http.StatusOK, utils.NewOkResponse(repo))
	}
}
func (h *ReporteHandler) FindByNotificaconID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewNoFindResponse("/<ID> invalido, no se encontro el registro"))
	}
	if repo, err := h.service.FindByNotificacionID(id); err != nil {
		return h.HandlerErrorMessage(c, err)
	} else {
		return c.JSON(http.StatusOK, utils.NewOkResponse(repo))
	}
}
