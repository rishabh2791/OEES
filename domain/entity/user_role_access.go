package entity

import (
	"oees/domain/value_objects"
	"oees/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRoleAccess struct {
	value_objects.BaseModel
	ID                string    `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	UserRoleID        string    `json:"user_role_id" gorm:"size:191;not null;"`
	UserRole          *UserRole `json:"user_role"`
	Table             string    `json:"tablename" gorm:"size:100;not null;"`
	AccessCode        string    `json:"access_code" gorm:"size:4;default:'0000';"`
	CreatedByUsername string    `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User     `json:"created_by"`
	UpdatedByUsername string    `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User     `json:"updated_by"`
}

var _ Model = &UserRoleAccess{}

func (userRoleAccess *UserRoleAccess) BeforeCreate(db *gorm.DB) error {
	userRoleAccess.ID = uuid.New().String()
	return nil
}

func (userRoleAccess *UserRoleAccess) Validate(action string) error {
	errors := map[string]interface{}{}
	if userRoleAccess.UserRoleID == "" || len(userRoleAccess.UserRoleID) == 0 {
		errors["role"] = "User Role Missing"
	}
	if userRoleAccess.Table == "" || len(userRoleAccess.Table) == 0 {
		errors["table"] = "Table Missing"
	}
	return utilities.ConvertMapToError(errors)
}

func (userRoleAccess *UserRoleAccess) Tablename() string {
	return "user_role_accesses"
}
