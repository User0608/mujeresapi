package repository

import (
	"log"

	"github.com/user0608/mujeresapi/models/control"
	"github.com/user0608/mujeresapi/utils"
	"gorm.io/gorm"
)

type NotificacionRepository struct {
	gorm *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificacionRepository {
	return &NotificacionRepository{gorm: db}
}

func (r *NotificacionRepository) CreateNotificacion(n *control.Notificacion) error {
	ch := ChTables{{"alerta", n.AlertaID}, {"institucion", n.InstitucionID}, {"colaborador", n.ColaboradorID}}
	if err := CheckTablesRecord(r.gorm, ch); err != nil {
		return err
	}
	if res := r.gorm.Create(n); res.Error != nil {
		log.Println("Error-0: NotificacionRepository.CreateNotificacion", res.Error.Error())
		return utils.ErrDataBaseError
	}
	return nil
}

func (r *NotificacionRepository) UpdateNotification(n *control.Notificacion) error {
	ch := ChTables{{"notificacion", n.ID}, {"alerta", n.AlertaID}, {"institucion", n.InstitucionID}, {"colaborador", n.ColaboradorID}}
	if err := CheckTablesRecord(r.gorm, ch); err != nil {
		return err
	}
	if res := r.gorm.Updates(n); res.Error != nil {
		log.Println("Error-0: NotificacionRepository.UpdateNotification:", res.Error.Error())
		return utils.ErrDataBaseError
	}
	return nil
}
func (r *NotificacionRepository) FindAll() ([]control.Notificacion, error) {
	notificaciones := []control.Notificacion{}
	if res := r.gorm.Find(&notificaciones); res.Error != nil {
		log.Println("Error-0: NotificacionRepository.FindAll:", res.Error.Error())
		return nil, utils.ErrDataBaseError
	}
	return notificaciones, nil
}
func (r *NotificacionRepository) FindByID(id int) (*control.Notificacion, error) {
	var notificacion control.Notificacion
	if res := r.gorm.Limit(1).Find(&notificacion, id); res.Error != nil {
		log.Println("Error-0: NotificacionRepository.FindByID:", res.Error.Error())
		return nil, utils.ErrDataBaseError
	} else {
		if res.RowsAffected == 0 {
			return nil, utils.ErrNothingFind
		}
	}
	return &notificacion, nil
}
func (r *NotificacionRepository) FindAllForIntitucionUsuario(usuarioID int) ([]control.Notificacion, error) {
	notificaciones := []control.Notificacion{}
	if res := r.gorm.Raw("select * from fn_find_all_notificacion_for_institucion(?)", usuarioID).Scan(&notificaciones); res.Error != nil {
		log.Println("Error-0: NotificacionRepository.FindAllForIntitucionUsuario: ", res.Error.Error())
		return nil, utils.ErrDataBaseError
	}
	return notificaciones, nil
}
func (r *NotificacionRepository) FindByIDForInstitucion(notificacionID, usuarioID int) (*control.Notificacion, error) {
	var notificacion control.Notificacion
	if res := r.gorm.Limit(1).Raw("select * from fn_find_notificacion_for_institucion(?,?)", notificacionID, usuarioID).Scan(&notificacion); res.Error != nil {
		log.Println("Error-0: NotificacionRepository.FindByIDForInstitucion:", res.Error.Error())
		return nil, utils.ErrDataBaseError
	} else {
		if res.RowsAffected == 0 {
			return nil, utils.ErrNothingFind
		} else {
			if err := r.gorm.Limit(1).Preload("Institucion.Direccion").Preload("Colaborador.Persona").Preload("Alerta.Multimedias").Preload("Alerta.Usuario.Direccion").Find(&notificacion, notificacion.ID).Error; err != nil {
				log.Println("Error-0: NotificacionRepository.FindByIDForInstitucion:", res.Error.Error())
				return nil, utils.ErrDataBaseError
			}
		}
	}
	return &notificacion, nil
}
func (r *NotificacionRepository) FindAlertaID(alertaID int) ([]control.Notificacion, error) {
	notificacion := []control.Notificacion{}
	if res := r.gorm.Limit(1).Find(&notificacion, "alerta_id = ?", alertaID); res.Error != nil {
		log.Println("Error-0: NotificacionRepository.FindAlertaID:", res.Error.Error())
		return nil, utils.ErrDataBaseError
	}
	return notificacion, nil
}
func (r *NotificacionRepository) FindColaboradorID(colaboradorID int) ([]control.Notificacion, error) {
	notificacion := []control.Notificacion{}
	if res := r.gorm.Limit(1).Find(&notificacion, "colaborador_id = ?", colaboradorID); res.Error != nil {
		log.Println("Error-0: NotificacionRepository.FindColaboradorID:", res.Error.Error())
		return nil, utils.ErrDataBaseError
	}
	return notificacion, nil
}
func (r *NotificacionRepository) FindInstitucionID(institucionID int) ([]control.Notificacion, error) {
	notificacion := []control.Notificacion{}
	if res := r.gorm.Limit(1).Find(&notificacion, "institucion_id = ?", institucionID); res.Error != nil {
		log.Println("Error-0: NotificacionRepository.FindInstitucionID:", res.Error.Error())
		return nil, utils.ErrDataBaseError
	}
	return notificacion, nil
}
