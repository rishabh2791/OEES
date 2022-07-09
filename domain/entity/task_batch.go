package entity

import (
	"oees/domain/value_objects"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskBatch struct {
	value_objects.BaseModel
	ID                string    `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	TaskID            string    `json:"task_id" gorm:"size:191;not null;"`
	Task              *Task     `json:"task"`
	BatchNumber       string    `json:"batch_number" gorm:"size:20;not null;"`
	StartTime         time.Time `json:"start_time"`
	EndTime           time.Time `json:"end_time"`
	Complete          bool      `json:"complete"`
	CreatedByUsername string    `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User     `json:"created_by"`
	UpdatedByUsername string    `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User     `json:"updated_by"`
}

var _ Model = &TaskBatch{}

func (taskBatch *TaskBatch) BeforeCreate(db *gorm.DB) error {
	taskBatch.ID = uuid.New().String()
	return nil
}

func (taskBatch *TaskBatch) Validate(action string) error {
	return nil
}

func (taskBatch *TaskBatch) Tablename() string {
	return "task_batches"
}
