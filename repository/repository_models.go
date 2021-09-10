package repository

type UsuarioConsult struct {
	Response string `gorm:"column:usuario_is_free"`
}

func (u *UsuarioConsult) OK() bool {
	return u.Response == "OK"
}

func (u *UsuarioConsult) NotExistUsuario() bool {
	return u.Response == "FK"
}
