package repository

import "oees/domain/entity"

type LineRepository interface {
	Create(device *entity.Line) (*entity.Line, error)
	Get(id string) (*entity.Line, error)
	List(conditions string) ([]entity.Line, error)
	Update(id string, update *entity.Line) (*entity.Line, error)
}
