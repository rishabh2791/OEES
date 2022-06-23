package repository

import "oees/domain/entity"

type JobRepository interface {
	Create(job *entity.Job) (*entity.Job, error)
	Get(id string) (*entity.Job, error)
	List(conditions string) ([]entity.Job, error)
}
