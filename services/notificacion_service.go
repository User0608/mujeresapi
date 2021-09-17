package services

import (
	"fmt"

	"github.com/user0608/mujeresapi/models/control"
	"github.com/user0608/mujeresapi/storage"
	"github.com/user0608/mujeresapi/utils"
)

type NotificacionService struct {
	Storage storage.NotificacionStorage
}

func NewNotificacionService(s storage.NotificacionStorage) *NotificacionService {
	return &NotificacionService{Storage: s}
}

func (s *NotificacionService) Create(n *control.Notificacion) error {
	if err := n.Validar(); err != nil {
		return err
	}
	return s.Storage.CreateNotificacion(n)
}
func (s *NotificacionService) Update(n *control.Notificacion) error {
	if n.ID == 0 {
		return fmt.Errorf("El ID, es necesarios, no puede estar vacio!")
	}
	if err := n.Validar(); err != nil {
		return err
	}
	return s.Storage.UpdateNotification(n)
}
func (s *NotificacionService) FindAll() ([]control.Notificacion, error) {
	return s.Storage.FindAll()
}
func (s *NotificacionService) FindAllForInstitucion(usuarioID int) ([]control.Notificacion, error) {
	return s.Storage.FindAllForIntitucionUsuario(usuarioID)
}
func (s *NotificacionService) FindByID(id int) (*control.Notificacion, error) {
	if id == 0 {
		return nil, utils.ErrNothingFind
	}
	return s.Storage.FindByID(id)
}
func (s *NotificacionService) FindByIDForInstitucion(notificacionID, usuarioID int) (*control.Notificacion, error) {
	if notificacionID == 0 || usuarioID == 0 {
		return nil, utils.ErrNothingFind
	}
	return s.Storage.FindByIDForInstitucion(notificacionID, usuarioID)
}
func (s *NotificacionService) FindAlertaID(alertaID int) ([]control.Notificacion, error) {
	if alertaID == 0 {
		return nil, utils.ErrNothingFind
	}
	return s.Storage.FindAlertaID(alertaID)
}
func (s *NotificacionService) FindColaboradorID(colaboradorID int) ([]control.Notificacion, error) {
	if colaboradorID == 0 {
		return nil, utils.ErrNothingFind
	}
	return s.Storage.FindColaboradorID(colaboradorID)
}
func (s *NotificacionService) FindInstitucionID(institucionID int) ([]control.Notificacion, error) {
	if institucionID == 0 {
		return nil, utils.ErrNothingFind
	}
	return s.Storage.FindInstitucionID(institucionID)
}
