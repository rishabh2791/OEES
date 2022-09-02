package entity

import (
	"oees/domain/value_objects"
	"oees/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Shift struct {
	value_objects.BaseModel
	ID                string `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	Code              string `json:"code" gorm:"size:2;not null;primaryKey;unique;"`
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
	errors := map[string]interface{}{}
	if shift.Code == "" || len(shift.Code) == 0 {
		errors["code"] = "Shift Code Missing"
	}
	if shift.Description == "" || len(shift.Description) == 0 {
		errors["description"] = "Shift Description Missing"
	}
	if shift.StartTime == "" || len(shift.StartTime) == 0 {
		errors["start_time"] = "Shift Start Time Missing"
	}
	if shift.EndTime == "" || len(shift.EndTime) == 0 {
		errors["end_time"] = "Shift End Time Missing"
	}
	if shift.CreatedByUsername == "" || len(shift.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Missing"
	}
	if shift.UpdatedByUsername == "" || len(shift.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Missing"
	}
	return utilities.ConvertMapToError(errors)
}

func (shift *Shift) Tablename() string {
	return "shifts"
}
