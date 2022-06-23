package entity

import (
	"oees/domain/value_objects"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SKUSpeed struct {
	value_objects.BaseModel
	ID                string  `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	LineID            string  `json:"line_id" gorm:"size:191;not null;"`
	Line              *Line   `json:"line" gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	SKUID             string  `json:"sku_id" gorm:"size:191;not null;column:sku_id;"`
	SKU               *SKU    `json:"sku" gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	Speed             float32 `json:"speed"`
	CreatedByUsername string  `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User   `json:"created_by"`
	UpdatedByUsername string  `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User   `json:"updated_by"`
}

var _ Model = &SKUSpeed{}

func (skuSpeed *SKUSpeed) BeforeCreate(db *gorm.DB) error {
	skuSpeed.ID = uuid.New().String()
	return nil
}

func (skuSpeed *SKUSpeed) Validate(action string) error {
	return nil
}

func (skuSpeed *SKUSpeed) Tablename() string {
	return "sku_speeds"
}
