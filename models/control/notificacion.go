package control

import (
	"fmt"
	"strings"
)

type Notificacion struct {
	ID            int    `json:"notificacion_id"`
	Titulo        string `json:"titulo"`
	AlertaID      int    `json:"alerta_id"`
	InstitucionID int    `json:"institucion_id"`
	Nivel         int    `json:"nivel"`
	Descripcion   string `json:"descripcion"`
	ColaboradorID int    `json:"colaborador_id"`
}

func (n *Notificacion) Validar() error {
	if n.AlertaID == 0 {
		return fmt.Errorf("Alerta ID, no puede estar vacio!")
	}
	if n.ColaboradorID == 0 {
		return fmt.Errorf("Colaborador ID, no puede estar vacio!")
	}
	if n.InstitucionID == 0 {
		return fmt.Errorf("Institucion ID, no puede estar vacio!")
	}
	if len(strings.Trim(n.Titulo, " ")) < 4 {
		return fmt.Errorf("El campo titulo, tiene que tener al menos 4 caracteres no vacios")
	}
	return nil
}
