package service

import (
	"time"
	"urlShortener/internal/models"
	"urlShortener/internal/repository"
)

type urlService struct {
	repo      repository.URLRepository
	generator URLGenerator
}

func (s *urlService) CreateShortURL(longURL string) (*models.URL, error) {
	url := &models.URL{
		OriginalURL: longURL,
		ShortCode:   s.generator.Generate(),
		CreatedAt:   time.Now(),
	}

	if err := s.repo.SaveURL(url); err != nil {
		return nil, err
	}

	return url, nil
}

func (s *urlService) GetLongURL(shortCode string) (*models.URL, error) {
	return s.repo.FindByShortCode(shortCode)
}
