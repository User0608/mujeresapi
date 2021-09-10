package services

import (
	"github.com/user0608/mujeresapi/models/control"
	"github.com/user0608/mujeresapi/storage"
	"github.com/user0608/mujeresapi/utils"
)

type PersonaService struct {
	storage storage.PersonaStorage
}

func NewPersonaService(s storage.PersonaStorage) *PersonaService {
	return &PersonaService{storage: s}
}
func (s *PersonaService) FindAll() ([]control.Persona, error) {
	return s.storage.FindAll()
}

func (s *PersonaService) FindByID(id int) (*control.Persona, error) {
	if id == 0 {
		return nil, utils.ErrIDInvalido
	}
	return s.storage.FindByID(id)
}

func (s *PersonaService) Register(p *control.Persona) error {
	if err := p.Validate(); err != nil {
		return utils.ErrInvalidData
	}
	return s.storage.Register(p)
}

func (s *PersonaService) Update(p *control.Persona) error {
	if err := p.Validate(); err != nil {
		return utils.ErrInvalidData
	}
	if p.ID <= 0 {
		return utils.ErrIDInvalido
	}
	return s.storage.Update(p)
}
