package entity

import (
	"oees/domain/value_objects"
	"oees/infrastructure/utilities"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	value_objects.BaseModel
	ID                string     `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	JobID             string     `json:"job_id" gorm:"size:191;not null;unique;"`
	LineID            string     `json:"line_id" gorm:"size:191;not null;"`
	ScheduledDate     *time.Time `json:"scheduled_date"`
	ShiftID           string     `json:"shift_id" gorm:"size:191;not null;"`
	StartTime         *time.Time `json:"start_time"`
	EndTime           *time.Time `json:"end_time"`
	Complete          bool       `json:"complete" gorm:"default:false;"`
	Plan              int16      `json:"plan" gorm:"default:1;"`
	Actual            int16      `json:"actual" gorm:"default:0;"`
	CreatedByUsername string     `json:"created_by_username" gorm:"size:20;not null;"`
	UpdatedByUsername string     `json:"updated_by_username" gorm:"size:20;not null;"`
}

var _ Model = &Task{}

func (task *Task) BeforeCreate(db *gorm.DB) error {
	task.ID = uuid.New().String()
	return nil
}

func (task *Task) Validate(action string) error {
	errors := map[string]interface{}{}
	if task.JobID == "" || len(task.JobID) == 0 {
		errors["job"] = "Job Missing"
	}
	if task.LineID == "" || len(task.LineID) == 0 {
		errors["line"] = "Line Missing"
	}
	if task.ShiftID == "" || len(task.ShiftID) == 0 {
		errors["shift"] = "Shift ID Missing"
	}
	if task.CreatedByUsername == "" || len(task.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Missing"
	}
	if task.UpdatedByUsername == "" || len(task.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Missing"
	}
	return utilities.ConvertMapToError(errors)
}

func (task *Task) Tablename() string {
	return "tasks"
}
