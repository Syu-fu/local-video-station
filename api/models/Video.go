package models

import "time"

type Video struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	TitleReading string    `json:"title_reading"`
	Url          string    `json:"url"`
	ThumbnailUrl string    `json:"thumbnail_url"`
	Tags         []Tag     `json:"tags"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}
