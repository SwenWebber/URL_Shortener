package service

import (
	"testing"
	"urlShortener/internal/models"
)

type MockURLRepository struct {
	urls map[string]*models.URL
}

func NewMockRepository() *MockURLRepository {
	return &MockURLRepository{
		urls: make(map[string]*models.URL),
	}
}

func (m *MockURLRepository) SaveURL(url *models.URL) error {
	m.urls[url.ShortCode] = url
	return nil
}

func (m *MockURLRepository) FindByLongUrl(code string) (*models.URL, error) {
	return nil, nil
}

func (m *MockURLRepository) FindByShortCode(code string) (*models.URL, error) {
	if url, exists := m.urls[code]; exists {
		return url, nil
	}
	return nil, nil
}

type MockURLGenerator struct {
	nextCode string
}

func NewMockGenerator() *MockURLGenerator {
	return &MockURLGenerator{
		nextCode: "abc123", // фиксированный код для тестов
	}
}

func (g *MockURLGenerator) Generate() string {
	return g.nextCode
}

func TestCreateShortUrl(t *testing.T) {
	mockRepo := NewMockRepository()
	mockGenerator := NewMockGenerator()
	urlService := NewUrlService(mockRepo, mockGenerator)

	tests := []struct {
		name    string
		longURL string
		wantErr bool
	}{
		{
			name:    "Valid URL",
			longURL: "https://google.com",
			wantErr: false,
		},
		{
			name:    "Empty URL",
			longURL: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := urlService.CreateShortURL(tt.longURL)

			if (err != nil) != tt.wantErr {
				t.Errorf("CreateShortUrl() error =%v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if result == nil {
					t.Error("Expected URL object, got nil")
					return
				}

				if result.OriginalURL != tt.longURL {
					t.Errorf("Expected long URL %s, got %s instead", tt.longURL, result.OriginalURL)
				}

				if result.ShortCode == "" {
					t.Error("Expected short code, got empty string")
				}
			}
		})
	}

}
