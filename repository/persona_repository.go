package repository

import (
	"github.com/user0608/mujeresapi/models/control"
	"github.com/user0608/mujeresapi/utils"
	"gorm.io/gorm"
)

type PersonaRepository struct {
	gdb *gorm.DB
}

func NewPersonaRepository(db *gorm.DB) *PersonaRepository {
	return &PersonaRepository{
		gdb: db,
	}
}
func (r *PersonaRepository) FindAll() ([]control.Persona, error) {
	var ps = []control.Persona{}
	return ps, r.gdb.Preload("Direccion").Find(&ps).Error
}
func (r *PersonaRepository) FindByID(id int) (*control.Persona, error) {
	var p = &control.Persona{}
	result := r.gdb.Preload("Direccion").Find(p, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, utils.ErrNothingFind
	}
	return p, nil
}
func (r *PersonaRepository) Register(p *control.Persona) error {
	return r.gdb.Create(p).Error
}
func (r *PersonaRepository) Update(p *control.Persona) error {
	if err := existRegister(r.gdb, &control.Persona{}, p.ID); err != nil {
		return err
	}
	return r.gdb.Session(&gorm.Session{FullSaveAssociations: true}).Updates(p).Error
}
