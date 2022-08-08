package entity

import (
	"oees/domain/value_objects"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SKU struct {
	value_objects.BaseModel
	ID                string  `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	Code              string  `json:"code" gorm:"size:20;not null;"`
	Description       string  `json:"description" gorm:"size:100;not null;"`
	CaseLot           float32 `json:"case_lot" gorm:"default:1;"`
	MinWeight         float32 `json:"min_weight" gorm:"default:1;"`
	ExpectedWeight    float32 `json:"expected_weight" gorm:"default:1;"`
	LowRunSpeed       int     `json:"low_run_speed" gorm:"default:1;"`
	HighRunSpeed      int     `json:"high_run_speed" gorm:"default:1;"`
	CreatedByUsername string  `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User   `json:"created_by"`
	UpdatedByUsername string  `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User   `json:"updated_by"`
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
