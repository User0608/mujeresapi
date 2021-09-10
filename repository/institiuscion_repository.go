package repository

import (
	"log"

	"github.com/user0608/mujeresapi/models/authentication"
	"github.com/user0608/mujeresapi/models/control"
	"github.com/user0608/mujeresapi/utils"
	"gorm.io/gorm"
)

type InstitucionRepository struct {
	gdb *gorm.DB
}

func NewInstitucionRepository(g *gorm.DB) *InstitucionRepository {
	return &InstitucionRepository{gdb: g}
}
func (r *InstitucionRepository) ConsultarDisponibilidadDeUsuario(usuarioID int) error {
	response := &UsuarioConsult{}
	if err := consultarUsuario(r.gdb, response, usuarioID); err != nil {
		log.Println("Error:InstitucionRepository.ConsultarDisponibilidadDeUsuario:", err.Error())
		return err
	}
	if response.NotExistUsuario() {
		return utils.ErrUsuarioNotExists
	}
	if !response.OK() {
		return utils.ErrUsuarioNoDisponible
	}
	return nil
}
func (r *InstitucionRepository) Create(i *control.Institucion) error {
	if err := r.ConsultarDisponibilidadDeUsuario(i.UsuarioID); err != nil {
		log.Println("Error:InstitucionRepository.Create:", err.Error())
		return err
	}
	return r.gdb.Create(i).Error
}
func (r *InstitucionRepository) Update(i *control.Institucion) error {
	oldEntity := &control.Institucion{}
	if result := r.gdb.Limit(1).Find(oldEntity, i.ID); result.Error != nil {
		log.Println("Error-1:InstitucionRepository.Update:", result.Error.Error())
		return utils.ErrDataBaseError
	} else {
		if result.RowsAffected == 0 {
			return utils.ErrNothingFind
		}
		if oldEntity.UsuarioID != i.UsuarioID {
			if err := r.ConsultarDisponibilidadDeUsuario(i.UsuarioID); err != nil {
				return err
			}
		}
	}
	if result := r.gdb.Session(&gorm.Session{FullSaveAssociations: true}).Updates(i); result.Error != nil {
		log.Println("Error-2:InstitucionRepository.Update:", result.Error.Error())
		return utils.ErrDataBaseError
	}
	return nil
}

func (r *InstitucionRepository) Find() ([]control.Institucion, error) {
	instituciones := []control.Institucion{}
	if err := r.gdb.Preload("Direccion").Find(&instituciones).Error; err != nil {
		log.Println("Error:InstitucionRepository.Find:", err.Error())
		return nil, err
	}
	return instituciones, nil
}
func (r *InstitucionRepository) FindByID(institucionID int) (*control.Institucion, error) {
	institucion := &control.Institucion{}
	if result := r.gdb.Preload("Direccion").Find(institucion, institucionID); result.Error != nil {
		log.Println("Error:InstitucionRepository.FindByID:", result.Error.Error())
		return nil, utils.ErrDataBaseError
	} else {
		if result.RowsAffected == 0 {
			return nil, utils.ErrNothingFind
		}
		return institucion, nil
	}
}
func (r *InstitucionRepository) FindByUserName(username string) (*control.Institucion, error) {
	usuario := &authentication.Usuario{}
	insitucion := &control.Institucion{}
	if result := r.gdb.Limit(1).Find(usuario, "username = ?", username); result.Error != nil {
		log.Println("Error 1 :InstitucionRepository.FindByUserName:", result.Error.Error())
		return nil, utils.ErrDataBaseError
	} else {
		if result.RowsAffected == 0 {
			return nil, utils.ErrUsuarioNotExists
		}
	}
	if result := r.gdb.Preload("Direccion").Limit(1).Find(insitucion, "usuario_id = ?", usuario.ID); result.Error != nil {
		log.Println("Error 2 :InstitucionRepository.FindByUserName:", result.Error.Error())
		return nil, utils.ErrDataBaseError
	} else {
		if result.RowsAffected == 0 {
			return nil, utils.ErrUsuarioNotExists
		}
	}
	return insitucion, nil
}
