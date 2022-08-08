package value_objects

type TotalDeviceData struct {
	Value    float32 `json:"value"`
	DeviceID string  `json:"device_id" gorm:"size:191;not null;"`
}
