package repository

import (
	"github.com/user0608/mujeresapi/models/authentication"
	"gorm.io/gorm"
)

type RolesRepository struct {
	gdb *gorm.DB
}

func NewRolesRepository(g *gorm.DB) *RolesRepository {
	return &RolesRepository{gdb: g}
}
func (r *RolesRepository) GetAll() ([]authentication.Roles, error) {
	roles := []authentication.Roles{}
	return roles, r.gdb.Find(&roles, []int{3, 4}).Error
}
