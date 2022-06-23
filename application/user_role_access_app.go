package application

import (
	"oees/domain/entity"
	"oees/domain/repository"
)

type userRoleAccessApp struct {
	userRoleAccessRepository repository.UserRoleAccessRepository
}

var _ UserRoleAccessAppInterface = &userRoleAccessApp{}

func NewUserRoleAccessApp(userRoleAccessRepository repository.UserRoleAccessRepository) *userRoleAccessApp {
	return &userRoleAccessApp{
		userRoleAccessRepository: userRoleAccessRepository,
	}
}

type UserRoleAccessAppInterface interface {
	Create(userRoleAccess *entity.UserRoleAccess) (*entity.UserRoleAccess, error)
	List(userRole string) ([]entity.UserRoleAccess, error)
	Update(userRole string, userRoleAccess *entity.UserRoleAccess) (*entity.UserRoleAccess, error)
}

func (userRoleAccessApp *userRoleAccessApp) Create(userRoleAccess *entity.UserRoleAccess) (*entity.UserRoleAccess, error) {
	return userRoleAccessApp.userRoleAccessRepository.Create(userRoleAccess)
}

func (userRoleAccessApp *userRoleAccessApp) List(userRole string) ([]entity.UserRoleAccess, error) {
	return userRoleAccessApp.userRoleAccessRepository.List(userRole)
}

func (userRoleAccessApp *userRoleAccessApp) Update(userRole string, userRoleAccess *entity.UserRoleAccess) (*entity.UserRoleAccess, error) {
	return userRoleAccessApp.userRoleAccessRepository.Update(userRole, userRoleAccess)
}
