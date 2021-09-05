package application

type Multimedia struct {
	ID       int    `json:"multimedia_id"`
	Url      string `json:"url"`
	Tipo     string `json:"tipo"`
	AlertaID int    `json:"alerta_id"`
}
