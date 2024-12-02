package models

import "time"

type URL struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	OriginalURL string    `json:"original_url" bson:"original_url" validate:"required,url"`
	ShortCode   string    `json:"short_code" bson:"short_code" validate:"required"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	ExpiresAt   time.Time `json:"expires_at,omitempty" bson:"expires_at,omitempty"`
}

type CreateURLRequest struct {
	URL string `json:"url"`
}

type CreateURLResponse struct {
	ShortURL string `json:"short_url"`
}
