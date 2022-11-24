package entity

import (
	"oees/domain/value_objects"
	"oees/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Job struct {
	value_objects.BaseModel
	ID                string  `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	Code              string  `json:"code" gorm:"size:10;not null;unique;primaryKey;"`
	SKUID             string  `json:"sku_id" gorm:"size:191;not null;column:sku_id;"`
	SKU               *SKU    `json:"sku"`
	Plan              float32 `json:"plan" gorm:"default:1;"`
	CreatedByUsername string  `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User   `json:"created_by"`
	UpdatedByUsername string  `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User   `json:"updated_by"`
}

var _ Model = &Job{}

func (job *Job) BeforeCreate(db *gorm.DB) error {
	job.ID = uuid.New().String()
	return nil
}

func (job *Job) Validate(action string) error {
	errors := map[string]interface{}{}
	if job.Code == "" || len(job.Code) == 0 {
		errors["code"] = "Job Code Missing."
	}
	if job.SKUID == "" || len(job.SKUID) == 0 {
		errors["sku"] = "SKU Missing."
	}
	if job.CreatedByUsername == "" || len(job.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Missing."
	}
	if job.UpdatedByUsername == "" || len(job.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Missing."
	}
	return utilities.ConvertMapToError(errors)
}

func (job *Job) Tablename() string {
	return "jobs"
}
