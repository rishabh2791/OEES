package repository

import "oees/domain/entity"

type UserRoleRepository interface {
	Create(userRole *entity.UserRole) (*entity.UserRole, error)
	List(conditions string) ([]entity.UserRole, error)
}
