package repository

import (
	"fmt"
	"log"

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
	if err := r.gdb.Create(c).Error; err != nil {
		return err
	}
	c.Persona = nil
	c.Usuario = nil
	return nil
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
func (r *ColaboradorRepository) AllColaboradores() ([]control.Colaborador, error) {
	colaboradores := []control.Colaborador{}
	if res := r.gdb.Preload("Persona").Find(&colaboradores); res.Error != nil {
		log.Println("Error-0: ColaboradorRepository.AllColaboradores:", res.Error.Error())
		return nil, utils.ErrDataBaseError
	}
	return colaboradores, nil
}
func (r *ColaboradorRepository) FindByUsuarioID(usuarioID int) (*control.Colaborador, error) {
	colaborador := &control.Colaborador{}
	fmt.Println(usuarioID)
	if res := r.gdb.Preload("Persona").Limit(1).Find(colaborador, "usuario_id=?", usuarioID); res.Error != nil {
		log.Println("Error-1: EfectivoRepository.FindAllByUsuarioID", res.Error.Error())
		return nil, utils.ErrDataBaseError
	}
	return colaborador, nil
}
