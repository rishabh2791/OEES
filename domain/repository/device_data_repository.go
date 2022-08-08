package repository

import (
	"oees/domain/entity"
	"oees/domain/value_objects"
)

type DeviceDataRepository interface {
	Create(deviceData *entity.DeviceData) (*entity.DeviceData, error)
	List(conditions string) ([]entity.DeviceData, error)
	TotalDeviceData(conditions string) (*value_objects.TotalDeviceData, error)
}
