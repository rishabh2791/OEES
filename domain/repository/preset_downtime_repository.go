package repository

import "oees/domain/entity"

type PresetDowntimeRepository interface {
	Create(downtime *entity.PresetDowntime) (*entity.PresetDowntime, error)
	List(conditions string) ([]entity.PresetDowntime, error)
}
