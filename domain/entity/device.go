package entity

import (
	"errors"
	"oees/domain/value_objects"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Device struct {
	value_objects.BaseModel
	ID                string `json:"id" gorm:"size:191;not null;unique;primaryKey;"`
	DeviceType        string `json:"device_type" gorm:"size:30;not null;"`
	LineID            string `json:"line_id" gorm:"size:191;not null;uniqueIndex:plant_line_sku;"`
	Line              *Line  `json:"line" gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	Code              string `json:"code" gorm:"size:20;not null;uniqueIndex:plant_line_sku;"`
	Description       string `json:"description" gorm:"size:100;not null;"`
	UseForOEE         bool   `json:"use_for_oee" gorm:"default:false;"`
	CreatedByUsername string `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User  `json:"created_by"`
	UpdatedByUsername string `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User  `json:"updated_by"`
}

var _ Model = &Device{}

func (device *Device) BeforeCreate(db *gorm.DB) error {
	device.ID = uuid.New().String()
	device.DeviceType = strings.ToUpper(device.DeviceType)
	oeeDevices := []Device{}
	db.Where("line_id = ? AND use_for_oee = TRUE", device.LineID).Find(&oeeDevices)
	if len(oeeDevices) == 0 && !device.UseForOEE {
		return errors.New("Each line should have at least one device to measure OEE.")
	}
	if len(oeeDevices) == 1 && device.UseForOEE {
		return errors.New("Each line can have only one device to measure OEE.")
	}
	return nil
}

func (device *Device) Validate(action string) error {
	return nil
}

func (device *Device) Tablename() string {
	return "devices"
}
