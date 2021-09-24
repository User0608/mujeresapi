package control

import "time"

type Asignacion struct {
	EfectivoID     int `json:"efectivo_id"`
	NotificacionID int `json:"notificacion_id"`
}
type AsignacionTable struct {
	ID        int       `json:"asignacion_id"`
	CreatedAt time.Time `json:"created_at"`
	Asignacion
	Notificacion *Notificacion `json:"notificacion,omitempty"`
	Efectivo     *Efectivo     `json:"efectivo,omitempty"`
}

func (*Asignacion) TableName() string {
	return "asignacion"
}
