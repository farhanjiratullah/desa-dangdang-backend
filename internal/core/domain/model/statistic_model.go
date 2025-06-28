package model

import (
	"time"

	"gorm.io/gorm"
)

type Statistic struct {
	ID        int64          `gorm:"id,primaryKey"`
	Name      string         `gorm:"name"`
	Total     int64          `gorm:"total"`
	Icon      string         `gorm:"icon"`
	CreatedAt time.Time      `gorm:"created_at"`
	UpdatedAt *time.Time     `gorm:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
