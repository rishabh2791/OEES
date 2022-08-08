package application

import (
	"oees/domain/entity"
	"oees/domain/repository"
	"oees/domain/value_objects"
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

func (deviceDataApp *deviceDataApp) TotalDeviceData(conditions string) (*value_objects.TotalDeviceData, error) {
	return deviceDataApp.deviceDataRepository.TotalDeviceData(conditions)
}

type deviceDataAppInterface interface {
	Create(deviceData *entity.DeviceData) (*entity.DeviceData, error)
	List(conditions string) ([]entity.DeviceData, error)
	TotalDeviceData(conditions string) (*value_objects.TotalDeviceData, error)
}
