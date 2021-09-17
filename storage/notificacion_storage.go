package storage

import "github.com/user0608/mujeresapi/models/control"

type NotificacionStorage interface {
	CreateNotificacion(n *control.Notificacion) error
	UpdateNotification(n *control.Notificacion) error
	FindAll() ([]control.Notificacion, error)
	FindAllForIntitucionUsuario(usuarioID int) ([]control.Notificacion, error)
	FindByID(id int) (*control.Notificacion, error)
	FindByIDForInstitucion(notificacionID, usuarioID int) (*control.Notificacion, error)
	FindAlertaID(alertaID int) ([]control.Notificacion, error)
	FindColaboradorID(colaboradorID int) ([]control.Notificacion, error)
	FindInstitucionID(institucionID int) ([]control.Notificacion, error)
}
