package model

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	ID            int64          `gorm:"id,primaryKey"`
	Title         string         `gorm:"title"`
	Content      string         `gorm:"content"`
	CreatedAt   time.Time      `gorm:"created_at"`
	UpdatedAt *time.Time     `gorm:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
