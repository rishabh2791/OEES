package persistance

import (
	"oees/domain/entity"
	"oees/domain/repository"

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
		Preload("SKU.Plant").
		Preload("SKU.Plant.CreatedBy").
		Preload("SKU.Plant.CreatedBy.UserRole").
		Preload("SKU.Plant.UpdatedBy").
		Preload("SKU.Plant.UpdatedBy.UserRole").
		Preload("SKU.CreatedBy").
		Preload("SKU.UpdatedBy").
		Preload("SKU.CreatedBy.UserRole").
		Preload("SKU.UpdatedBy.UserRole").
		Preload("Line.Plant").
		Preload("Line.Plant.CreatedBy").
		Preload("Line.Plant.CreatedBy.UserRole").
		Preload("Line.Plant.UpdatedBy").
		Preload("Line.Plant.UpdatedBy.UserRole").
		Preload("Line.CreatedBy").
		Preload("Line.UpdatedBy").
		Preload("Line.CreatedBy.UserRole").
		Preload("Line.UpdatedBy.UserRole").
		Preload("Shift.Plant").
		Preload("Shift.Plant.CreatedBy").
		Preload("Shift.Plant.CreatedBy.UserRole").
		Preload("Shift.Plant.UpdatedBy").
		Preload("Shift.Plant.UpdatedBy.UserRole").
		Preload("Shift.CreatedBy").
		Preload("Shift.UpdatedBy").
		Preload("Shift.CreatedBy.UserRole").
		Preload("Shift.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where("id = ?", id).Take(&task).Error
	return &task, getErr
}

func (taskRepo *taskRepo) List(conditions string) ([]entity.Task, error) {
	tasks := []entity.Task{}
	getErr := taskRepo.db.
		Preload("SKU.Plant").
		Preload("SKU.Plant.CreatedBy").
		Preload("SKU.Plant.CreatedBy.UserRole").
		Preload("SKU.Plant.UpdatedBy").
		Preload("SKU.Plant.UpdatedBy.UserRole").
		Preload("SKU.CreatedBy").
		Preload("SKU.UpdatedBy").
		Preload("SKU.CreatedBy.UserRole").
		Preload("SKU.UpdatedBy.UserRole").
		Preload("Line.Plant").
		Preload("Line.Plant.CreatedBy").
		Preload("Line.Plant.CreatedBy.UserRole").
		Preload("Line.Plant.UpdatedBy").
		Preload("Line.Plant.UpdatedBy.UserRole").
		Preload("Line.CreatedBy").
		Preload("Line.UpdatedBy").
		Preload("Line.CreatedBy.UserRole").
		Preload("Line.UpdatedBy.UserRole").
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
		Preload("SKU.Plant").
		Preload("SKU.Plant.CreatedBy").
		Preload("SKU.Plant.CreatedBy.UserRole").
		Preload("SKU.Plant.UpdatedBy").
		Preload("SKU.Plant.UpdatedBy.UserRole").
		Preload("SKU.CreatedBy").
		Preload("SKU.UpdatedBy").
		Preload("SKU.CreatedBy.UserRole").
		Preload("SKU.UpdatedBy.UserRole").
		Preload("Line.Plant").
		Preload("Line.Plant.CreatedBy").
		Preload("Line.Plant.CreatedBy.UserRole").
		Preload("Line.Plant.UpdatedBy").
		Preload("Line.Plant.UpdatedBy.UserRole").
		Preload("Line.CreatedBy").
		Preload("Line.UpdatedBy").
		Preload("Line.CreatedBy.UserRole").
		Preload("Line.UpdatedBy.UserRole").
		Preload("Shift.Plant").
		Preload("Shift.Plant.CreatedBy").
		Preload("Shift.Plant.CreatedBy.UserRole").
		Preload("Shift.Plant.UpdatedBy").
		Preload("Shift.Plant.UpdatedBy.UserRole").
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
