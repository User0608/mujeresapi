package repository

import (
	"github.com/user0608/mujeresapi/utils"
	"gorm.io/gorm"
)

type UsuarioConsult struct {
	Response string `gorm:"column:usuario_is_free"`
}

func (u *UsuarioConsult) OK() bool {
	return u.Response == "OK"
}

func existRegister(entity interface{}, id int, g *gorm.DB) error {
	r := g.Find(entity, id)
	if r.Error != nil {
		return r.Error
	}
	if r.RowsAffected == 0 {
		return utils.ErrNothingFind
	}
	return nil
}
