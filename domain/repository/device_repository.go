package repository

import "oees/domain/entity"

type DeviceRepository interface {
	Create(device *entity.Device) (*entity.Device, error)
	Get(id string) (*entity.Device, error)
	List(conditions string) ([]entity.Device, error)
	Update(id string, update *entity.Device) (*entity.Device, error)
}
