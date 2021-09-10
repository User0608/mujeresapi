package services

import (
	"log"

	"github.com/user0608/mujeresapi/models/control"
	"github.com/user0608/mujeresapi/storage"
	"github.com/user0608/mujeresapi/utils"
)

type InsitucionService struct {
	storage storage.InstitucionStorage
}

func NewInstitucionService(s storage.InstitucionStorage) *InsitucionService {
	return &InsitucionService{storage: s}
}

func (s *InsitucionService) CreateInstitucion(institucion *control.Institucion) error {
	if institucion == nil {
		return utils.ErrNullEntity
	}
	if err := institucion.Validate(); err != nil {
		log.Println("Error: InsitucionService.CreateInstitucion:", err.Error())
		return utils.ErrInvalidData
	}
	return s.storage.Create(institucion)
}
func (s *InsitucionService) UpdateInstitucion(institucion *control.Institucion) error {
	if institucion == nil {
		return utils.ErrNullEntity
	}
	if institucion.ID == 0 {
		return utils.ErrIdIsNeeded
	}
	if err := institucion.Validate(); err != nil {
		log.Println("Error: InsitucionService.CreateInstitucion:", err.Error())
		return utils.ErrInvalidData
	}
	return s.storage.Update(institucion)
}
func (s *InsitucionService) FindInstitucion() ([]control.Institucion, error) {
	return s.storage.Find()
}
func (s *InsitucionService) FindInstitucionByID(institucionID int) (*control.Institucion, error) {
	if institucionID == 0 {
		return nil, utils.ErrInvalidData
	}
	return s.storage.FindByID(institucionID)
}
func (s *InsitucionService) FindInstitucionByUserName(username string) (*control.Institucion, error) {
	if len(username) < 5 {
		return nil, utils.ErrInvalidData
	}
	return s.storage.FindByUserName(username)
}
