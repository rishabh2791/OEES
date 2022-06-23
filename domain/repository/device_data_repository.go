package repository

import "oees/domain/entity"

type DeviceDataRepository interface {
	Create(deviceData *entity.DeviceData) (*entity.DeviceData, error)
	List(conditions string) ([]entity.DeviceData, error)
}
