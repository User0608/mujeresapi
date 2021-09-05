package repository

import (
	"github.com/user0608/mujeresapi/models/application"
	"github.com/user0608/mujeresapi/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AlertaRepository struct {
	gdb *gorm.DB
}

func NewAlertaRepository(db *gorm.DB) *AlertaRepository {
	return &AlertaRepository{
		gdb: db,
	}
}

func (a *AlertaRepository) Create(alerta *application.Alerta) error {
	return a.gdb.Omit(clause.Associations).Create(alerta).Error
}
func (a *AlertaRepository) FetchAllByUserAppID(userID int) ([]application.Alerta, error) {
	alertas := []application.Alerta{}
	if err := a.gdb.Preload("Multimedias").Order("created_at desc").Find(&alertas, "usuario_id = ?", userID).Error; err != nil {
		return []application.Alerta{}, err
	}
	return alertas, nil
}
func (a *AlertaRepository) FetchAllByUserID(userid int) ([]application.Alerta, error) {
	alertas := []application.Alerta{}
	if err := a.gdb.Preload("Multimedias").Where("usuario_id = ?", userid).Order("created_at desc").Find(&alertas).Error; err != nil {
		return []application.Alerta{}, err
	}
	return alertas, nil
}
func (a *AlertaRepository) FetchAll() ([]application.Alerta, error) {
	alertas := []application.Alerta{}
	if err := a.gdb.Preload("Multimedias").Find(&alertas).Error; err != nil {
		return []application.Alerta{}, err
	}
	return alertas, nil
}

func (a *AlertaRepository) FetchByID(alertaID int) (*application.Alerta, error) {
	alerta := &application.Alerta{}
	result := a.gdb.Preload("Multimedias").First(alerta, alertaID)
	if err := result.Error; err != nil {
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, utils.ErrNothingFind
	}
	return alerta, nil
}
