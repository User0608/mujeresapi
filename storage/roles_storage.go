package storage

import "github.com/user0608/mujeresapi/models/authentication"

type RoleStorage interface {
	GetAll() ([]authentication.Roles, error)
}
