package handlers

import (
	"urlShortener/internal/models"
	"urlShortener/internal/service"

	"github.com/labstack/echo/v4"
)

type URlHandler struct {
	service service.URLService
	baseURL string
}

func NewURLHandler(service service.URLService, baseURL string) *URlHandler {
	return &URlHandler{
		service: service,
		baseURL: baseURL,
	}
}

func (h *URlHandler) CreateShortURL(c echo.Context) error {
	var req models.CreateURLRequest

}
