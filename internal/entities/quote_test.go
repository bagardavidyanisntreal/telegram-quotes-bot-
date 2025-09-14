package entities

import (
	"testing"
)

func TestQuote(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		author   string
		expected string
	}{
		{
			name:     "Valid quote",
			text:     "The only way to do great work is to love what you do.",
			author:   "Steve Jobs",
			expected: "The only way to do great work is to love what you do.",
		},
		{
			name:     "Empty quote",
			text:     "",
			author:   "Unknown",
			expected: "",
		},
		{
			name:     "Quote with special characters",
			text:     "Don't be afraid to give up the good to go for the great.",
			author:   "John D. Rockefeller",
			expected: "Don't be afraid to give up the good to go for the great.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quote := &Quote{
				Text:   tt.text,
				Author: tt.author,
			}

			if quote.Text != tt.expected {
				t.Errorf("Expected text %q, got %q", tt.expected, quote.Text)
			}

			if quote.Author != tt.author {
				t.Errorf("Expected author %q, got %q", tt.author, quote.Author)
			}
		})
	}
}
