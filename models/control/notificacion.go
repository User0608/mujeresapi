package control

type Notificacion struct {
	ID            int    `json:"notificacion_id"`
	AlertaID      int    `json:"alerta_id"`
	InstitucionID int    `json:"institucion_id"`
	Titulo        string `json:"titulo"`
	Nivel         string `json:"nivel"`
	Descripcion   string `json:"description"`
	ColaboradorID int    `json:"colaborador_id"`
}
