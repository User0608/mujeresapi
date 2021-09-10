package storage

import "github.com/user0608/mujeresapi/models/control"

type PersonaStorage interface {
	FindAll() ([]control.Persona, error)
	FindByID(id int) (*control.Persona, error)
	Register(p *control.Persona) error
	Update(p *control.Persona) error
}
