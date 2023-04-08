package models

import "time"

type Tag struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	NameReading string    `json:"nameReading"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
