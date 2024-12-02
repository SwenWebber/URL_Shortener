package repository

import (
	"context"
	"urlShortener/internal/models"
)

type URLRepository interface {
	Create(ctx context.Context, url *models.URL) error
	GetByShortCode(ctx context.Context, shortCode string) (*models.URL, error)
	GetByOriginalURL(ctx context.Context, originalURL string) (*models.URL, error)
}
