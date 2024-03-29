package persistance

import (
	"fmt"
	"oees/domain/entity"
	"oees/domain/repository"
	"time"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	getErr := taskRepo.db.
		Preload("Job.SKU").
		Preload("Job.SKU.CreatedBy").
		Preload("Job.SKU.UpdatedBy").
		Preload("Job.SKU.CreatedBy.UserRole").
		Preload("Job.SKU.UpdatedBy.UserRole").
		Preload("Job.CreatedBy").
		Preload("Job.UpdatedBy").
		Preload("Job.CreatedBy.UserRole").
		Preload("Job.UpdatedBy.UserRole").
		Preload("Line.CreatedBy").
		Preload("Line.UpdatedBy").
		Preload("Line.CreatedBy.UserRole").
		Preload("Line.UpdatedBy.UserRole").
		Preload("Shift.CreatedBy").
		Preload("Shift.UpdatedBy").
		Preload("Shift.CreatedBy.UserRole").
		Preload("Shift.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where("id = ?", id).Take(&task).Error

	return &task, getErr
}

func (taskRepo *taskRepo) GetLast(lineID string, taskID string) (*entity.Task, error) {
	currentTask := entity.Task{}

	currentTaskGetErr := taskRepo.db.Preload("Job.SKU").
		Preload("Job.SKU.CreatedBy").
		Preload("Job.SKU.UpdatedBy").
		Preload("Job.SKU.CreatedBy.UserRole").
		Preload("Job.SKU.UpdatedBy.UserRole").
		Preload("Job.CreatedBy").
		Preload("Job.UpdatedBy").
		Preload("Job.CreatedBy.UserRole").
		Preload("Job.UpdatedBy.UserRole").
		Preload("Line.CreatedBy").
		Preload("Line.UpdatedBy").
		Preload("Line.CreatedBy.UserRole").
		Preload("Line.UpdatedBy.UserRole").
		Preload("Shift.CreatedBy").
		Preload("Shift.UpdatedBy").
		Preload("Shift.CreatedBy.UserRole").
		Preload("Shift.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where("id = ?", taskID).Take(&currentTask).Error
	if currentTaskGetErr != nil {
		return nil, currentTaskGetErr
	}

	currentTaskStartTime := time.Now()

	if currentTask.StartTime != nil {
		currentTaskStartTime = *currentTask.StartTime
	}

	tasks := []entity.Task{}

	queryString := fmt.Sprintf("line_id = '%s' AND start_time <= '%s'", lineID, currentTaskStartTime)
	getErr := taskRepo.db.Preload("Job.SKU").
		Preload("Job.SKU.CreatedBy").
		Preload("Job.SKU.UpdatedBy").
		Preload("Job.SKU.CreatedBy.UserRole").
		Preload("Job.SKU.UpdatedBy.UserRole").
		Preload("Job.CreatedBy").
		Preload("Job.UpdatedBy").
		Preload("Job.CreatedBy.UserRole").
		Preload("Job.UpdatedBy.UserRole").
		Preload("Line.CreatedBy").
		Preload("Line.UpdatedBy").
		Preload("Line.CreatedBy.UserRole").
		Preload("Line.UpdatedBy.UserRole").
		Preload("Shift.CreatedBy").
		Preload("Shift.UpdatedBy").
		Preload("Shift.CreatedBy.UserRole").
		Preload("Shift.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where(queryString).Order("start_time desc").Limit(2).Find(&tasks).Error

	if len(tasks) > 1 {
		return &tasks[1], getErr
	}
	return &tasks[0], nil
}

func (taskRepo *taskRepo) List(conditions string) ([]entity.Task, error) {
	tasks := []entity.Task{}
	getErr := taskRepo.db.
		Preload("Job.SKU").
		Preload("Job.SKU.CreatedBy").
		Preload("Job.SKU.UpdatedBy").
		Preload("Job.SKU.CreatedBy.UserRole").
		Preload("Job.SKU.UpdatedBy.UserRole").
		Preload("Job.CreatedBy").
		Preload("Job.UpdatedBy").
		Preload("Job.CreatedBy.UserRole").
		Preload("Job.UpdatedBy.UserRole").
		Preload("Line.CreatedBy").
		Preload("Line.UpdatedBy").
		Preload("Line.CreatedBy.UserRole").
		Preload("Line.UpdatedBy.UserRole").
		Preload("Shift.CreatedBy").
		Preload("Shift.UpdatedBy").
		Preload("Shift.CreatedBy.UserRole").
		Preload("Shift.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where(conditions).Find(&tasks).Error
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
		Preload("Job.SKU").
		Preload("Job.SKU.CreatedBy").
		Preload("Job.SKU.UpdatedBy").
		Preload("Job.SKU.CreatedBy.UserRole").
		Preload("Job.SKU.UpdatedBy.UserRole").
		Preload("Job.CreatedBy").
		Preload("Job.UpdatedBy").
		Preload("Job.CreatedBy.UserRole").
		Preload("Job.UpdatedBy.UserRole").
		Preload("Line.CreatedBy").
		Preload("Line.UpdatedBy").
		Preload("Line.CreatedBy.UserRole").
		Preload("Line.UpdatedBy.UserRole").
		Preload("Shift.CreatedBy").
		Preload("Shift.UpdatedBy").
		Preload("Shift.CreatedBy.UserRole").
		Preload("Shift.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).
		Where("id = ?", id).Take(&updated)

	return &updated, nil
}

func (taskRepo *taskRepo) Delete(id string) error {
	deletionErr := taskRepo.db.Where("id = ?", id).Delete(&entity.Task{}).Error
	return deletionErr
}
