package entity

import (
	"oees/domain/value_objects"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Line struct {
	value_objects.BaseModel
	ID                string `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	Code              string `json:"code" gorm:"size:10;not null;uniqueIndex:plant_line;"`
	Name              string `json:"name" gorm:"size:100;not null;"`
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
	return nil
}

func (line *Line) Tablename() string {
	return "lines"
}
