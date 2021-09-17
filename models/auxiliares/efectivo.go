package auxiliares

import (
	"github.com/user0608/mujeresapi/models/application"
	"github.com/user0608/mujeresapi/models/control"
)

type Efectivo struct {
	InstitucionID   int    `json:"institucion_id"`
	Nombre          string `json:"nombre"`
	ApellidoPaterno string `json:"apellido_paterno"`
	ApellidoMaterno string `json:"apellido_materno"`
	Telefono        string `json:"telefono"`
	DNI             string `json:"dni"`
	Direccion       string `json:"direccion"`
	Provincia       string `json:"provincia"`
	Distrito        string `json:"distrito"`
	Referencia      string `json:"referencia"`
}

func (e *Efectivo) GetDireccion() *application.Direccion {
	return &application.Direccion{
		Provincia:  e.Provincia,
		Distrito:   e.Distrito,
		Direccion:  e.Direccion,
		Referencia: e.Referencia,
	}
}
func (e *Efectivo) GetPersona() *control.Persona {
	return &control.Persona{
		Nombre:          e.Nombre,
		ApellidoPaterno: e.ApellidoPaterno,
		ApellidoMaterno: e.ApellidoMaterno,
		Telefono:        e.Telefono,
		Dni:             e.DNI,
		Direccion:       e.GetDireccion(),
	}
}
func (e *Efectivo) GetEfectivo() *control.Efectivo {
	return &control.Efectivo{
		InstitucionID: e.InstitucionID,
		Persona:       e.GetPersona(),
	}
}
