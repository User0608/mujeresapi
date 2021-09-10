package control

import (
	"github.com/user0608/kcheck"
	"github.com/user0608/mujeresapi/models/application"
)

type Institucion struct {
	ID          int                    `json:"institucion_id"`
	Persona     string                 `json:"persona"`
	Telefono    string                 `json:"telefono"`
	Email       string                 `json:"email"`
	Tipo        string                 `json:"tipo"`
	UsuarioID   int                    `json:"usuario_id,omitempty"`
	DireccionID int                    `json:"direccion_id"`
	Direccion   *application.Direccion `json:"direccion"`
}

func (i *Institucion) Validate() error {
	chk := kcheck.New()
	if err := chk.Target("min=2 basic", i.Persona).Ok(); err != nil {
		return err
	}
	if err := chk.Target("num min=8", i.Telefono).Ok(); err != nil {
		return err
	}
	if err := chk.Target("min=8", i.Email).Ok(); err != nil {
		return err
	}
	return nil
}
