package service

import (
	"log"
	"time"
	"urlShortener/internal/models"
	"urlShortener/internal/repository"
)

type urlService struct {
	repo      repository.URLRepository
	generator URLGenerator
}

func NewUrlService(repo repository.URLRepository, generator URLGenerator) URLService {
	return &urlService{
		repo:      repo,
		generator: generator,
	}
}

func (s *urlService) CreateShortURL(longURL string) (*models.URL, error) {

	existing, _ := s.repo.FindByLongUrl(longURL)

	if existing != nil {
		return existing, nil
	}

	newUrl := &models.URL{
		OriginalURL: longURL,
		ShortCode:   s.generator.Generate(),
		CreatedAt:   time.Now(),
	}

	if err := s.repo.SaveURL(newUrl); err != nil {
		log.Printf("Error saving URL:%v", err)
		return nil, err
	}

	return newUrl, nil
}

func (s *urlService) GetLongURL(shortCode string) (*models.URL, error) {
	return s.repo.FindByShortCode(shortCode)
}
