package models

import "time"

// BaseModel is used by other models to have default columns
type BaseModel struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
