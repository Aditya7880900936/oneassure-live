package handler

import (
	"errors"
	"net/http"

	"one-assure/livedemo/internal/model"
	"one-assure/livedemo/internal/services"

	"github.com/gin-gonic/gin"
)

type MarkdownHandler struct {
	service *services.MarkdownService
}

func NewMarkdownHandler(service *services.MarkdownService) *MarkdownHandler {
	return &MarkdownHandler{
		service: service,
	}
}

// RenderMarkdown converts Markdown supplied through the GET query string.
// Example: GET /markdown?content=%23%20Hello
func (h *MarkdownHandler) RenderMarkdown(c *gin.Context) {
	content, exists := c.GetQuery("content")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "content query parameter is required",
		})
		return
	}

	html, err := h.service.ConvertMarkdown(content)
	if err != nil {
		writeConversionError(c, err)
		return
	}

	c.JSON(http.StatusOK, model.ConversionResponse{
		HTML: html,
	})
}

// ConvertMarkdown converts Markdown supplied in a JSON request body.
func (h *MarkdownHandler) ConvertMarkdown(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(
		c.Writer,
		c.Request.Body,
		services.MaxMarkdownSize+4096,
	)

	var request model.Markdown

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request: expected JSON with a non-empty content field",
		})
		return
	}

	html, err := h.service.ConvertMarkdown(request.Content)
	if err != nil {
		writeConversionError(c, err)
		return
	}

	c.JSON(http.StatusOK, model.ConversionResponse{
		HTML: html,
	})
}

func writeConversionError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, services.ErrEmptyMarkdown):
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	case errors.Is(err, services.ErrMarkdownTooLarge):
		c.JSON(http.StatusRequestEntityTooLarge, gin.H{
			"error": err.Error(),
		})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to convert Markdown",
		})
	}
}
