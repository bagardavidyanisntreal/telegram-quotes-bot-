package usecases

import (
	"context"
	"errors"
	"testing"

	"telegram-quotes-bot/internal/entities"
)

// mockTelegramSender - мок для тестирования
type mockTelegramSender struct {
	err error
}

func (m *mockTelegramSender) SendMessage(ctx context.Context, message string) error {
	return m.err
}

func TestSendQuoteService_SendQuote(t *testing.T) {
	tests := []struct {
		name        string
		mockSender  *mockTelegramSender
		quote       *entities.Quote
		expectedErr bool
	}{
		{
			name: "Success",
			mockSender: &mockTelegramSender{
				err: nil,
			},
			quote: &entities.Quote{
				Text:   "Test quote",
				Author: "Test Author",
			},
			expectedErr: false,
		},
		{
			name: "Send Error",
			mockSender: &mockTelegramSender{
				err: errors.New("send error"),
			},
			quote: &entities.Quote{
				Text:   "Test quote",
				Author: "Test Author",
			},
			expectedErr: true,
		},
		{
			name: "Nil Quote",
			mockSender: &mockTelegramSender{
				err: nil,
			},
			quote:       nil,
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewSendQuoteService(tt.mockSender)
			ctx := context.Background()

			err := service.SendQuote(ctx, tt.quote)

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
