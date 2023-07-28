package entity

import (
	"oees/domain/value_objects"
	"oees/infrastructure/utilities"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Downtime struct {
	value_objects.BaseModel
	ID                string     `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	LineID            string     `json:"line_id" gorm:"size:191;not null;uniqueIndex:unique_dt_line_start_end;uniqueIndex:unique_dt_line_start;"`
	Line              *Line      `json:"line"`
	Planned           bool       `json:"planned" gorm:"default:false;"`
	Controlled        bool       `json:"controlled" gorm:"default:false;"`
	StartTime         *time.Time `json:"start_time" gorm:"uniqueIndex:unique_dt_line_start_end;uniqueIndex:unique_dt_line_start;"`
	EndTime           *time.Time `json:"end_time" gorm:"uniqueIndex:unique_dt_line_start_end;"`
	Preset            string     `json:"preset" gorm:"size:1000;"`
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
	errors := map[string]interface{}{}
	if downtime.LineID == "" || len(downtime.LineID) == 0 {
		errors["line"] = "Line Details Missing."
	}
	if downtime.UpdatedByUsername == "" || len(downtime.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Missing."
	}
	return utilities.ConvertMapToError(errors)
}

func (downtime *Downtime) Tablename() string {
	return "downtimes"
}
