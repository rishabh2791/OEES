package entity

import (
	"oees/domain/value_objects"
	"oees/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PresetDowntime struct {
	value_objects.BaseModel
	ID            string `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	Type          string `json:"type" gorm:"size:50;"`
	Description   string `json:"description" gorm:"size:1000;not null;"`
	DefaultPeriod int16  `json:"default_period"`
}

var _ Model = &PresetDowntime{}

func (downtime *PresetDowntime) BeforeCreate(db *gorm.DB) error {
	downtime.ID = uuid.New().String()
	return nil
}

func (downtime *PresetDowntime) Validate(action string) error {
	errors := map[string]interface{}{}
	if downtime.Type == "" || len(downtime.Type) == 0 {
		errors["type"] = "Downtime Type Missing."
	}
	if downtime.Description == "" || len(downtime.Description) == 0 {
		errors["description"] = "Downtime Description Missing."
	}
	return utilities.ConvertMapToError(errors)
}

func (downtime *PresetDowntime) Tablename() string {
	return "preset_downtimes"
}
