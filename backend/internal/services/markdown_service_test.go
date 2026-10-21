package services

import (
	"strings"
	"testing"
)

func TestConvertMarkdown(t *testing.T) {
	service := NewMarkdownService()

	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "heading",
			input: "# Hello",
			want:  []string{"<h1", "Hello", "</h1>"},
		},
		{
			name:  "bold text",
			input: "**Hello**",
			want:  []string{"<strong>Hello</strong>"},
		},
		{
			name:  "unordered list",
			input: "- one\n- two",
			want:  []string{"<ul>", "<li>one</li>", "<li>two</li>"},
		},
		{
			name:  "GFM strikethrough",
			input: "~~removed~~",
			want:  []string{"<del>removed</del>"},
		},
		{
			name:  "raw HTML is not rendered as HTML",
			input: "<script>alert(1)</script>",
			want:  []string{"raw HTML omitted"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.ConvertMarkdown(tt.input)
			if err != nil {
				t.Fatalf("ConvertMarkdown() error = %v", err)
			}

			for _, expected := range tt.want {
				if !strings.Contains(got, expected) {
					t.Errorf("output %q does not contain %q", got, expected)
				}
			}
		})
	}
}

func TestConvertMarkdownRejectsEmptyInput(t *testing.T) {
	service := NewMarkdownService()

	_, err := service.ConvertMarkdown("   ")
	if err != ErrEmptyMarkdown {
		t.Fatalf("expected ErrEmptyMarkdown, got %v", err)
	}
}

func TestConvertMarkdownRejectsOversizedInput(t *testing.T) {
	service := NewMarkdownService()
	input := strings.Repeat("a", MaxMarkdownSize+1)

	_, err := service.ConvertMarkdown(input)
	if err != ErrMarkdownTooLarge {
		t.Fatalf("expected ErrMarkdownTooLarge, got %v", err)
	}
}

func TestConvertMarkdownDoesNotRenderRawHTML(t *testing.T) {
	service := NewMarkdownService()

	got, err := service.ConvertMarkdown("<script>alert(1)</script>")
	if err != nil {
		t.Fatalf("ConvertMarkdown() error = %v", err)
	}

	if strings.Contains(got, "<script>") {
		t.Fatalf("raw script tag must not appear in output: %q", got)
	}

	if !strings.Contains(got, "raw HTML omitted") {
		t.Fatalf("expected Goldmark to omit raw HTML, got %q", got)
	}
}
