package application

import (
	"oees/domain/entity"
	"oees/domain/repository"
)

type skuApp struct {
	skuRepository repository.SKURepository
}

var _ skuAppInterface = &skuApp{}

func newskuApp(skuRepository repository.SKURepository) *skuApp {
	return &skuApp{
		skuRepository: skuRepository,
	}
}

func (skuApp *skuApp) Create(sku *entity.SKU) (*entity.SKU, error) {
	return skuApp.skuRepository.Create(sku)
}

func (skuApp *skuApp) Get(skuname string) (*entity.SKU, error) {
	return skuApp.skuRepository.Get(skuname)
}

func (skuApp *skuApp) List(conditions string) ([]entity.SKU, error) {
	return skuApp.skuRepository.List(conditions)
}

func (skuApp *skuApp) Update(skuname string, update *entity.SKU) (*entity.SKU, error) {
	return skuApp.skuRepository.Update(skuname, update)
}

type skuAppInterface interface {
	Create(sku *entity.SKU) (*entity.SKU, error)
	Get(skuname string) (*entity.SKU, error)
	List(conditions string) ([]entity.SKU, error)
	Update(skuname string, update *entity.SKU) (*entity.SKU, error)
}
