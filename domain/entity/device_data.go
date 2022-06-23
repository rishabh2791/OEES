package entity

import (
	"oees/domain/value_objects"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeviceData struct {
	value_objects.BaseModel
	ID               string  `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	DeviceID         string  `json:"device_id" gorm:"size:191;not null;"`
	Device           *Device `json:"device" gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	TheoreticalValue float32 `json:"theoretical_value" gorm:"default:0.0;"`
	Value            float32 `json:"value" gorm:"default:0.0;"`
}

var _ Model = &DeviceData{}

func (deviceData *DeviceData) BeforeCreate(db *gorm.DB) error {
	deviceData.ID = uuid.New().String()
	return nil
}

func (deviceData *DeviceData) Validate(action string) error {
	return nil
}

func (deviceData *DeviceData) Tablename() string {
	return "device_data"
}
