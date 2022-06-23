package repository

import "oees/domain/entity"

type PlantRepository interface {
	Create(device *entity.Plant) (*entity.Plant, error)
	Get(id string) (*entity.Plant, error)
	List(conditions string) ([]entity.Plant, error)
	Update(id string, update *entity.Plant) (*entity.Plant, error)
}
