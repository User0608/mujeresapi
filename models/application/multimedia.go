package application

type Multimedia struct {
	ID       int    `json:"multimedia_id"`
	Url      string `json:"url"`
	Tipo     string `json:"tipo"`
	AlertaID uint   `json:"alerta_id"`
}
