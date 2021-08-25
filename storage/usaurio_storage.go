package storage

import (
	"github.com/user0608/mujeresapi/models/application"
	"github.com/user0608/mujeresapi/models/authentication"
)

type UsuarioStorage interface {
	GetAllUsuarios() ([]authentication.Usuario, error)
	LoginUsuario(username, password string) (*authentication.Usuario, error)
	GetUser(username string) (*authentication.Usuario, bool, error)
	RegistrarUsuario(user *authentication.Usuario) error
	RegistrarDetalleAppUsuario(appuser *application.AppUser) error
	UpdateDetalleAppUsuario(appuser *application.AppUser) error
	GetUsuarioDetalleApp(userid int) (*application.AppUser, error)
}
