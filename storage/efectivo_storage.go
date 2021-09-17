package storage

import (
	"github.com/user0608/mujeresapi/models/auxiliares"
	"github.com/user0608/mujeresapi/models/control"
)

type EfectivoStorage interface {
	Create(ae *auxiliares.Efectivo) (*control.Efectivo, error)
	FindAll() ([]control.Efectivo, error)
	FindAllByInstitucionID(institucionID int) ([]control.Efectivo, error)
	FindAllByUsuarioID(usuarioID int) ([]control.Efectivo, error)
}
