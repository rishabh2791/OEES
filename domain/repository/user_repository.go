package repository

import "oees/domain/entity"

type UserRepository interface {
	Create(device *entity.User) (*entity.User, error)
	Get(id string) (*entity.User, error)
	List(conditions string) ([]entity.User, error)
	Update(username string, update *entity.User) (*entity.User, error)
}
