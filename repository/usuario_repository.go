package repository

import (
	"github.com/user0608/mujeresapi/models/authentication"
	"gorm.io/gorm"
)

type UsuarioRepository struct {
	gdb *gorm.DB
}

func NewUsuarioRepository(db *gorm.DB) *UsuarioRepository {
	return &UsuarioRepository{
		gdb: db,
	}
}

func (u *UsuarioRepository) GetAllUsuarios() ([]authentication.Usuario, error) {
	usuarios := []authentication.Usuario{}
	if err := u.gdb.Preload("Roles").Omit("Password").Find(&usuarios).Error; err != nil {
		return nil, err
	}
	return usuarios, nil
}
