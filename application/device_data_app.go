package application

import (
	"oees/domain/entity"
	"oees/domain/repository"
)

type deviceDataApp struct {
	deviceDataRepository repository.DeviceDataRepository
}

var _ deviceDataAppInterface = &deviceDataApp{}

func newdeviceDataApp(deviceDataRepository repository.DeviceDataRepository) *deviceDataApp {
	return &deviceDataApp{
		deviceDataRepository: deviceDataRepository,
	}
}

func (deviceDataApp *deviceDataApp) Create(deviceData *entity.DeviceData) (*entity.DeviceData, error) {
	return deviceDataApp.deviceDataRepository.Create(deviceData)
}

func (deviceDataApp *deviceDataApp) List(conditions string) ([]entity.DeviceData, error) {
	return deviceDataApp.deviceDataRepository.List(conditions)
}

type deviceDataAppInterface interface {
	Create(deviceData *entity.DeviceData) (*entity.DeviceData, error)
	List(conditions string) ([]entity.DeviceData, error)
}
