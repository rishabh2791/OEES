package entity

import (
	"oees/domain/value_objects"
	"oees/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Line struct {
	value_objects.BaseModel
	ID                string `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	Code              string `json:"code" gorm:"size:6;not null;primaryKey;unique;"`
	Name              string `json:"name" gorm:"size:100;not null;"`
	SpeedType         int    `json:"speed_type" gorm:"default:1;"`
	IPAddress         string `json:"ip_address" gorm:"size:16;"`
	CreatedByUsername string `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User  `json:"created_by"`
	UpdatedByUsername string `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User  `json:"updated_by"`
}

var _ Model = &Line{}

func (line *Line) BeforeCreate(db *gorm.DB) error {
	line.ID = uuid.New().String()
	return nil
}

func (line *Line) Validate(action string) error {
	errors := map[string]interface{}{}
	if line.Code == "" || len(line.Code) == 0 {
		errors["code"] = "Line Code Missing."
	}
	if line.Name == "" || len(line.Name) == 0 {
		errors["name"] = "Line Name Missing."
	}
	if line.IPAddress == "" || len(line.IPAddress) == 0 {
		errors["ip_address"] = "Line IP Address Missing."
	}
	if line.CreatedByUsername == "" || len(line.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Missing."
	}
	if line.UpdatedByUsername == "" || len(line.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Missing."
	}
	return utilities.ConvertMapToError(errors)
}

func (line *Line) Tablename() string {
	return "lines"
}
