package service

import "urlShortener/internal/models"

type URLService interface {
	CreateShortURL(longURL string) (*models.URL, error)
	GetLongURL(shortcode string) (*models.URL, error)
}

type URLGenerator interface {
	Generate() string
}
