package repository

import "oees/domain/entity"

type SKURepository interface {
	Create(device *entity.SKU) (*entity.SKU, error)
	Get(id string) (*entity.SKU, error)
	List(conditions string) ([]entity.SKU, error)
	Update(id string, update *entity.SKU) (*entity.SKU, error)
}
