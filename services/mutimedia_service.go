package services

import (
	"github.com/user0608/mujeresapi/models/application"
	"github.com/user0608/mujeresapi/storage"
	"github.com/user0608/mujeresapi/utils"
)

type MultimediaService struct {
	storage storage.MultimediaStorage
}

func NewMultimediaService(s storage.MultimediaStorage) *MultimediaService {
	return &MultimediaService{storage: s}
}

func (s *MultimediaService) RegistrarMultimedia(m []application.Multimedia) error {
	return s.storage.RegistrarMultimedia(m)
}

func (s *MultimediaService) VerificarAlertaID(alertaID int) error {
	if alertaID == 0 {
		return utils.ErrNothingFind
	}
	return s.storage.VerificarAlertaID(alertaID)
}
