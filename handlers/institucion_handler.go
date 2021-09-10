package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/user0608/mujeresapi/authorization"
	"github.com/user0608/mujeresapi/models/control"
	"github.com/user0608/mujeresapi/services"
	"github.com/user0608/mujeresapi/utils"
)

type InstitucionHandlear struct {
	service *services.InsitucionService
	binder  echo.DefaultBinder
}

func NewInstitucionHandler(s *services.InsitucionService) *InstitucionHandlear {
	return &InstitucionHandlear{service: s, binder: echo.DefaultBinder{}}
}
func (h *InstitucionHandlear) Create(c echo.Context) error {
	return h.createOrUpdateAction(c, true)
}
func (h *InstitucionHandlear) Update(c echo.Context) error {
	return h.createOrUpdateAction(c, false)
}
func (h *InstitucionHandlear) createOrUpdateAction(c echo.Context, create bool) error {
	institucion := &control.Institucion{}
	if err := h.binder.BindBody(c, institucion); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse("Estructura de datos incorrecta"))
	}
	var err error
	if create {
		err = h.service.CreateInstitucion(institucion)
	} else {
		err = h.service.UpdateInstitucion(institucion)
		if err == utils.ErrNothingFind {
			return c.JSON(http.StatusBadRequest, utils.NewBadResponse(err.Error()))
		}
	}
	if err != nil {
		if err == utils.ErrUsuarioNoDisponible || err == utils.ErrUsuarioNotExists {
			return c.JSON(http.StatusBadRequest, utils.NewBadResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(institucion))
}

func (h *InstitucionHandlear) Find(c echo.Context) error {
	instituciones, err := h.service.FindInstitucion()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(instituciones))
}
func (h *InstitucionHandlear) FindByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse("ID invalido"))
	}
	institucion, err := h.service.FindInstitucionByID(id)
	if err != nil {
		if err == utils.ErrNothingFind {
			return c.JSON(http.StatusNotFound, utils.NewBadResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(institucion))
}
func (h *InstitucionHandlear) findByUserName(c echo.Context, username string) error {

	institucion, err := h.service.FindInstitucionByUserName(username)
	if err != nil {
		if err == utils.ErrUsuarioNotExists {
			return c.JSON(http.StatusNotFound, utils.NewBadResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(institucion))
}

func (h *InstitucionHandlear) FindMyInfo(c echo.Context) error {
	username, ok := c.Get(authorization.USERNAME_KEY).(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse("username no se pudo leer"))
	}
	return h.findByUserName(c, username)
}

func (h *InstitucionHandlear) FindByUserName(c echo.Context) error {
	username := c.Param("username")
	return h.findByUserName(c, username)
}
