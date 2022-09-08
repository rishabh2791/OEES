package repository

import "oees/domain/entity"

type UserRoleAccessRepository interface {
	Create(userRoleAccess *entity.UserRoleAccess) (*entity.UserRoleAccess, error)
	List(userRole string) ([]entity.UserRoleAccess, error)
	Update(userRole string, userRoleAccess *entity.UserRoleAccess) (*entity.UserRoleAccess, error)
}
