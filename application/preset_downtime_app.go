package application

import (
	"oees/domain/entity"
	"oees/domain/repository"
)

type presetDowntimeApp struct {
	presetDowntimeRepository repository.PresetDowntimeRepository
}

var _ PresetDowntimeAppInterface = &presetDowntimeApp{}

func newPresetDowntimeApp(presetDowntimeRepository repository.PresetDowntimeRepository) *presetDowntimeApp {
	return &presetDowntimeApp{
		presetDowntimeRepository: presetDowntimeRepository,
	}
}

type PresetDowntimeAppInterface interface {
	Create(downtime *entity.PresetDowntime) (*entity.PresetDowntime, error)
	List(conditions string) ([]entity.PresetDowntime, error)
}

func (presetDowntimeApp *presetDowntimeApp) Create(downtime *entity.PresetDowntime) (*entity.PresetDowntime, error) {
	return presetDowntimeApp.presetDowntimeRepository.Create(downtime)
}

func (presetDowntimeApp *presetDowntimeApp) List(conditions string) ([]entity.PresetDowntime, error) {
	return presetDowntimeApp.presetDowntimeRepository.List(conditions)
}
