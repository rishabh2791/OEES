package application

import (
	"oees/domain/entity"
	"oees/domain/repository"
)

type plantApp struct {
	plantRepository repository.PlantRepository
}

var _ plantAppInterface = &plantApp{}

func newplantApp(plantRepository repository.PlantRepository) *plantApp {
	return &plantApp{
		plantRepository: plantRepository,
	}
}

func (plantApp *plantApp) Create(plant *entity.Plant) (*entity.Plant, error) {
	return plantApp.plantRepository.Create(plant)
}

func (plantApp *plantApp) Get(plantname string) (*entity.Plant, error) {
	return plantApp.plantRepository.Get(plantname)
}

func (plantApp *plantApp) List(conditions string) ([]entity.Plant, error) {
	return plantApp.plantRepository.List(conditions)
}

func (plantApp *plantApp) Update(plantname string, update *entity.Plant) (*entity.Plant, error) {
	return plantApp.plantRepository.Update(plantname, update)
}

type plantAppInterface interface {
	Create(plant *entity.Plant) (*entity.Plant, error)
	Get(plantname string) (*entity.Plant, error)
	List(conditions string) ([]entity.Plant, error)
	Update(plantname string, update *entity.Plant) (*entity.Plant, error)
}
