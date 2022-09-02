package entity

import (
	"oees/domain/value_objects"
	"oees/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole struct {
	value_objects.BaseModel
	ID          string `json:"id" gorm:"size:191;not null;unique;primaryKey;"`
	Description string `json:"description" gorm:"size:200;not null;"`
}

var _ Model = &UserRole{}

func (userRole *UserRole) BeforeCreate(db *gorm.DB) error {
	userRole.ID = uuid.New().String()
	return nil
}

func (userRole *UserRole) Validate(action string) error {
	errors := map[string]interface{}{}
	if userRole.Description == "" || len(userRole.Description) == 0 {
		errors["description"] = "User Role Description Required."
	}
	return utilities.ConvertMapToError(errors)
}

func (userRole *UserRole) Tablename() string {
	return "user_roles"
}
