package application

import (
	"oees/domain/entity"
	"oees/domain/repository"
)

type skuSpeedApp struct {
	skuSpeedRepository repository.SKUSpeedRepository
}

var _ skuSpeedAppInterface = &skuSpeedApp{}

func newskuSpeedApp(skuSpeedRepository repository.SKUSpeedRepository) *skuSpeedApp {
	return &skuSpeedApp{
		skuSpeedRepository: skuSpeedRepository,
	}
}

func (skuSpeedApp *skuSpeedApp) Create(skuSpeed *entity.SKUSpeed) (*entity.SKUSpeed, error) {
	return skuSpeedApp.skuSpeedRepository.Create(skuSpeed)
}

func (skuSpeedApp *skuSpeedApp) Get(skuSpeedname string) (*entity.SKUSpeed, error) {
	return skuSpeedApp.skuSpeedRepository.Get(skuSpeedname)
}

func (skuSpeedApp *skuSpeedApp) List(conditions string) ([]entity.SKUSpeed, error) {
	return skuSpeedApp.skuSpeedRepository.List(conditions)
}

func (skuSpeedApp *skuSpeedApp) Update(skuSpeedname string, update *entity.SKUSpeed) (*entity.SKUSpeed, error) {
	return skuSpeedApp.skuSpeedRepository.Update(skuSpeedname, update)
}

type skuSpeedAppInterface interface {
	Create(skuSpeed *entity.SKUSpeed) (*entity.SKUSpeed, error)
	Get(skuSpeedname string) (*entity.SKUSpeed, error)
	List(conditions string) ([]entity.SKUSpeed, error)
	Update(skuSpeedname string, update *entity.SKUSpeed) (*entity.SKUSpeed, error)
}
