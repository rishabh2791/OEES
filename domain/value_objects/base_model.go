package value_objects

import "time"

type BaseModel struct {
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime;"`
}
