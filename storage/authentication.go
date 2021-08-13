package storage

import "github.com/user0608/mujeresapi/models/authentication"

type UsuarioStorage interface {
	GetAllUsuarios() ([]authentication.Usuario, error)
}
