package control

import (
	"github.com/user0608/mujeresapi/models/application"
)

type Persona struct {
	ID              int                    `json:"id"`
	Nombre          string                 `json:"nombre"`
	ApellidoMaterno string                 `json:"apellido_materno"`
	ApellidoPaterno string                 `json:"apellido_paterno"`
	Telefono        string                 `json:"telefon"`
	Dni             string                 `json:"dni"`
	DireccionID     int                    `json:"direccion_id"`
	Direccion       *application.Direccion `json:"direccion"`
}
