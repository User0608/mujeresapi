package repository

import "gorm.io/gorm"

type InstitucionRepository struct {
	gdb *gorm.DB
}

func NewInstitucionRepository(db *gorm.DB) *UsuarioRepository {
	return &UsuarioRepository{
		gdb: db,
	}
}
