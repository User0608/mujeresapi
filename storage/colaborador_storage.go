package storage

import "github.com/user0608/mujeresapi/models/control"

type ColaboradorStorage interface {
	AsociarPersonaUsuario(c *control.Colaborador) error
	FindColaborador(colaboradorID int) (*control.Colaborador, error)
}
