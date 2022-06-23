package entity

import (
	"oees/domain/value_objects"

	"gorm.io/gorm"
)

type Plant struct {
	value_objects.BaseModel
	Code              string `json:"code" gorm:"size:6;not null;unique;primaryKey;"`
	Description       string `json:"description" gorm:"size:100;not null;"`
	CreatedByUsername string `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User  `json:"created_by"`
	UpdatedByUsername string `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User  `json:"updated_by"`
}

var _ Model = &Plant{}

func (plant *Plant) BeforeCreate(db *gorm.DB) error {
	return nil
}

func (plant *Plant) Validate(action string) error {
	return nil
}

func (plant *Plant) Tablename() string {
	return "plants"
}
