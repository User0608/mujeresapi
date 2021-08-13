package authentication

type Roles struct {
	ID          int    `json:"role_id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}
