package storage

import "github.com/user0608/mujeresapi/models/control"

type InstitucionStorage interface {
	Create(i *control.Institucion) error
	Update(i *control.Institucion) error
	Find() ([]control.Institucion, error)
	FindByID(institucionID int) (*control.Institucion, error)
	FindByUserName(username string) (*control.Institucion, error)
}
