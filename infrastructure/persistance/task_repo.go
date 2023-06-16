package persistance

import (
	"oees/domain/entity"
	"oees/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
)

type taskRepo struct {
	db     *gorm.DB
	logger hclog.Logger
}

var _ repository.TaskRepository = &taskRepo{}

func NewTaskRepo(db *gorm.DB, logger hclog.Logger) *taskRepo {
	return &taskRepo{
		db:     db,
		logger: logger,
	}
}

func (taskRepo *taskRepo) Create(task *entity.Task) (*entity.Task, error) {
	validationErr := task.Validate("create")
	if validationErr != nil {
		return nil, validationErr
	}
	creationErr := taskRepo.db.Create(&task).Error
	return task, creationErr
}

func (taskRepo *taskRepo) Get(id string) (*entity.Task, error) {
	task := entity.Task{}
	getErr := taskRepo.db.Where("id = ?", id).Take(&task).Error
	return &task, getErr
}

func (taskRepo *taskRepo) List(conditions string) ([]entity.Task, error) {
	tasks := []entity.Task{}
	getErr := taskRepo.db.Where(conditions).Find(&tasks).Error
	return tasks, getErr
}

func (taskRepo *taskRepo) Update(id string, update *entity.Task) (*entity.Task, error) {
	existingTask := entity.Task{}

	err := taskRepo.db.Where("id = ?", id).Take(&existingTask).Error
	if err != nil {
		return nil, err
	}

	updationErr := taskRepo.db.Table(existingTask.Tablename()).Where("id = ?", id).Updates(update).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.Task{}
	taskRepo.db.
		Where("id = ?", id).Take(&updated)

	return &updated, nil
}

func (taskRepo *taskRepo) Delete(id string) error {
	deletionErr := taskRepo.db.Where("id = ?", id).Delete(&entity.Task{}).Error
	return deletionErr
}
