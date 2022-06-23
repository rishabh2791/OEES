package entity

import (
	"oees/domain/value_objects"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Job struct {
	value_objects.BaseModel
	ID                string `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	Code              string `json:"code" gorm:"size:10;not null;"`
	PlantCode         string `json:"plant_code" gorm:"size:191;not null;"`
	Plant             *Plant `json:"plant"`
	SKUID             string `json:"sku_id" gorm:"size:191;not null;column:sku_id;"`
	SKU               *SKU   `json:"sku" gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	Plan              int16  `json:"plan" gorm:"default:1;"`
	CreatedByUsername string `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User  `json:"created_by"`
	UpdatedByUsername string `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User  `json:"updated_by"`
}

var _ Model = &Job{}

func (job *Job) BeforeCreate(db *gorm.DB) error {
	job.ID = uuid.New().String()
	return nil
}

func (job *Job) Validate(action string) error {
	return nil
}

func (job *Job) Tablename() string {
	return "tasks"
}
