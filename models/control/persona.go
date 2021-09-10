package control

import (
	"github.com/user0608/kcheck"
	"github.com/user0608/mujeresapi/models/application"
)

type Persona struct {
	ID              int                    `json:"id"`
	Nombre          string                 `json:"nombre"`
	ApellidoMaterno string                 `json:"apellido_materno"`
	ApellidoPaterno string                 `json:"apellido_paterno"`
	Telefono        string                 `json:"telefono"`
	Dni             string                 `json:"dni"`
	DireccionID     int                    `json:"direccion_id,omitempty"`
	Direccion       *application.Direccion `json:"direccion,omitempty"`
}

func (p *Persona) Validate() error {
	chk := kcheck.New()
	if err := chk.Target("min=2 basic", p.Nombre, p.ApellidoPaterno, p.ApellidoMaterno).Ok(); err != nil {
		return err
	}
	if err := chk.Target("num len=8", p.Dni).Ok(); err != nil {
		return err
	}
	if err := chk.Target("num min=8 max=15", p.Telefono).Ok(); err != nil {
		return err
	}
	return nil
}
