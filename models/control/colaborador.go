package control

import "github.com/user0608/mujeresapi/models/authentication"

type Colaborador struct {
	ID      int
	Persona Persona
	Usuario authentication.Usuario
}
