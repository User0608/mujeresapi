package repository

import (
	"github.com/user0608/mujeresapi/models/control"
	"github.com/user0608/mujeresapi/utils"
	"gorm.io/gorm"
)

type ColaboradorRepository struct {
	gdb *gorm.DB
}

func NewColaboradorRepository(db *gorm.DB) *ColaboradorRepository {
	return &ColaboradorRepository{
		gdb: db,
	}
}
func (r *ColaboradorRepository) AsociarPersonaUsuario(c *control.Colaborador) error {
	res := UsuarioConsult{}
	if err := r.gdb.Raw("select usuario_is_free(?)", c.UsuarioID).Scan(&res).Error; err != nil {
		return err
	}
	if !res.OK() {
		return utils.ErrIDInvalido
	}
	return r.gdb.Create(c).Error
}
func (r *ColaboradorRepository) FindColaborador(colaboradorID int) (*control.Colaborador, error) {
	colaborador := &control.Colaborador{}
	if result := r.gdb.Preload("Persona").Preload("Usuario").Find(colaborador, colaboradorID); result.Error != nil {
		return nil, result.Error
	} else {
		if result.RowsAffected == 0 {
			return nil, utils.ErrNothingFind
		}
		if colaborador.Usuario != nil {
			colaborador.Usuario.Password = ""
		}
		return colaborador, nil
	}
}
