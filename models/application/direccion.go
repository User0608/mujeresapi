package application

type Direccion struct {
	ID         int    `json:"direccion_id"`
	Provincia  string `json:"provincia"`
	Distrito   string `json:"distrito"`
	Direccion  string `json:"direccion"`
	Referencia string `json:"referencia"`
}
