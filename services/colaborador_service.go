package services

import (
	"github.com/user0608/mujeresapi/models/control"
	"github.com/user0608/mujeresapi/storage"
	"github.com/user0608/mujeresapi/utils"
)

type ColaboradorService struct {
	storage storage.ColaboradorStorage
}

func NewColaboradorService(s storage.ColaboradorStorage) *ColaboradorService {
	return &ColaboradorService{
		storage: s,
	}
}
func (s *ColaboradorService) Asociar(c *control.Colaborador) error {
	if c == nil {
		return utils.ErrInvalidData
	}
	if c.Persona != nil || c.Usuario != nil || c.PersonaID == 0 || c.UsuarioID == 0 {
		return utils.ErrInvalidData
	}
	return s.storage.AsociarPersonaUsuario(c)
}

func (s *ColaboradorService) FindByID(colaboradorID int) (*control.Colaborador, error) {
	if colaboradorID == 0 {
		return nil, utils.ErrInvalidData
	}
	return s.storage.FindColaborador(colaboradorID)
}

func (s *ColaboradorService) AllColaboradores() ([]control.Colaborador, error) {
	return s.storage.AllColaboradores()
}
