package entity

import (
	"oees/domain/value_objects"
	"oees/infrastructure/utilities"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskBatch struct {
	value_objects.BaseModel
	ID                string     `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	TaskID            string     `json:"task_id" gorm:"size:191;not null;uniqueIndex:task_batch;"`
	BatchNumber       string     `json:"batch_number" gorm:"size:20;not null;uniqueIndex:task_batch;"`
	BatchSize         float32    `json:"batch_size" gorm:"default:0;"`
	StartTime         time.Time  `json:"start_time"`
	EndTime           *time.Time `json:"end_time"`
	Complete          bool       `json:"complete"`
	CreatedByUsername string     `json:"created_by_username" gorm:"size:20;not null;"`
	UpdatedByUsername string     `json:"updated_by_username" gorm:"size:20;not null;"`
}

var _ Model = &TaskBatch{}

func (taskBatch *TaskBatch) BeforeCreate(db *gorm.DB) error {
	taskBatch.ID = uuid.New().String()
	return nil
}

func (taskBatch *TaskBatch) Validate(action string) error {
	errors := map[string]interface{}{}
	if taskBatch.TaskID == "" || len(taskBatch.TaskID) == 0 {
		errors["task"] = "Task Missing."
	}
	if taskBatch.BatchNumber == "" || len(taskBatch.BatchNumber) == 0 {
		errors["batch_number"] = "Batch Number Missing."
	}
	if taskBatch.CreatedByUsername == "" || len(taskBatch.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Missing."
	}
	if taskBatch.UpdatedByUsername == "" || len(taskBatch.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Task Missing."
	}
	return utilities.ConvertMapToError(errors)
}

func (taskBatch *TaskBatch) Tablename() string {
	return "task_batches"
}
