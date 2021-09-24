package repository

import (
	"log"
	"time"

	"github.com/user0608/mujeresapi/models/control"
	"github.com/user0608/mujeresapi/utils"
	"gorm.io/gorm"
)

type AsignarEfectivoRepository struct {
	gorm *gorm.DB
}

func NewAsignarEfectivoRepositoryRepository(g *gorm.DB) *AsignarEfectivoRepository {
	return &AsignarEfectivoRepository{gorm: g}
}

func (r *AsignarEfectivoRepository) Create(a control.Asignacion) (*control.AsignacionTable, error) {
	asignacion := &control.AsignacionTable{
		Asignacion: a,
		CreatedAt:  time.Now(),
	}
	if res := r.gorm.Omit("created_at").Save(asignacion); res.Error != nil {
		log.Println("Error-1:AsignarEfectivoRepository.Create", res.Error.Error())
		return nil, utils.ErrInvalidData
	}
	return asignacion, nil
}

func (r *AsignarEfectivoRepository) Update(id int, a control.Asignacion) (*control.AsignacionTable, error) {
	if err := CheckTablesRecord(r.gorm, ChTables{{"asignacion", id}}); err != nil {
		return nil, err
	}
	asignacion := &control.AsignacionTable{
		ID:         id,
		Asignacion: a,
	}
	if res := r.gorm.Omit("created_at").Save(asignacion); res.Error != nil {
		log.Println("Error-1:AsignarEfectivoRepository.Update:", res.Error.Error())
		return nil, utils.ErrDataBaseError
	}
	return asignacion, nil
}

func (r *AsignarEfectivoRepository) Delete(id int) error {
	if err := CheckTablesRecord(r.gorm, ChTables{{"asignacion", id}}); err != nil {
		return err
	}
	if err := r.gorm.Delete(&control.AsignacionTable{ID: id}).Error; err != nil {
		log.Println("Error-1:AsignarEfectivoRepository.Delete:", err.Error())
		return utils.ErrDataBaseError
	}
	return nil
}

func (r *AsignarEfectivoRepository) GetAllByEfectivoID(id int) ([]control.AsignacionTable, error) {
	asignacion := []control.AsignacionTable{}
	if err := r.gorm.Preload("Notificacion").Find(&asignacion, "efectivo_id =?", id).Error; err != nil {
		log.Println("Error-1:AsignarEfectivoRepository.GetAllByEfectivoID:", err.Error())
		return nil, utils.ErrDataBaseError
	}
	return asignacion, nil
}
func (r *AsignarEfectivoRepository) GetByID(id int) (*control.AsignacionTable, error) {
	asignacion := &control.AsignacionTable{}
	if err := r.gorm.Preload("Notificacion").Limit(1).Find(asignacion, id).Error; err != nil {
		log.Println("Error-1:AsignarEfectivoRepository.GetAllByID:", err.Error())
		return nil, utils.ErrDataBaseError
	}
	return asignacion, nil
}
func (r *AsignarEfectivoRepository) GetAllByNotificatioID(id int) (*control.AsignacionTable, error) {
	asignacion := &control.AsignacionTable{}
	if err := r.gorm.Preload("Notificacion").Limit(1).Find(asignacion, "notificacion_id=?", id).Error; err != nil {
		log.Println("Error-1:AsignarEfectivoRepository.GetAllByID:", err.Error())
		return nil, utils.ErrDataBaseError
	}
	return asignacion, nil
}
