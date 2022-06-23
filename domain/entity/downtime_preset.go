package entity

import (
	"oees/domain/value_objects"

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
	return nil
}

func (downtime *PresetDowntime) Tablename() string {
	return "preset_downtimes"
}
