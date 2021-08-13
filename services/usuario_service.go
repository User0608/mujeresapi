package services

import (
	"github.com/user0608/mujeresapi/models/authentication"
	"github.com/user0608/mujeresapi/storage"
)

type UsuarioService struct {
	storage storage.UsuarioStorage
}

func NewUsuarioService(storage storage.UsuarioStorage) *UsuarioService {
	return &UsuarioService{storage}
}

func (u *UsuarioService) AllUsuarios() ([]authentication.Usuario, error) {
	return u.storage.GetAllUsuarios()
}
