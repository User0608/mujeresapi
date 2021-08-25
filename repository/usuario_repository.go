package repository

import (
	"github.com/user0608/mujeresapi/models/application"
	"github.com/user0608/mujeresapi/models/authentication"
	"github.com/user0608/mujeresapi/utils"
	"gorm.io/gorm"
)

type UsuarioRepository struct {
	gdb *gorm.DB
}

func NewUsuarioRepository(db *gorm.DB) *UsuarioRepository {
	return &UsuarioRepository{
		gdb: db,
	}
}

func (u *UsuarioRepository) GetAllUsuarios() ([]authentication.Usuario, error) {
	usuarios := []authentication.Usuario{}
	if err := u.gdb.Preload("Roles").Omit("Password").Find(&usuarios).Error; err != nil {
		return nil, err
	}
	return usuarios, nil
}
func (u *UsuarioRepository) GetUser(username string) (*authentication.Usuario, bool, error) {
	var users = []authentication.Usuario{}
	result := u.gdb.Find(&users, "username = ?", username)
	if err := result.Error; err != nil {
		return nil, false, err
	}
	if result.RowsAffected == 0 {
		return nil, true, nil
	}
	return &users[0], false, nil
}

func (u *UsuarioRepository) LoginUsuario(username, password string) (*authentication.Usuario, error) {
	usuario := &authentication.Usuario{}
	query := "select * from ps_sign_in(?,?)"
	result := u.gdb.Raw(query, username, password).Scan(usuario)
	if err := result.Error; err != nil {
		return nil, err
	} else {
		if result.RowsAffected == 1 {
			err = u.gdb.Preload("Roles").Omit("Password").Find(&usuario, usuario.ID).Error
			if err != nil {
				return nil, err
			}
		} else {
			return nil, utils.ErrNothingFind
		}
	}
	return usuario, nil
}

func (u *UsuarioRepository) RegistrarUsuario(user *authentication.Usuario) error {
	if u.gdb.Find(&[]authentication.Usuario{}, "username = ?", user.Username).RowsAffected > 0 {
		return utils.ErrExistUsername
	}
	return u.gdb.Create(user).Error
}

func (u *UsuarioRepository) RegistrarDetalleAppUsuario(appuser *application.AppUser) error {
	return u.gdb.Create(appuser).Error
}
func (u *UsuarioRepository) UpdateDetalleAppUsuario(appuser *application.AppUser) error {
	if u.gdb.First(&application.AppUser{}, appuser.ID).RowsAffected == 0 {
		return utils.ErrNothingFind
	}
	return u.gdb.Session(&gorm.Session{FullSaveAssociations: true}).Updates(appuser).Error
}

func (u *UsuarioRepository) GetUsuarioDetalleApp(userid int) (*application.AppUser, error) {
	usuario := &application.AppUser{}
	response := u.gdb.Preload("Direccion").Find(usuario, "id = ?", userid)
	if err := response.Error; err != nil {
		return nil, err
	}
	if response.RowsAffected == 0 {
		return nil, utils.ErrNothingFind
	}
	return usuario, nil
}
