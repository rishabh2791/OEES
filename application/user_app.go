package application

import (
	"oees/domain/entity"
	"oees/domain/repository"
)

type userApp struct {
	userRepository repository.UserRepository
}

var _ UserAppInterface = &userApp{}

func newUserApp(userRepository repository.UserRepository) *userApp {
	return &userApp{
		userRepository: userRepository,
	}
}

func (userApp *userApp) Create(user *entity.User) (*entity.User, error) {
	return userApp.userRepository.Create(user)
}

func (userApp *userApp) Get(username string) (*entity.User, error) {
	return userApp.userRepository.Get(username)
}

func (userApp *userApp) List(conditions string) ([]entity.User, error) {
	return userApp.userRepository.List(conditions)
}

func (userApp *userApp) Update(username string, update *entity.User) (*entity.User, error) {
	return userApp.userRepository.Update(username, update)
}

type UserAppInterface interface {
	Create(user *entity.User) (*entity.User, error)
	Get(username string) (*entity.User, error)
	List(conditions string) ([]entity.User, error)
	Update(username string, update *entity.User) (*entity.User, error)
}
