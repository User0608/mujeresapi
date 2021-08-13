package authentication

type Usuario struct {
	ID       int      `json:"usuario_id"`
	Username string   `json:"username"`
	Password string   `json:"password,omitempty"`
	Estado   bool     `json:"estado"`
	Roles    []*Roles `gorm:"many2many:usuario_roles;" json:"roles"`
}

func (u *Usuario) GetRoles() []string {
	roles := []string{}
	if u.Roles != nil {
		for _, rol := range u.Roles {
			roles = append(roles, rol.Nombre)
		}
	}
	return roles
}
