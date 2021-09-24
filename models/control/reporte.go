package control

import (
	"errors"
	"time"
)

type Reporte struct {
	Descripcion     string `json:"descripcion"`
	Notificacion_Id int    `json:"notificacion_id"`
}

type ReporteTable struct {
	ID int `json:"id"`
	Reporte
	Notificacion *Notificacion `json:"notificacion,omitempty"`
	Created_At   time.Time     `json:"created_at"`
	Updated_At   time.Time     `json:"create_at"`
}

func (r *Reporte) Valid() error {
	if r.Descripcion == "" {
		return errors.New("La descripcion no puede estar vacia")
	}
	if r.Notificacion_Id == 0 {
		return errors.New("Notificacion ID no puede ser 0")
	}
	return nil
}
func (*ReporteTable) TableName() string {
	return "reporte"
}
