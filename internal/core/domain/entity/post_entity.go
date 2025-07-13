package entity

import "time"

type PostEntity struct {
	ID            int64
	Title         string
	Slug          string
	Author       string
	FeaturedImage string
	Content      string
	PublishedAt  time.Time
}
