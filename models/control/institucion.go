package control

import "github.com/user0608/mujeresapi/models/application"

type Institucion struct {
	ID          int    `json:"institucion_id" gorm:"autoIncrement:false"`
	Persona     string `json:"persona"`
	Teleono     string `json:"telefon"`
	Email       string `json:"email"`
	Tipo        string `json:"tipo"`
	UsuarioID   int    `json:"usuario_id,omitempty"`
	DireccionID int
	Direccion   *application.Direccion
}
