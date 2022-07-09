package application

import (
	"oees/domain/entity"
	"oees/domain/repository"
)

type taskBatchApp struct {
	taskBatchRepository repository.TaskBatchRepository
}

var _ taskBatchAppInterface = &taskBatchApp{}

func newTaskBatchApp(taskBatchRepository repository.TaskBatchRepository) *taskBatchApp {
	return &taskBatchApp{
		taskBatchRepository: taskBatchRepository,
	}
}

func (taskBatchApp *taskBatchApp) Create(taskBatch *entity.TaskBatch) (*entity.TaskBatch, error) {
	return taskBatchApp.taskBatchRepository.Create(taskBatch)
}

func (taskBatchApp *taskBatchApp) List(taskID string) ([]entity.TaskBatch, error) {
	return taskBatchApp.taskBatchRepository.List(taskID)
}

func (taskBatchApp *taskBatchApp) Update(taskID string, taskBatch *entity.TaskBatch) (*entity.TaskBatch, error) {
	return taskBatchApp.taskBatchRepository.Update(taskID, taskBatch)
}

type taskBatchAppInterface interface {
	Create(taskBatch *entity.TaskBatch) (*entity.TaskBatch, error)
	List(taskID string) ([]entity.TaskBatch, error)
	Update(taskID string, taskBatch *entity.TaskBatch) (*entity.TaskBatch, error)
}
