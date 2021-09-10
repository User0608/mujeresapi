package repository

import (
	"github.com/user0608/mujeresapi/utils"
	"gorm.io/gorm"
)

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
