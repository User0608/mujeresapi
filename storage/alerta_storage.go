package storage

import (
	"github.com/user0608/mujeresapi/models/application"
)

type AlertaStorage interface {
	Create(alerta *application.Alerta) error
	FetchAll() ([]application.Alerta, error)
	FetchAllByUserID(userid int) ([]application.Alerta, error)
	FetchByID(alertaID int) (*application.Alerta, error)
}
