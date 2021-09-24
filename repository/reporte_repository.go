package repository

import (
	"log"
	"time"

	"github.com/user0608/mujeresapi/models/control"
	"github.com/user0608/mujeresapi/utils"
	"gorm.io/gorm"
)

type ReporteRepository struct {
	gorm *gorm.DB
}

func NewReporteRepository(g *gorm.DB) *ReporteRepository {
	return &ReporteRepository{gorm: g}
}

func (rr *ReporteRepository) Create(r control.Reporte) (*control.ReporteTable, error) {
	reporte := &control.ReporteTable{
		Reporte:    r,
		Created_At: time.Now(),
		Updated_At: time.Now(),
	}
	if err := rr.gorm.Create(reporte).Error; err != nil {
		log.Println("Error-1:ReporteRepository.Create:", err.Error())
		return nil, utils.ErrDataBaseError
	}
	return reporte, nil
}
func (rr *ReporteRepository) Update(id int, r control.Reporte) (*control.ReporteTable, error) {
	reporte := &control.ReporteTable{
		ID:         id,
		Reporte:    r,
		Updated_At: time.Now(),
	}
	if err := rr.gorm.Omit("created_at").Updates(reporte).Error; err != nil {
		log.Println("Error-1:ReporteRepository.Update:", err.Error())
		return nil, utils.ErrDataBaseError
	}
	return reporte, nil
}
func (rr *ReporteRepository) Delete(id int) error {
	if res := rr.gorm.Delete(&control.ReporteTable{ID: id}); res.Error != nil {
		log.Println("Error-1:ReporteRepository.Update:", res.Error.Error())
		return utils.ErrDataBaseError
	} else {
		if res.RowsAffected == 0 {
			return utils.ErrNothingFind
		}
	}
	return nil
}
func (rr *ReporteRepository) FindAll() ([]control.ReporteTable, error) {
	reportes := []control.ReporteTable{}
	if res := rr.gorm.Find(&reportes); res.Error != nil {
		log.Println("Error-1:ReporteRepository.Update:", res.Error.Error())
		return nil, utils.ErrDataBaseError
	}
	return reportes, nil
}
func (rr *ReporteRepository) FindByID(id int) (*control.ReporteTable, error) {
	reporte := &control.ReporteTable{}
	if res := rr.gorm.Limit(1).Preload("Notificacion").Find(reporte, id); res.Error != nil {
		log.Println("Error-1:ReporteRepository.Update:", res.Error.Error())
		return nil, utils.ErrDataBaseError
	} else {
		if res.RowsAffected == 0 {
			return nil, utils.ErrNothingFind
		}
	}
	return reporte, nil
}
func (rr *ReporteRepository) FindByNotificacionID(notificacionID int) (*control.ReporteTable, error) {
	reporte := &control.ReporteTable{}
	if res := rr.gorm.Limit(1).Find(reporte, "notificacion_id=?", notificacionID); res.Error != nil {
		log.Println("Error-1:ReporteRepository.Update:", res.Error.Error())
		return nil, utils.ErrDataBaseError
	} else {
		if res.RowsAffected == 0 {
			return nil, utils.ErrNothingFind
		}
	}
	return reporte, nil
}
