package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/user0608/mujeresapi/authorization"
	"github.com/user0608/mujeresapi/models/control"
	"github.com/user0608/mujeresapi/services"
	"github.com/user0608/mujeresapi/utils"
)

type ColaboradorHandlear struct {
	binder  echo.DefaultBinder
	service *services.ColaboradorService
}

func NewColaboradorHandler(s *services.ColaboradorService) *ColaboradorHandlear {
	return &ColaboradorHandlear{service: s, binder: echo.DefaultBinder{}}
}

func (h *ColaboradorHandlear) CreateColaborador(c echo.Context) error {
	colaborador := &control.Colaborador{}
	if err := h.binder.BindBody(c, colaborador); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse("Estructura JSON incorrecta."))
	}
	if err := h.service.Asociar(colaborador); err != nil {
		log.Println(err.Error())
		if err == utils.ErrInvalidData {
			return c.JSON(http.StatusBadRequest, utils.NewBadResponse("Datos incorrectos"))
		}
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(colaborador))
}

func (h *ColaboradorHandlear) GetByID(c echo.Context) error {
	colaboradorID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse("ID invalido"))
	}
	colaborador, err := h.service.FindByID(colaboradorID)
	log.Println(colaboradorID)
	if err != nil {
		log.Println(err.Error())
		if err == utils.ErrNothingFind {
			return c.JSON(http.StatusNotFound, utils.NewBadResponse("No se encontro el registro"))
		}
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(colaborador))
}
func (h *ColaboradorHandlear) FindAll(c echo.Context) error {
	colaboradores, err := h.service.AllColaboradores()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(colaboradores))
}
func (h *ColaboradorHandlear) ME(c echo.Context) error {
	id, ok := c.Get(authorization.USUARIO_ID_KEY).(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	if res, err := h.service.ByUserID(id); err != nil {
		if err == utils.ErrDataBaseError {
			return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
		}
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse(err.Error()))
	} else {
		return c.JSON(http.StatusOK, utils.NewOkResponse(res))
	}
}
