package application

import (
	"oees/domain/entity"
	"oees/domain/repository"
)

type deviceApp struct {
	deviceRepository repository.DeviceRepository
}

var _ deviceAppInterface = &deviceApp{}

func newdeviceApp(deviceRepository repository.DeviceRepository) *deviceApp {
	return &deviceApp{
		deviceRepository: deviceRepository,
	}
}

func (deviceApp *deviceApp) Create(device *entity.Device) (*entity.Device, error) {
	return deviceApp.deviceRepository.Create(device)
}

func (deviceApp *deviceApp) Get(devicename string) (*entity.Device, error) {
	return deviceApp.deviceRepository.Get(devicename)
}

func (deviceApp *deviceApp) List(conditions string) ([]entity.Device, error) {
	return deviceApp.deviceRepository.List(conditions)
}

func (deviceApp *deviceApp) Update(devicename string, update *entity.Device) (*entity.Device, error) {
	return deviceApp.deviceRepository.Update(devicename, update)
}

type deviceAppInterface interface {
	Create(device *entity.Device) (*entity.Device, error)
	Get(devicename string) (*entity.Device, error)
	List(conditions string) ([]entity.Device, error)
	Update(devicename string, update *entity.Device) (*entity.Device, error)
}
