package repository

import (
	"urlShortener/internal/models"
)

type URLRepository interface {
	FindByLongUrl(longURL string) (*models.URL, error)
	SaveURL(url *models.URL) error
	FindByShortCode(code string) (*models.URL, error)
}
