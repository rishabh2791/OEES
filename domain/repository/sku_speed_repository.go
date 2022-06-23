package repository

import "oees/domain/entity"

type SKUSpeedRepository interface {
	Create(device *entity.SKUSpeed) (*entity.SKUSpeed, error)
	Get(id string) (*entity.SKUSpeed, error)
	List(conditions string) ([]entity.SKUSpeed, error)
	Update(id string, update *entity.SKUSpeed) (*entity.SKUSpeed, error)
}
