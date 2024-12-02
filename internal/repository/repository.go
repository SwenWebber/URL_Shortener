package repository

import (
	"urlShortener/internal/models"
)

type URLRepository interface {
	SaveURL(url *models.URL) error
	FindByShortCode(code string) (*models.URL, error)
}
