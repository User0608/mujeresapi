package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/user0608/mujeresapi/models/control"
	"github.com/user0608/mujeresapi/services"
	"github.com/user0608/mujeresapi/utils"
)

type PersonaHandlear struct {
	service *services.PersonaService
	parser  *echo.DefaultBinder
}

func NewPersonaHandler(s *services.PersonaService) *PersonaHandlear {
	return &PersonaHandlear{
		service: s,
		parser:  &echo.DefaultBinder{},
	}
}
func (h *PersonaHandlear) FindAll(c echo.Context) error {
	p, err := h.service.FindAll()
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(p))
}

func (h *PersonaHandlear) FindByID(c echo.Context) error {
	personaID, err := strconv.Atoi(c.Param("persona_id"))
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse("ID incorrecto"))
	}
	p, err := h.service.FindByID(personaID)
	if err != nil {
		if err == utils.ErrIDInvalido {
			log.Println(err.Error())
			return c.JSON(http.StatusBadRequest, utils.NewBadResponse("ID incorrecto"))
		}
		if err == utils.ErrNothingFind {
			log.Println(err.Error())
			return c.JSON(http.StatusNotFound, utils.NewBadResponse("Recuso no encontrado!"))
		}
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(p))
}

func (h *PersonaHandlear) Register(c echo.Context) error {
	p := &control.Persona{}
	if err := h.parser.BindBody(c, p); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse("Estructura de datos incorrecta"))
	}
	if err := h.service.Register(p); err != nil {
		log.Println(err.Error())
		if err == utils.ErrInvalidData {
			return c.JSON(http.StatusBadRequest, utils.NewBadResponse("Datos incorrectos"))
		}
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(p))
}

func (h *PersonaHandlear) Update(c echo.Context) error {
	p := &control.Persona{}
	if err := h.parser.BindBody(c, p); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse("Estructura de datos incorrecta"))
	}
	if err := h.service.Update(p); err != nil {
		log.Println(err.Error())
		if err == utils.ErrNothingFind {
			return c.JSON(http.StatusNotFound, utils.NewBadResponse("Registro no encontrado"))
		}
		if err == utils.ErrInvalidData {
			return c.JSON(http.StatusBadRequest, utils.NewBadResponse("Datos incorrectos"))
		}
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(p))
}
