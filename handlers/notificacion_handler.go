package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/user0608/mujeresapi/authorization"
	"github.com/user0608/mujeresapi/authorization/roles"
	"github.com/user0608/mujeresapi/models/control"
	"github.com/user0608/mujeresapi/services"
	"github.com/user0608/mujeresapi/utils"
)

type NotificacionHandler struct {
	binder  echo.DefaultBinder
	service *services.NotificacionService
}

func NewNotificacionHandler(s *services.NotificacionService) *NotificacionHandler {
	return &NotificacionHandler{
		service: s,
		binder:  echo.DefaultBinder{},
	}
}

func (h *NotificacionHandler) Create(c echo.Context) error {
	n := &control.Notificacion{}
	if err := h.binder.BindBody(c, n); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse("Estructura JSON Invalida!"))
	}
	if err := h.service.Create(n); err != nil {
		if err == utils.ErrDataBaseError {
			return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
		}
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(n))
}

func (h *NotificacionHandler) Update(c echo.Context) error {
	n := &control.Notificacion{}
	if err := h.binder.BindBody(c, n); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse("Estructura JSON Invalida!"))
	}
	if err := h.service.Update(n); err != nil {
		if err == utils.ErrDataBaseError {
			return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
		}
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(n))
}
func (h *NotificacionHandler) FindAll(c echo.Context) error {
	if notificaciones, err := h.service.FindAll(); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(err.Error()))
	} else {
		return c.JSON(http.StatusOK, utils.NewOkResponse(notificaciones))
	}
}

func (h *NotificacionHandler) FindByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	var notificacion *control.Notificacion
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewBadResponse("No se encontro el registro"))
	}
	userid, ok := c.Get(authorization.USUARIO_ID_KEY).(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	if rls, ok := c.Get(authorization.ROLES_KEY).([]string); ok {
		var ins bool
		var err error
		for _, r := range rls {
			if r == roles.INSTITUCION_ROLE {
				ins = true
				break
			}
		}
		if ins {
			notificacion, err = h.service.FindByIDForInstitucion(id, userid)
		} else {
			notificacion, err = h.service.FindByID(id)
		}
		if err != nil {
			if err == utils.ErrDataBaseError {
				return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
			} else if err == utils.ErrNothingFind {
				return c.JSON(http.StatusNotFound, utils.NewBadResponse("No se encontro el registro"))
			} else {
				return c.JSON(http.StatusBadRequest, utils.NewBadResponse(err.Error()))
			}
		}
	} else {
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(notificacion))
}

func (h *NotificacionHandler) FindByInstitucionID(c echo.Context) error {
	institucionID, err := strconv.Atoi(c.Param("institucion_id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewBadResponse("No se encontro el recurso"))
	}
	if notificaciones, err := h.service.FindInstitucionID(institucionID); err != nil {
		if err == utils.ErrDataBaseError {
			return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
		}
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse(err.Error()))
	} else {
		return c.JSON(http.StatusOK, utils.NewOkResponse(notificaciones))
	}
}

// FindNotificationByInstitucion devuelve, todas las notificaciones, pero este
// handler tiene que ejecutarse con rol insitucion.

func (h *NotificacionHandler) FindNotificationByInstitucion(c echo.Context) error {
	usuarioID, ok := c.Get(authorization.USUARIO_ID_KEY).(int)
	if !ok {
		log.Println("Error-0: NotificacionHandler.FindNotificationByInstitucion: No se pudo encontrar el usuario ID ")
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	if notificaciones, err := h.service.FindAllForInstitucion(usuarioID); err != nil {
		if err == utils.ErrDataBaseError {
			return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
		}
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse(err.Error()))
	} else {
		return c.JSON(http.StatusOK, utils.NewOkResponse(notificaciones))
	}

}
func (h *NotificacionHandler) FindByColaboradorID(c echo.Context) error {
	colaboradorID, err := strconv.Atoi(c.Param("colaborador_id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewBadResponse("Registro no encontrado"))

	}
	if notificaciones, err := h.service.FindColaboradorID(colaboradorID); err != nil {
		if err == utils.ErrDataBaseError {
			return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
		}
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse(err.Error()))
	} else {
		return c.JSON(http.StatusOK, utils.NewOkResponse(notificaciones))
	}
}
