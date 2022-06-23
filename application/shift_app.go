package application

import (
	"oees/domain/entity"
	"oees/domain/repository"
)

type shiftApp struct {
	shiftRepository repository.ShiftRepository
}

var _ shiftAppInterface = &shiftApp{}

func newshiftApp(shiftRepository repository.ShiftRepository) *shiftApp {
	return &shiftApp{
		shiftRepository: shiftRepository,
	}
}

func (shiftApp *shiftApp) Create(shift *entity.Shift) (*entity.Shift, error) {
	return shiftApp.shiftRepository.Create(shift)
}

func (shiftApp *shiftApp) Get(shiftname string) (*entity.Shift, error) {
	return shiftApp.shiftRepository.Get(shiftname)
}

func (shiftApp *shiftApp) List(conditions string) ([]entity.Shift, error) {
	return shiftApp.shiftRepository.List(conditions)
}

func (shiftApp *shiftApp) Update(shiftname string, update *entity.Shift) (*entity.Shift, error) {
	return shiftApp.shiftRepository.Update(shiftname, update)
}

type shiftAppInterface interface {
	Create(shift *entity.Shift) (*entity.Shift, error)
	Get(shiftname string) (*entity.Shift, error)
	List(conditions string) ([]entity.Shift, error)
	Update(shiftname string, update *entity.Shift) (*entity.Shift, error)
}
