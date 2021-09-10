package control_test

import (
	"testing"

	"github.com/user0608/mujeresapi/models/control"
)

func TestPersana(t *testing.T) {
	p := control.Persona{
		Nombre:          "Kevin",
		ApellidoPaterno: "Saucedo",
		ApellidoMaterno: "Huaman",
		Telefono:        "995831649",
		Dni:             "73491346",
	}
	if err := p.Validate(); err != nil {
		t.Error(err.Error())
	}
}
