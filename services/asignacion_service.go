package services

import (
	"github.com/user0608/mujeresapi/models/control"
	"github.com/user0608/mujeresapi/storage"
	"github.com/user0608/mujeresapi/utils"
)

type AsignacionService struct {
	storage storage.AsignacionStorage
}

func NewAsignacionService(s storage.AsignacionStorage) *AsignacionService {
	return &AsignacionService{storage: s}
}

func (s *AsignacionService) Create(a control.Asignacion) (*control.AsignacionTable, error) {
	if a.EfectivoID == 0 || a.NotificacionID == 0 {
		return nil, utils.ErrInvalidData
	}
	return s.storage.Create(a)
}
func (s *AsignacionService) Update(id int, a control.Asignacion) (*control.AsignacionTable, error) {
	if a.EfectivoID == 0 || a.NotificacionID == 0 || id == 0 {
		return nil, utils.ErrInvalidData
	}
	return s.storage.Update(id, a)
}
func (s *AsignacionService) Delete(id int) error {
	if id == 0 {
		return utils.ErrInvalidData
	}
	return s.storage.Delete(id)
}
func (s *AsignacionService) GetAllByEfectivoID(id int) ([]control.AsignacionTable, error) {
	if id == 0 {
		return nil, utils.ErrInvalidData
	}
	return s.storage.GetAllByEfectivoID(id)
}
func (s *AsignacionService) GetByID(id int) (*control.AsignacionTable, error) {
	if id == 0 {
		return nil, utils.ErrInvalidData
	}
	return s.storage.GetByID(id)
}
func (s *AsignacionService) GetAllByNotificatioID(id int) (*control.AsignacionTable, error) {
	if id == 0 {
		return nil, utils.ErrInvalidData
	}
	return s.storage.GetAllByNotificatioID(id)
}
