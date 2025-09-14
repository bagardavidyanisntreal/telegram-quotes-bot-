package usecases

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"telegram-quotes-bot/internal/entities"
	"telegram-quotes-bot/internal/interfaces"
)

// SendQuoteService предоставляет методы для отправки цитат в Telegram-канал.
type SendQuoteService struct {
	telegram interfaces.TelegramSender // Интерфейс для отправки сообщений в Telegram
}

// NewSendQuoteService создаёт новый экземпляр SendQuoteService.
// Принимает интерфейс TelegramSender для отправки сообщений в Telegram.
func NewSendQuoteService(telegram interfaces.TelegramSender) *SendQuoteService {
	return &SendQuoteService{telegram: telegram}
}

// SendQuote отправляет цитату в Telegram-канал.
// Форматирует цитату в удобочитаемый вид и отправляет её через TelegramSender.
// Возвращает ошибку, если отправка не удалась.
func (s *SendQuoteService) SendQuote(ctx context.Context, quote *entities.Quote) error {
	// Проверяем, что цитата не nil
	if quote == nil {
		return fmt.Errorf("цитата не может быть nil")
	}

	// Форматируем цитату с красивым оформлением
	message := s.FormatQuote(quote)

	// Отправляем сформированное сообщение через TelegramSender
	err := s.telegram.SendMessage(ctx, message)
	if err != nil {
		// Если произошла ошибка при отправке, возвращаем её с описанием
		return fmt.Errorf("не удалось отправить сообщение: %w", err)
	}

	// Если всё прошло успешно, возвращаем nil
	return nil
}

// FormatQuote создает красиво отформатированное сообщение с цитатой (публичная функция для тестирования)
func (s *SendQuoteService) FormatQuote(quote *entities.Quote) string {
	rand.Seed(time.Now().UnixNano())

	// Ограничиваем длину цитаты для лучшего отображения
	text := quote.Text
	if len(text) > 200 {
		text = text[:197] + "..."
	}

	// Выбираем случайный стиль форматирования
	styles := []func(string, string) string{
		s.formatStyle1,
		s.formatStyle2,
		s.formatStyle3,
		s.formatStyle4,
	}

	style := styles[rand.Intn(len(styles))]
	return style(text, quote.Author)
}

// formatStyle1 - Стиль с рамкой
func (s *SendQuoteService) formatStyle1(text, author string) string {
	quoteEmojis := []string{"💭", "✨", "🌟", "💫", "🎯", "🔥", "💡", "🌈", "🦋", "🌸"}
	authorEmojis := []string{"✍️", "👤", "🎭", "📝", "🖋️", "✏️", "📖", "📚", "🎨", "🎪"}

	quoteEmoji := quoteEmojis[rand.Intn(len(quoteEmojis))]
	authorEmoji := authorEmojis[rand.Intn(len(authorEmojis))]

	return fmt.Sprintf(
		"%s *Цитата дня*\n\n"+
			"┌─────────────────────────┐\n"+
			"│  %s  │\n"+
			"│                         │\n"+
			"│  %s  │\n"+
			"└─────────────────────────┘\n\n"+
			"%s *%s*",
		quoteEmoji,
		text,
		strings.Repeat("─", 25),
		authorEmoji,
		author,
	)
}

// formatStyle2 - Стиль с кавычками
func (s *SendQuoteService) formatStyle2(text, author string) string {
	emojis := []string{"💫", "✨", "🌟", "🎯", "🔥", "💡", "🌈", "🦋", "🌸", "🎪"}
	emoji := emojis[rand.Intn(len(emojis))]

	return fmt.Sprintf(
		"%s *Мудрая мысль*\n\n"+
			"❝ %s ❞\n\n"+
			"    — *%s* ✍️",
		emoji,
		text,
		author,
	)
}

// formatStyle3 - Стиль с разделителями
func (s *SendQuoteService) formatStyle3(text, author string) string {
	emojis := []string{"🌟", "💫", "✨", "🎯", "🔥", "💡", "🌈", "🦋", "🌸", "🎨"}
	emoji := emojis[rand.Intn(len(emojis))]

	return fmt.Sprintf(
		"%s *Вдохновение дня*\n\n"+
			"━━━━━━━━━━━━━━━━━━━━━━━━━━\n"+
			"  %s\n"+
			"━━━━━━━━━━━━━━━━━━━━━━━━━━\n\n"+
			"👤 *%s*",
		emoji,
		text,
		author,
	)
}

// formatStyle4 - Стиль с эмодзи-рамкой
func (s *SendQuoteService) formatStyle4(text, author string) string {
	emojis := []string{"💭", "✨", "🌟", "💫", "🎯", "🔥", "💡", "🌈", "🦋", "🌸"}
	emoji := emojis[rand.Intn(len(emojis))]

	return fmt.Sprintf(
		"%s *Цитата дня*\n\n"+
			"🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦\n"+
			"🟦  %s  🟦\n"+
			"🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦\n\n"+
			"✍️ *%s*",
		emoji,
		text,
		author,
	)
}
