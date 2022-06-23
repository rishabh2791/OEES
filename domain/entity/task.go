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
	Code              string     `json:"code" gorm:"size:10;not null;"`
	SKUID             string     `json:"sku_id" gorm:"size:191;not null;column:sku_id;"`
	SKU               *SKU       `json:"sku" gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	LineID            string     `json:"line_id" gorm:"size:191;not null;"`
	Line              *Line      `json:"line" gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	StartTime         *time.Time `json:"start_time"`
	EndTime           *time.Time `json:"end_time"`
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
