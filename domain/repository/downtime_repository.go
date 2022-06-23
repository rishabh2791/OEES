package repository

import "oees/domain/entity"

type DowntimeRepository interface {
	Create(device *entity.Downtime) (*entity.Downtime, error)
	Get(id string) (*entity.Downtime, error)
	List(conditions string) ([]entity.Downtime, error)
	Update(id string, update *entity.Downtime) (*entity.Downtime, error)
}
