package storage

import "github.com/user0608/mujeresapi/models/control"

type ReporteStorage interface {
	Create(r control.Reporte) (*control.ReporteTable, error)
	Update(id int, r control.Reporte) (*control.ReporteTable, error)
	Delete(id int) error
	FindAll() ([]control.ReporteTable, error)
	FindByID(id int) (*control.ReporteTable, error)
	FindByNotificacionID(notificacionID int) (*control.ReporteTable, error)
}
