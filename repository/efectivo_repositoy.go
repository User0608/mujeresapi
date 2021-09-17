package repository

import (
	"log"

	"github.com/user0608/mujeresapi/models/auxiliares"
	"github.com/user0608/mujeresapi/models/control"
	"github.com/user0608/mujeresapi/utils"
	"gorm.io/gorm"
)

type EfectivoRepository struct {
	gorm *gorm.DB
}

func NewEfectivoRepository(g *gorm.DB) *EfectivoRepository {
	return &EfectivoRepository{gorm: g}
}

func (r *EfectivoRepository) Create(ae *auxiliares.Efectivo) (*control.Efectivo, error) {
	efectivo := ae.GetEfectivo()
	var direccionID int
	var personaID int
	var efectivoID int
	conn, err := r.gorm.DB()
	if err != nil {
		log.Println("Error-0: EfectivoRepository.Create:", err.Error())
		return nil, utils.ErrDataBaseError
	}
	tx, err := conn.Begin()
	if err != nil {
		log.Println("Error-1: EfectivoRepository.Creaste:", err.Error())
		return nil, utils.ErrDataBaseError
	}
	query := "insert into direccion(provincia,distrito,direccion,referencia) values($1,$2,$3,$4) returning id"
	if err := tx.QueryRow(query, ae.Provincia, ae.Direccion, ae.Direccion, ae.Referencia).Scan(&direccionID); err != nil {
		log.Println("Error-2: EfectivoRepository.Creaste:", err.Error())
		if err := tx.Rollback(); err != nil {
			log.Println("Error-3: EfectivoRepository.Creaste:", err.Error())
		}
		return nil, utils.ErrDataBaseError
	}
	efectivo.Persona.Direccion.ID = int(direccionID)
	efectivo.Persona.DireccionID = int(direccionID)
	query = "insert into persona(nombre,apellido_materno,apellido_paterno,telefono,dni,direccion_id) values($1,$2,$3,$4,$5,$6) returning id"
	if err := tx.QueryRow(query, ae.Nombre, ae.ApellidoMaterno, ae.ApellidoPaterno, ae.Telefono, ae.DNI, direccionID).Scan(&personaID); err != nil {
		log.Println("Error-6: EfectivoRepository.Creaste:", err.Error())
		if err := tx.Rollback(); err != nil {
			log.Println("Error-7: EfectivoRepository.Creaste:", err.Error())
		}
		return nil, utils.ErrDataBaseError
	}
	efectivo.Persona.ID = int(personaID)
	efectivo.PersonaID = int(personaID)

	query = "insert into efectivo(institucion_id,persona_id) values($1,$2)  returning id"
	if err := tx.QueryRow(query, ae.InstitucionID, personaID).Scan(&efectivoID); err != nil {
		log.Println("Error-10: EfectivoRepository.Creaste:", err.Error())
		if err := tx.Rollback(); err != nil {
			log.Println("Error-11: EfectivoRepository.Creaste:", err.Error())
		}
		return nil, utils.ErrDataBaseError
	}
	efectivo.ID = int(efectivoID)
	if err := tx.Commit(); err != nil {
		log.Println("Error-14: EfectivoRepository.Creaste:", err.Error())
		return nil, utils.ErrDataBaseError
	}
	return efectivo, nil
}

func (r *EfectivoRepository) FindAll() ([]control.Efectivo, error) {
	efectivos := []control.Efectivo{}
	if res := r.gorm.Preload("Persona").Find(&efectivos); res.Error != nil {
		log.Println("Error-0: EfectivoRepository.FindAll", res.Error.Error())
		return nil, utils.ErrDataBaseError
	}
	return efectivos, nil
}

func (r *EfectivoRepository) FindAllByInstitucionID(institucionID int) ([]control.Efectivo, error) {
	efectivos := []control.Efectivo{}
	if res := r.gorm.Preload("Persona").Find(&efectivos, "institucion_id = ?", institucionID); res.Error != nil {
		log.Println("Error-0: EfectivoRepository.FindAll", res.Error.Error())
		return nil, utils.ErrDataBaseError
	}
	return efectivos, nil
}
func (r *EfectivoRepository) FindAllByUsuarioID(usuarioID int) ([]control.Efectivo, error) {
	var institucionID int
	row := r.gorm.Table("institucion").Where("usuario_id = ?", usuarioID).Select("id").Row()
	if err := row.Err(); err != nil {
		log.Println("Error-0: EfectivoRepository.FindAllByUsuarioID", err.Error())
		return nil, utils.ErrDataBaseError
	}
	if err := row.Scan(&institucionID); err != nil {
		return nil, utils.ErrIDInvalido
	}
	efectivos := []control.Efectivo{}
	if res := r.gorm.Preload("Persona").Find(&efectivos, "institucion_id = ?", institucionID); res.Error != nil {
		log.Println("Error-1: EfectivoRepository.FindAllByUsuarioID", res.Error.Error())
		return nil, utils.ErrDataBaseError
	}
	return efectivos, nil
}
