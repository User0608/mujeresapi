package services

import (
	"github.com/user0608/kcheck"
	"github.com/user0608/mujeresapi/models/application"
	"github.com/user0608/mujeresapi/models/authentication"
	"github.com/user0608/mujeresapi/storage"
	"github.com/user0608/mujeresapi/utils"
)

type UsuarioService struct {
	storage storage.UsuarioStorage
}

func NewUsuarioService(storage storage.UsuarioStorage) *UsuarioService {
	return &UsuarioService{storage}
}
func (u *UsuarioService) Loggin(r utils.RequestLoging) (*authentication.Usuario, error) {
	if r.Password == "" || r.Username == "" {
		return nil, utils.ErrNoUsernameOrPassword
	}
	user, err := u.storage.LoginUsuario(r.Username, r.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (u *UsuarioService) AllUsuarios() ([]authentication.Usuario, error) {
	return u.storage.GetAllUsuarios()
}

func (u *UsuarioService) RegistrarUsuarioApp(user *authentication.Usuario) error {
	if user == nil {
		return utils.ErrNullEntity
	}
	if len(user.Roles) > 0 {
		return utils.ErrInvalidData
	}
	user.Roles = []*authentication.Roles{{ID: 2}}
	return u.RegistrarUsuario(user)
}

func (u *UsuarioService) RegistrarUsuario(user *authentication.Usuario) error {
	if user == nil {
		return utils.ErrNullEntity
	}
	if len(user.Roles) == 0 {
		user.Roles = []*authentication.Roles{{ID: 3}}
	}
	user.Estado = true
	check := kcheck.New()
	if err := check.Target("min=8 no-spaces", user.Username, user.Password).Ok(); err != nil {
		return utils.ErrInvalidData
	}
	return u.storage.RegistrarUsuario(user)
}

func (u *UsuarioService) ValidarUsuario(appuser *application.AppUser) error {
	if appuser == nil {
		return utils.ErrNullEntity
	}
	check := kcheck.New()
	valid := check.Target("basic", appuser.Nombre, appuser.ApellidoPaterno, appuser.ApellidoMaterno)
	if err := valid.Ok(); err != nil {
		return utils.ErrInvalidData
	}
	if err := check.Target("num len=8", appuser.Dni).Ok(); err != nil {
		return utils.ErrInvalidData
	}
	if err := check.Target("num max=20", appuser.Telefono).Ok(); err != nil {
		return utils.ErrInvalidData
	}
	return nil
}
func (u *UsuarioService) RegistrarDetalleAppUsuario(appuser *application.AppUser) error {
	if err := u.ValidarUsuario(appuser); err != nil {
		return err
	}
	return u.storage.RegistrarDetalleAppUsuario(appuser)
}

func (u *UsuarioService) UpdateDetalleAppUsuario(appuser *application.AppUser) error {
	if err := u.ValidarUsuario(appuser); err != nil {
		return err
	}
	return u.storage.UpdateDetalleAppUsuario(appuser)
}

func (u *UsuarioService) GetUsuarioDetalleApp(userid int) (*application.AppUser, error) {
	return u.storage.GetUsuarioDetalleApp(userid)
}
