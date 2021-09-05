package storage

import "github.com/user0608/mujeresapi/models/application"

type MultimediaStorage interface {
	RegistrarMultimedia(m []application.Multimedia) error
	VerificarAlertaID(alertaID int) error
}
