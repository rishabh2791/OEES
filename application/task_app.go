package application

import (
	"oees/domain/entity"
	"oees/domain/repository"
)

type taskApp struct {
	taskRepository repository.TaskRepository
}

var _ taskAppInterface = &taskApp{}

func newTaskApp(taskRepository repository.TaskRepository) *taskApp {
	return &taskApp{
		taskRepository: taskRepository,
	}
}

type taskAppInterface interface {
	Create(task *entity.Task) (*entity.Task, error)
	Get(id string) (*entity.Task, error)
	GetLast(lineID string, taskID string) (*entity.Task, error)
	List(conditions string) ([]entity.Task, error)
	Update(id string, update *entity.Task) (*entity.Task, error)
	Delete(id string) error
}

func (taskApp *taskApp) Create(task *entity.Task) (*entity.Task, error) {
	return taskApp.taskRepository.Create(task)
}

func (taskApp *taskApp) Get(id string) (*entity.Task, error) {
	return taskApp.taskRepository.Get(id)
}

func (taskApp *taskApp) GetLast(lineID string, taskID string) (*entity.Task, error) {
	return taskApp.taskRepository.GetLast(lineID, taskID)
}

func (taskApp *taskApp) List(conditions string) ([]entity.Task, error) {
	return taskApp.taskRepository.List(conditions)
}

func (taskApp *taskApp) Update(id string, update *entity.Task) (*entity.Task, error) {
	return taskApp.taskRepository.Update(id, update)
}

func (taskApp *taskApp) Delete(id string) error {
	return taskApp.taskRepository.Delete(id)
}
