package application

import (
	"oees/domain/entity"
	"oees/domain/repository"
)

type downtimeApp struct {
	downtimeRepository repository.DowntimeRepository
}

var _ downtimeAppInterface = &downtimeApp{}

func newdowntimeApp(downtimeRepository repository.DowntimeRepository) *downtimeApp {
	return &downtimeApp{
		downtimeRepository: downtimeRepository,
	}
}

func (downtimeApp *downtimeApp) Create(downtime *entity.Downtime) (*entity.Downtime, error) {
	return downtimeApp.downtimeRepository.Create(downtime)
}

func (downtimeApp *downtimeApp) Get(downtimename string) (*entity.Downtime, error) {
	return downtimeApp.downtimeRepository.Get(downtimename)
}

func (downtimeApp *downtimeApp) List(conditions string) ([]entity.Downtime, error) {
	return downtimeApp.downtimeRepository.List(conditions)
}

func (downtimeApp *downtimeApp) Update(downtimename string, update *entity.Downtime) (*entity.Downtime, error) {
	return downtimeApp.downtimeRepository.Update(downtimename, update)
}

type downtimeAppInterface interface {
	Create(downtime *entity.Downtime) (*entity.Downtime, error)
	Get(downtimename string) (*entity.Downtime, error)
	List(conditions string) ([]entity.Downtime, error)
	Update(downtimename string, update *entity.Downtime) (*entity.Downtime, error)
}
