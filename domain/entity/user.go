package entity

import (
	"oees/domain/value_objects"

	"gorm.io/gorm"
)

type User struct {
	value_objects.BaseModel
	UserRoleID string    `json:"user_role_id" gorm:"size:191;not null;"`
	UserRole   *UserRole `json:"user_role"`
	Username   string    `json:"username" gorm:"size:20;not null;unique;primaryKey;"`
	FirstName  string    `json:"first_name" gorm:"size:100;not null;"`
	LastName   string    `json:"last_name" gorm:"size:100;"`
	Password   string    `json:"password" gorm:"size:200;not null;"`
	Email      string    `json:"email" gorm:"size:100;not null;"`
	ProfilePic string    `json:"profile_pic" gorm:"size:1000;not null;default:'public/profile_pics/default.jpg'"`
	Active     bool      `json:"active" gorm:"default:true;"`
	SecretKey  string    `json:"secret_key" gorm:"size:191;"`
}

var _ Model = &User{}

func (user *User) BeforeCreate(db *gorm.DB) error {
	return nil
}

func (user *User) Validate(action string) error {
	return nil
}

func (user *User) Tablename() string {
	return "users"
}
