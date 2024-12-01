package models

import "time"

type URL struct {
	ID int
	OriginalURL string
	ShortCode string
	CreatedAt time.Time
	ExpiresAt time.Time
}