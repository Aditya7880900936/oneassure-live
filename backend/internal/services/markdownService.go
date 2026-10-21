package services

import (
	"bytes"
	"errors"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

const MaxMarkdownSize = 1 << 20 // 1 MiB

var ErrEmptyMarkdown = errors.New("markdown content cannot be empty")

var ErrMarkdownTooLarge = errors.New("markdown content exceeds 1 MiB limit")

type MarkdownService struct {
	converter goldmark.Markdown
}

func NewMarkdownService() *MarkdownService {
	converter := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
	)

	return &MarkdownService{
		converter: converter,
	}
}

func (s *MarkdownService) ConvertMarkdown(content string) (string, error) {
	if strings.TrimSpace(content) == "" {
		return "", ErrEmptyMarkdown
	}

	if len(content) > MaxMarkdownSize {
		return "", ErrMarkdownTooLarge
	}

	var buf bytes.Buffer

	if err := s.converter.Convert([]byte(content), &buf); err != nil {
		return "", err
	}

	return buf.String(), nil
}
