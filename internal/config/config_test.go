package config

import (
	"log/slog"
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	tests := []struct {
		name              string
		botToken          string
		chatID            string
		sendTestQuote     string
		expectedErr       bool
		expectedTestQuote bool
	}{
		{
			name:              "Valid config with default test quote",
			botToken:          "123456789:ABCdefGHIjklMNOpqrsTUVwxyz",
			chatID:            "123456789",
			sendTestQuote:     "",
			expectedErr:       false,
			expectedTestQuote: true,
		},
		{
			name:              "Valid config with test quote enabled",
			botToken:          "123456789:ABCdefGHIjklMNOpqrsTUVwxyz",
			chatID:            "123456789",
			sendTestQuote:     "true",
			expectedErr:       false,
			expectedTestQuote: true,
		},
		{
			name:              "Valid config with test quote disabled",
			botToken:          "123456789:ABCdefGHIjklMNOpqrsTUVwxyz",
			chatID:            "123456789",
			sendTestQuote:     "false",
			expectedErr:       false,
			expectedTestQuote: false,
		},
		{
			name:          "Empty bot token",
			botToken:      "",
			chatID:        "123456789",
			sendTestQuote: "",
			expectedErr:   true,
		},
		{
			name:          "Empty chat ID",
			botToken:      "123456789:ABCdefGHIjklMNOpqrsTUVwxyz",
			chatID:        "",
			sendTestQuote: "",
			expectedErr:   true,
		},
		{
			name:          "Invalid chat ID",
			botToken:      "123456789:ABCdefGHIjklMNOpqrsTUVwxyz",
			chatID:        "invalid",
			sendTestQuote: "",
			expectedErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Сохраняем оригинальные значения
			originalBotToken := os.Getenv("BOT_TOKEN")
			originalChatID := os.Getenv("CHAT_ID")
			originalSendTestQuote := os.Getenv("SEND_TEST_QUOTE")

			// Устанавливаем тестовые значения
			os.Setenv("BOT_TOKEN", tt.botToken)
			os.Setenv("CHAT_ID", tt.chatID)
			os.Setenv("SEND_TEST_QUOTE", tt.sendTestQuote)

			// Восстанавливаем оригинальные значения после теста
			defer func() {
				os.Setenv("BOT_TOKEN", originalBotToken)
				os.Setenv("CHAT_ID", originalChatID)
				os.Setenv("SEND_TEST_QUOTE", originalSendTestQuote)
			}()

			config, err := LoadConfig(logger)

			if tt.expectedErr {
				if err == nil {
					t.Error("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
				if config == nil {
					t.Error("Expected config, got nil")
				} else {
					if config.BotToken != tt.botToken {
						t.Errorf("Expected bot token %q, got %q", tt.botToken, config.BotToken)
					}
					if config.ChatID != 123456789 {
						t.Errorf("Expected chat ID %d, got %d", 123456789, config.ChatID)
					}
					if config.SendTestQuote != tt.expectedTestQuote {
						t.Errorf("Expected SendTestQuote %v, got %v", tt.expectedTestQuote, config.SendTestQuote)
					}
				}
			}
		})
	}
}

func TestMaskToken(t *testing.T) {
	tests := []struct {
		name     string
		token    string
		expected string
	}{
		{
			name:     "Empty token",
			token:    "",
			expected: "***",
		},
		{
			name:     "Short token",
			token:    "12345678",
			expected: "***",
		},
		{
			name:     "Long token",
			token:    "123456789:ABCdefGHIjklMNOpqrsTUVwxyz",
			expected: "1234***wxyz",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maskToken(tt.token)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}
