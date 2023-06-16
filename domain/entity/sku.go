package entity

import (
	"oees/domain/value_objects"
	"oees/infrastructure/utilities"

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
	MaxWeight         float32 `json:"max_weight" gorm:"default:1;"`
	ExpectedWeight    float32 `json:"expected_weight" gorm:"default:1;"`
	LowRunSpeed       int     `json:"low_run_speed" gorm:"default:1;"`
	HighRunSpeed      int     `json:"high_run_speed" gorm:"default:1;"`
	CreatedByUsername string  `json:"created_by_username" gorm:"size:20;not null;"`
	UpdatedByUsername string  `json:"updated_by_username" gorm:"size:20;not null;"`
}

var _ Model = &SKU{}

func (sku *SKU) BeforeCreate(db *gorm.DB) error {
	sku.ID = uuid.New().String()
	return nil
}

func (sku *SKU) Validate(action string) error {
	errors := map[string]interface{}{}
	if sku.Code == "" || len(sku.Code) == 0 {
		errors["sku_code"] = "SKU Code Missing"
	}
	if sku.Description == "" || len(sku.Description) == 0 {
		errors["sku_description"] = "SKU Description Missing"
	}
	if sku.CreatedByUsername == "" || len(sku.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Missing"
	}
	if sku.UpdatedByUsername == "" || len(sku.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Missing"
	}
	return utilities.ConvertMapToError(errors)
}

func (sku *SKU) Tablename() string {
	return "skus"
}
