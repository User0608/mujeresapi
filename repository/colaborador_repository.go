package repository

import (
	"github.com/user0608/mujeresapi/models/application"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ColaboradorRepository struct {
	gdb *gorm.DB
}

func NewColaboradorRepository(db *gorm.DB) *ColaboradorRepository {
	return &ColaboradorRepository{
		gdb: db,
	}
}

func (a *ColaboradorRepository) Create(alerta *application.Alerta) error {
	return a.gdb.Omit(clause.Associations).Create(alerta).Error
}
