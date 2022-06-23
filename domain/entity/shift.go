package entity

import (
	"oees/domain/value_objects"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Shift struct {
	value_objects.BaseModel
	ID                string `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	PlantCode         string `json:"plant_code" gorm:"size:6;not null;uniqueIndex:plant_shift;"`
	Plant             *Plant `json:"plant"`
	Code              string `json:"code" gorm:"size:2;not null;uniqueIndex:plant_shift;"`
	Description       string `json:"description" gorm:"size:100;not null;"`
	StartTime         string `json:"start_time" gorm:"size:5;not null;"`
	EndTime           string `json:"end_time" gorm:"size:5;not null;"`
	CreatedByUsername string `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User  `json:"created_by"`
	UpdatedByUsername string `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User  `json:"updated_by"`
}

var _ Model = &Shift{}

func (shift *Shift) BeforeCreate(db *gorm.DB) error {
	shift.ID = uuid.New().String()
	return nil
}

func (shift *Shift) Validate(action string) error {
	return nil
}

func (shift *Shift) Tablename() string {
	return "shifts"
}
