package services

import (
	"github.com/user0608/mujeresapi/models/authentication"
	"github.com/user0608/mujeresapi/storage"
)

type RoleService struct {
	storage storage.RoleStorage
}

func NewRoleService(s storage.RoleStorage) *RoleService {
	return &RoleService{
		storage: s,
	}
}
func (s *RoleService) FindRoles() ([]authentication.Roles, error) {
	return s.storage.GetAll()
}
