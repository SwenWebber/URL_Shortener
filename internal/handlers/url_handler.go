package handlers

import (
	"net/http"
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
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if req.URL == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "URL is required"})
	}

	url, err := h.service.CreateShortURL(req.URL)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error:": "Failed to create short URL"})
	}

	shortURL := h.baseURL + "/" + url.ShortCode
	return c.JSON(http.StatusCreated, models.CreateURLResponse{ShortURL: shortURL})
}

func (h *URlHandler) RedirectToURL(c echo.Context) error {
	code := c.Param("code")

	url, err := h.service.GetLongURL(code)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "URL not found"})
	}

	return c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
}
