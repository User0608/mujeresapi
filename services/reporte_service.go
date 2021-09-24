package services

import (
	"errors"

	"github.com/user0608/mujeresapi/models/control"
	"github.com/user0608/mujeresapi/storage"
)

type ReporteService struct {
	storage storage.ReporteStorage
}

func NewReporteService(s storage.ReporteStorage) *ReporteService {
	return &ReporteService{storage: s}
}
func (s *ReporteService) Create(r control.Reporte) (*control.ReporteTable, error) {
	if err := r.Valid(); err != nil {
		return nil, err
	}
	return s.storage.Create(r)
}
func (s *ReporteService) Update(id int, r control.Reporte) (*control.ReporteTable, error) {
	if id == 0 {
		return nil, errors.New("/<ID> no pude ser 0")
	}
	if err := r.Valid(); err != nil {
		return nil, err
	}
	return s.storage.Update(id, r)
}
func (s *ReporteService) Delete(id int) error {
	if id == 0 {
		return errors.New("/<ID> no pude ser 0")
	}
	return s.storage.Delete(id)
}
func (s *ReporteService) FindAll() ([]control.ReporteTable, error) {
	return s.storage.FindAll()
}
func (s *ReporteService) FindByID(id int) (*control.ReporteTable, error) {
	if id == 0 {
		return nil, errors.New("/<ID> no pude ser 0")
	}
	return s.storage.FindByID(id)
}
func (s *ReporteService) FindByNotificacionID(notificacionID int) (*control.ReporteTable, error) {
	if notificacionID == 0 {
		return nil, errors.New("/<ID> no pude ser 0")
	}
	return s.storage.FindByNotificacionID(notificacionID)
}
