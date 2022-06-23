package entity

import (
	"oees/domain/value_objects"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SKU struct {
	value_objects.BaseModel
	ID                string `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	PlantCode         string `json:"plant_code" gorm:"size:6;not null;uniqueIndex:plant_sku;"`
	Plant             *Plant `json:"plant" gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	Code              string `json:"code" gorm:"size:20;not null;uniqueIndex:plant_sku;"`
	Description       string `json:"description" gorm:"size:100;not null;"`
	CreatedByUsername string `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User  `json:"created_by"`
	UpdatedByUsername string `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User  `json:"updated_by"`
}

var _ Model = &SKU{}

func (sku *SKU) BeforeCreate(db *gorm.DB) error {
	sku.ID = uuid.New().String()
	return nil
}

func (sku *SKU) Validate(action string) error {
	return nil
}

func (sku *SKU) Tablename() string {
	return "skus"
}
