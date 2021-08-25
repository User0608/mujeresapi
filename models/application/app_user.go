package application

type AppUser struct {
	ID              int        `json:"appuser_id" gorm:"autoIncrement:false"`
	Nombre          string     `json:"nombre"`
	ApellidoMaterno string     `json:"apellido_materno"`
	ApellidoPaterno string     `json:"apellido_paterno"`
	Telefono        string     `json:"telefono"`
	Dni             string     `json:"dni"`
	DireccionID     int        `json:"direccion_id"`
	Direccion       *Direccion `json:"direccion"`
}
