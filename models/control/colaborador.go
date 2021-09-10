package control

import "github.com/user0608/mujeresapi/models/authentication"

type Colaborador struct {
	ID        int                     `json:"colaborador_id"`
	PersonaID int                     `json:"persona_id"`
	Persona   *Persona                `json:"persona,omitempty"`
	UsuarioID int                     `json:"usuario_id"`
	Usuario   *authentication.Usuario `json:"usuario,omitempty"`
}
