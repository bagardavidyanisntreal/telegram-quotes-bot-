package validators

import (
	"testing"
)

func TestValidateQuote(t *testing.T) {
	tests := []struct {
		name        string
		text        string
		author      string
		expectedErr bool
	}{
		{
			name:        "Valid quote",
			text:        "The only way to do great work is to love what you do.",
			author:      "Steve Jobs",
			expectedErr: false,
		},
		{
			name:        "Empty text",
			text:        "",
			author:      "Steve Jobs",
			expectedErr: true,
		},
		{
			name:        "Empty author",
			text:        "The only way to do great work is to love what you do.",
			author:      "",
			expectedErr: true,
		},
		{
			name:        "Whitespace only text",
			text:        "   ",
			author:      "Steve Jobs",
			expectedErr: true,
		},
		{
			name:        "Whitespace only author",
			text:        "The only way to do great work is to love what you do.",
			author:      "   ",
			expectedErr: true,
		},
		{
			name:        "Text too long",
			text:        string(make([]byte, 1001)),
			author:      "Steve Jobs",
			expectedErr: true,
		},
		{
			name:        "Author too long",
			text:        "The only way to do great work is to love what you do.",
			author:      string(make([]byte, 101)),
			expectedErr: true,
		},
		{
			name:        "Text with dangerous chars",
			text:        "Hello <script>alert('xss')</script>",
			author:      "Steve Jobs",
			expectedErr: true,
		},
		{
			name:        "Author with dangerous chars",
			text:        "The only way to do great work is to love what you do.",
			author:      "Steve <script>alert('xss')</script>",
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateQuote(tt.text, tt.author)
			if tt.expectedErr {
				if err == nil {
					t.Error("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
			}
		})
	}
}

func TestValidateBotToken(t *testing.T) {
	tests := []struct {
		name        string
		token       string
		expectedErr bool
	}{
		{
			name:        "Valid token",
			token:       "123456789:ABCdefGHIjklMNOpqrsTUVwxyz",
			expectedErr: false,
		},
		{
			name:        "Empty token",
			token:       "",
			expectedErr: true,
		},
		{
			name:        "Token without colon",
			token:       "123456789ABCdefGHIjklMNOpqrsTUVwxyz",
			expectedErr: true,
		},
		{
			name:        "Token too short",
			token:       "123:abc",
			expectedErr: true,
		},
		{
			name:        "Token too long",
			token:       string(make([]byte, 101)),
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateBotToken(tt.token)
			if tt.expectedErr {
				if err == nil {
					t.Error("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
			}
		})
	}
}

func TestValidateChatID(t *testing.T) {
	tests := []struct {
		name        string
		chatID      int64
		expectedErr bool
	}{
		{
			name:        "Valid positive chat ID",
			chatID:      123456789,
			expectedErr: false,
		},
		{
			name:        "Valid negative chat ID",
			chatID:      -123456789,
			expectedErr: false,
		},
		{
			name:        "Zero chat ID",
			chatID:      0,
			expectedErr: true,
		},
		{
			name:        "Chat ID too large",
			chatID:      1000000000000000,
			expectedErr: true,
		},
		{
			name:        "Chat ID too small",
			chatID:      -1000000000000000,
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateChatID(tt.chatID)
			if tt.expectedErr {
				if err == nil {
					t.Error("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
			}
		})
	}
}
