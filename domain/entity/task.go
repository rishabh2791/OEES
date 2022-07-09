package entity

import (
	"oees/domain/value_objects"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	value_objects.BaseModel
	ID                string     `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	JobID             string     `json:"job_id" gorm:"size:191;not null;uniqueIndex:job_line;"`
	Job               *Job       `json:"job"`
	LineID            string     `json:"line_id" gorm:"size:191;not null;uniqueIndex:job_line;"`
	Line              *Line      `json:"line" gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	ScheduledDate     *time.Time `json:"scheduled_date"`
	ShiftID           string     `json:"shift_id" gorm:"size:191;not null;"`
	Shift             *Shift     `json:"shift"`
	StartTime         *time.Time `json:"start_time"`
	EndTime           *time.Time `json:"end_time"`
	Complete          bool       `json:"complete" gorm:"default:false;"`
	Plan              int16      `json:"plan" gorm:"default:1;"`
	Actual            int16      `json:"actual" gorm:"default:0;"`
	CreatedByUsername string     `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User      `json:"created_by"`
	UpdatedByUsername string     `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User      `json:"updated_by"`
}

var _ Model = &Task{}

func (task *Task) BeforeCreate(db *gorm.DB) error {
	task.ID = uuid.New().String()
	return nil
}

func (task *Task) Validate(action string) error {
	return nil
}

func (task *Task) Tablename() string {
	return "tasks"
}
