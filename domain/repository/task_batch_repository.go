package repository

import "oees/domain/entity"

type TaskBatchRepository interface {
	Create(taskBatch *entity.TaskBatch) (*entity.TaskBatch, error)
	List(taskID string) ([]entity.TaskBatch, error)
	Update(taskID string, taskBatch *entity.TaskBatch) (*entity.TaskBatch, error)
}
