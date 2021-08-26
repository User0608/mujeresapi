package services

import (
	"strconv"

	"github.com/user0608/mujeresapi/models/application"
	"github.com/user0608/mujeresapi/storage"
	"github.com/user0608/mujeresapi/utils"
)

type AlertaService struct {
	storage storage.AlertaStorage
}

func NewAlertaService(storage storage.AlertaStorage) *AlertaService {
	return &AlertaService{storage}
}
func (a *AlertaService) GatAll() ([]application.Alerta, error) {
	return a.storage.FetchAll()
}
func (a *AlertaService) GatAllWithUserID(userid int) ([]application.Alerta, error) {
	return a.storage.FetchAllByUserID(userid)
}
func (a *AlertaService) GatByID(alertaID string) (*application.Alerta, error) {
	id, _ := strconv.Atoi(alertaID)
	return a.storage.FetchByID(id)
}
func (a *AlertaService) Create(alerta *application.Alerta) error {
	if alerta == nil {
		return utils.ErrNullEntity
	}
	alerta.Multimedias = make([]application.Multimedia, 0)
	if err := utils.ValidarDecimal(alerta.Longitude); err != nil {
		return utils.ErrInvalidData
	}
	if err := utils.ValidarDecimal(alerta.Latitude); err != nil {
		return utils.ErrInvalidData
	}
	return a.storage.Create(alerta)
}
