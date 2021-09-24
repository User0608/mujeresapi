package storage

import "github.com/user0608/mujeresapi/models/control"

type AsignacionStorage interface {
	Create(a control.Asignacion) (*control.AsignacionTable, error)
	Update(id int, a control.Asignacion) (*control.AsignacionTable, error)
	Delete(id int) error
	GetAllByEfectivoID(id int) ([]control.AsignacionTable, error)
	GetByID(id int) (*control.AsignacionTable, error)
	GetAllByNotificatioID(id int) (*control.AsignacionTable, error)
}
