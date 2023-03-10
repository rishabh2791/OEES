package repository

import "oees/domain/entity"

type TaskRepository interface {
	Create(task *entity.Task) (*entity.Task, error)
	Get(id string) (*entity.Task, error)
	List(conditions string) ([]entity.Task, error)
	Update(id string, update *entity.Task) (*entity.Task, error)
	Delete(id string) error
}
