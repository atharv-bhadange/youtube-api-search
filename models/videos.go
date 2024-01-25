package models

import "time"

// VideoDetails struct represents a video
type VideoDetails struct {
	Id          string    `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description"`
	PublishedAt time.Time `gorm:"not null" json:"published_at"`
	Thumbnail   string    `gorm:"not null" json:"thumbnail"`
	Query       string    `gorm:"not null" json:"query"`
}

type VideoResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []VideoDetails `json:"data"`
}
