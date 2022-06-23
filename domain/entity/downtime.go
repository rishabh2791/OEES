package entity

import (
	"oees/domain/value_objects"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Downtime struct {
	value_objects.BaseModel
	ID                string     `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	LineID            string     `json:"line_id" gorm:"size:191;not null;"`
	Line              *Line      `json:"line"`
	Planned           bool       `json:"planned" gorm:"default:false;"`
	Controlled        bool       `json:"controlled" gorm:"default:false;"`
	StartTime         *time.Time `json:"start_time"`
	EndTime           *time.Time `json:"end_time"`
	Description       string     `json:"description" gorm:"size:1000;"`
	UpdatedByUsername string     `json:"updated_by_username" gorm:"size:20;"`
	UpdatedBy         *User      `json:"updated_by"`
}

var _ Model = &Downtime{}

func (downtime *Downtime) BeforeCreate(db *gorm.DB) error {
	downtime.ID = uuid.New().String()
	return nil
}

func (downtime *Downtime) Validate(action string) error {
	return nil
}

func (downtime *Downtime) Tablename() string {
	return "downtimes"
}
