package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID            int64          `gorm:"id,primaryKey"`
	Title         string         `gorm:"title"`
	Slug          string         `gorm:"slug"`
	Author       string         `gorm:"author"`
	FeaturedImage string         `gorm:"featured_image"`
	Content      string         `gorm:"content"`
	PublishedAt  time.Time      `gorm:"published_at"`
	CreatedAt   time.Time      `gorm:"created_at"`
	UpdatedAt *time.Time     `gorm:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
