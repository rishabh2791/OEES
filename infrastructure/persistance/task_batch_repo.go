package persistance

import (
	"oees/domain/entity"
	"oees/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type taskBatchRepo struct {
	db     *gorm.DB
	logger hclog.Logger
}

var _ repository.TaskBatchRepository = &taskBatchRepo{}

func newTaskBatchRepo(db *gorm.DB, logger hclog.Logger) *taskBatchRepo {
	return &taskBatchRepo{
		db:     db,
		logger: logger,
	}
}

func (taskBatchRepo *taskBatchRepo) Create(taskBatch *entity.TaskBatch) (*entity.TaskBatch, error) {
	taskID := taskBatch.TaskID
	openTaskBatch := entity.TaskBatch{}

	openTaskBatchErr := taskBatchRepo.db.Where("task_id = ? AND complete = 0", taskID).Take(&openTaskBatch).Error
	if openTaskBatchErr != nil {
		update := entity.TaskBatch{}
		update.Complete = true
		taskBatchRepo.db.Where("id = ?", openTaskBatch.ID).Updates(update)
	}

	creationErr := taskBatchRepo.db.Create(&taskBatch).Error

	createdTaskBatch := entity.TaskBatch{}
	taskBatchRepo.db.
		Preload("Task.Job.SKU").
		Preload("Task.Job.SKU.CreatedBy").
		Preload("Task.Job.SKU.UpdatedBy").
		Preload("Task.Job.SKU.CreatedBy.UserRole").
		Preload("Task.Job.SKU.UpdatedBy.UserRole").
		Preload("Task.Job.CreatedBy").
		Preload("Task.Job.UpdatedBy").
		Preload("Task.Job.CreatedBy.UserRole").
		Preload("Task.Job.UpdatedBy.UserRole").
		Preload("Task.Line.CreatedBy").
		Preload("Task.Line.UpdatedBy").
		Preload("Task.Line.CreatedBy.UserRole").
		Preload("Task.Line.UpdatedBy.UserRole").
		Preload("Task.Shift.CreatedBy").
		Preload("Task.Shift.UpdatedBy").
		Preload("Task.Shift.CreatedBy.UserRole").
		Preload("Task.Shift.UpdatedBy.UserRole").
		Preload("Task.CreatedBy.UserRole").
		Preload("Task.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where("task_id = ?", taskID).Find(&createdTaskBatch)

	return &createdTaskBatch, creationErr
}

func (taskBatchRepo *taskBatchRepo) List(taskID string) ([]entity.TaskBatch, error) {
	taskBatches := []entity.TaskBatch{}
	getErr := taskBatchRepo.db.
		Preload("Task.Job.SKU").
		Preload("Task.Job.SKU.CreatedBy").
		Preload("Task.Job.SKU.UpdatedBy").
		Preload("Task.Job.SKU.CreatedBy.UserRole").
		Preload("Task.Job.SKU.UpdatedBy.UserRole").
		Preload("Task.Job.CreatedBy").
		Preload("Task.Job.UpdatedBy").
		Preload("Task.Job.CreatedBy.UserRole").
		Preload("Task.Job.UpdatedBy.UserRole").
		Preload("Task.Line.CreatedBy").
		Preload("Task.Line.UpdatedBy").
		Preload("Task.Line.CreatedBy.UserRole").
		Preload("Task.Line.UpdatedBy.UserRole").
		Preload("Task.Shift.CreatedBy").
		Preload("Task.Shift.UpdatedBy").
		Preload("Task.Shift.CreatedBy.UserRole").
		Preload("Task.Shift.UpdatedBy.UserRole").
		Preload("Task.CreatedBy.UserRole").
		Preload("Task.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where("task_id = ?", taskID).Find(&taskBatches).Error
	return taskBatches, getErr
}

func (taskBatchRepo *taskBatchRepo) Update(taskID string, update *entity.TaskBatch) (*entity.TaskBatch, error) {
	existingTaskBatch := entity.TaskBatch{}

	getErr := taskBatchRepo.db.
		Preload("Task.Job.SKU").
		Preload("Task.Job.SKU.CreatedBy").
		Preload("Task.Job.SKU.UpdatedBy").
		Preload("Task.Job.SKU.CreatedBy.UserRole").
		Preload("Task.Job.SKU.UpdatedBy.UserRole").
		Preload("Task.Job.CreatedBy").
		Preload("Task.Job.UpdatedBy").
		Preload("Task.Job.CreatedBy.UserRole").
		Preload("Task.Job.UpdatedBy.UserRole").
		Preload("Task.Line.CreatedBy").
		Preload("Task.Line.UpdatedBy").
		Preload("Task.Line.CreatedBy.UserRole").
		Preload("Task.Line.UpdatedBy.UserRole").
		Preload("Task.Shift.CreatedBy").
		Preload("Task.Shift.UpdatedBy").
		Preload("Task.Shift.CreatedBy.UserRole").
		Preload("Task.Shift.UpdatedBy.UserRole").
		Preload("Task.CreatedBy.UserRole").
		Preload("Task.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where("id = ?", taskID).Find(&existingTaskBatch).Error

	if getErr != nil {
		return nil, getErr
	}

	updationErr := taskBatchRepo.db.Table(existingTaskBatch.Tablename()).Where("id = ?", taskID).Updates(update).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.TaskBatch{}
	taskBatchRepo.db.
		Preload("Task.Job.SKU").
		Preload("Task.Job.SKU.CreatedBy").
		Preload("Task.Job.SKU.UpdatedBy").
		Preload("Task.Job.SKU.CreatedBy.UserRole").
		Preload("Task.Job.SKU.UpdatedBy.UserRole").
		Preload("Task.Job.CreatedBy").
		Preload("Task.Job.UpdatedBy").
		Preload("Task.Job.CreatedBy.UserRole").
		Preload("Task.Job.UpdatedBy.UserRole").
		Preload("Task.Line.CreatedBy").
		Preload("Task.Line.UpdatedBy").
		Preload("Task.Line.CreatedBy.UserRole").
		Preload("Task.Line.UpdatedBy.UserRole").
		Preload("Task.Shift.CreatedBy").
		Preload("Task.Shift.UpdatedBy").
		Preload("Task.Shift.CreatedBy.UserRole").
		Preload("Task.Shift.UpdatedBy.UserRole").
		Preload("Task.CreatedBy.UserRole").
		Preload("Task.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where("id = ?", taskID).Take(&updated)

	return &updated, nil
}
