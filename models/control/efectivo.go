package control

type Efectivo struct {
	ID            int          `json:"efectivo_id"`
	InstitucionID int          `json:"institucion_id"`
	Institucion   *Institucion `json:"institucion,omitempty"`
	PersonaID     int          `json:"persona_id"`
	Persona       *Persona     `json:"persona,omitempty"`
}
