package application

import (
	"oees/domain/entity"
	"oees/domain/repository"
)

type userRoleApp struct {
	userRoleRepository repository.UserRoleRepository
}

var _ userRoleAppInterface = &userRoleApp{}

func newuserRoleApp(userRoleRepository repository.UserRoleRepository) *userRoleApp {
	return &userRoleApp{
		userRoleRepository: userRoleRepository,
	}
}

func (userRoleApp *userRoleApp) Create(userRole *entity.UserRole) (*entity.UserRole, error) {
	return userRoleApp.userRoleRepository.Create(userRole)
}

func (userRoleApp *userRoleApp) List(conditions string) ([]entity.UserRole, error) {
	return userRoleApp.userRoleRepository.List(conditions)
}

type userRoleAppInterface interface {
	Create(userRole *entity.UserRole) (*entity.UserRole, error)
	List(conditions string) ([]entity.UserRole, error)
}
