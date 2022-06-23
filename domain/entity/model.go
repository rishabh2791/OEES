package entity

import "gorm.io/gorm"

type Model interface {
	BeforeCreate(db *gorm.DB) error
	Validate(action string) error
	Tablename() string
}
