package services

import (
	"github.com/user0608/mujeresapi/models/auxiliares"
	"github.com/user0608/mujeresapi/models/control"
	"github.com/user0608/mujeresapi/storage"
)

type EfectivoService struct {
	storage storage.EfectivoStorage
}

func NewEfectivoService(s storage.EfectivoStorage) *EfectivoService {
	return &EfectivoService{storage: s}
}

func (s *EfectivoService) Create(ae *auxiliares.Efectivo) (*control.Efectivo, error) {
	if err := ae.GetPersona().Validate(); err != nil {
		return nil, err
	}
	return s.storage.Create(ae)
}
func (s *EfectivoService) GetAll() ([]control.Efectivo, error) {
	return s.storage.FindAll()
}
func (s *EfectivoService) GetAllByInstitucionID(institucionID int) ([]control.Efectivo, error) {
	return s.storage.FindAllByInstitucionID(institucionID)
}
func (s *EfectivoService) GetAllByUsuarioID(usuarioID int) ([]control.Efectivo, error) {
	return s.storage.FindAllByUsuarioID(usuarioID)
}
