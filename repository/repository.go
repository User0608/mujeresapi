package repository

import (
	"github.com/user0608/mujeresapi/utils"
	"gorm.io/gorm"
)

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
