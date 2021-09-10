package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/user0608/kcheck"
	"github.com/user0608/mujeresapi/authorization"
	"github.com/user0608/mujeresapi/models/application"
	"github.com/user0608/mujeresapi/models/authentication"
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
func (u *UsuarioHandlear) LogginUser(c echo.Context) error {
	Binder := &echo.DefaultBinder{}
	Request := utils.RequestLoging{}
	if err := Binder.BindBody(c, &Request); err != nil {
		log.Println("No se pudo parsear los datos", c.Request().URL, " Err:", err.Error())
		return echo.ErrBadRequest
	}
	user, err := u.service.Loggin(Request)
	if err != nil {
		log.Println("Login invalid : ", err.Error())
		if err == utils.ErrNothingFind {
			return c.JSON(http.StatusBadRequest, utils.NewBadResponse("Usuario o password invalidos"))
		}
		return echo.ErrBadRequest
	}

	token, err := authorization.GenerageToken(*user)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, utils.LogginResponse{Usuario: user, Token: token, Code: utils.CODE_OK})
}

func (u *UsuarioHandlear) RegistrarUsuariosGeneral(isApp bool) echo.HandlerFunc {
	return func(c echo.Context) error {
		Binder := &echo.DefaultBinder{}
		user := &authentication.Usuario{}
		if err := Binder.BindBody(c, user); err != nil {
			log.Println("Peticion con estructura de datos incorrecta", err.Error())
			return echo.ErrBadRequest
		}
		var err error
		if isApp {
			err = u.service.RegistrarUsuarioApp(user)
		} else {
			err = u.service.RegistrarUsuario(user)
		}
		if err != nil {
			if err == utils.ErrExistUsername {
				return c.JSON(http.StatusBadRequest, utils.Response{
					Code:    utils.COD_USUARIO_EXISTE,
					Message: "El username ya esta en uso",
				})
			}
			log.Println("Error DB,", err.Error())
			return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
		}
		return c.JSON(http.StatusOK, utils.NewOkResponse(user))
	}
}

func (u *UsuarioHandlear) CrateUpdateUser(isCreate bool) echo.HandlerFunc {
	return func(c echo.Context) error {
		Binder := &echo.DefaultBinder{}
		appuser := &application.AppUser{}
		if err := Binder.BindBody(c, appuser); err != nil {
			return echo.ErrBadRequest
		}
		var err error
		if isCreate {
			appuser.ID = c.Get(authorization.USUARIO_ID_KEY).(int)
			err = u.service.RegistrarDetalleAppUsuario(appuser)
			if err == utils.ErrIdIsNeeded {
				log.Println(err.Error())
				return c.JSON(http.StatusBadRequest, utils.NewBadResponse(err.Error()))
			}
		} else {
			err = u.service.UpdateDetalleAppUsuario(appuser)
			if err == utils.ErrNothingFind {
				log.Println(err.Error())
				return c.JSON(http.StatusBadRequest, utils.NewNoFindResponse("Registro no encontrado"))
			}
		}
		if err != nil {
			log.Println(err.Error())
			return echo.ErrInternalServerError
		}
		return c.JSON(http.StatusOK, utils.NewOkResponse(appuser))
	}
}

func (u *UsuarioHandlear) AllUsersHandler(c echo.Context) error {
	usuarios, err := u.service.AllUsuarios()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(usuarios))
}
func (u *UsuarioHandlear) AllFreeUsuarios(c echo.Context) error {
	usuarios, err := u.service.AllFreeUsuarios()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(usuarios))
}
func (u *UsuarioHandlear) GetUserAppById(c echo.Context) error {
	id := c.Param("usuario_id")
	if err := kcheck.New().Target("num", id).Ok(); err != nil {
		return echo.ErrBadRequest
	}
	usuarioID, _ := strconv.Atoi(id)
	usuario, err := u.service.GetUsuarioDetalleApp(usuarioID)
	if err != nil {
		if err == utils.ErrNothingFind {
			return c.JSON(http.StatusBadRequest, utils.NewNoFindResponse("Registro no encontrado"))
		}
		log.Println(err.Error())
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(usuario))
}
