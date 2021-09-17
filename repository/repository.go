package repository

import (
	"fmt"
	"log"

	"github.com/user0608/mujeresapi/utils"
	"gorm.io/gorm"
)

type ChTables []struct {
	tableName string
	recordID  int
}

func consultarUsuario(g *gorm.DB, uc *UsuarioConsult, usuario_id int) error {
	if err := g.Raw("select usuario_is_free(?)", usuario_id).Scan(&uc).Error; err != nil {
		return err
	}
	return nil
}
func existRegister(g *gorm.DB, entity interface{}, id int) error {
	r := g.Find(entity, id)
	if r.Error != nil {
		return r.Error
	}
	if r.RowsAffected == 0 {
		return utils.ErrNothingFind
	}
	return nil
}

func CheckTablesRecord(g *gorm.DB, tables ChTables) error {
	for _, t := range tables {
		var r int
		row := g.Raw("select fn_consult(?,?)", t.tableName, t.recordID).Row()
		if err := row.Err(); err != nil {
			log.Println("Error-0: repository.CheckTableRecord:", err.Error())
			return utils.ErrDataBaseError
		}
		if err := row.Scan(&r); err != nil {
			log.Println("Error-1: repository.CheckTableRecord:", err.Error())
			return utils.ErrDataBaseError
		}
		if r == 0 {
			return fmt.Errorf("El registro con ID: `%d`, no existe existe en los datos de %s", t.recordID, t.tableName)
		}
	}
	return nil
}
