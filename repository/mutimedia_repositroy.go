package repository

import (
	"github.com/user0608/mujeresapi/models/application"
	"gorm.io/gorm"
)

type MultimediaRepository struct {
	gdb *gorm.DB
}

func NewMultimediaRepository(g *gorm.DB) *MultimediaRepository {
	return &MultimediaRepository{gdb: g}
}

func (r *MultimediaRepository) RegistrarMultimedia(m []application.Multimedia) error {
	return r.gdb.Create(m).Error
}

func (r *MultimediaRepository) VerificarAlertaID(alertaID int) error {
	return existRegister(&application.Alerta{}, alertaID, r.gdb)
}
