package validators

import (
	"errors"
	"strings"
)

// ValidateQuote проверяет валидность цитаты
func ValidateQuote(text, author string) error {
	if strings.TrimSpace(text) == "" {
		return errors.New("текст цитаты не может быть пустым")
	}

	if strings.TrimSpace(author) == "" {
		return errors.New("автор цитаты не может быть пустым")
	}

	// Проверяем длину текста
	if len(text) > 1000 {
		return errors.New("текст цитаты слишком длинный (максимум 1000 символов)")
	}

	// Проверяем длину автора
	if len(author) > 100 {
		return errors.New("имя автора слишком длинное (максимум 100 символов)")
	}

	// Проверяем на потенциально опасные символы
	if containsDangerousChars(text) {
		return errors.New("текст цитаты содержит недопустимые символы")
	}

	if containsDangerousChars(author) {
		return errors.New("имя автора содержит недопустимые символы")
	}

	return nil
}

// ValidateBotToken проверяет валидность токена бота
func ValidateBotToken(token string) error {
	if strings.TrimSpace(token) == "" {
		return errors.New("токен бота не может быть пустым")
	}

	// Проверяем формат токена Telegram бота (должен содержать двоеточие)
	if !strings.Contains(token, ":") {
		return errors.New("неверный формат токена бота")
	}

	// Проверяем длину токена
	if len(token) < 20 || len(token) > 100 {
		return errors.New("неверная длина токена бота")
	}

	return nil
}

// ValidateChatID проверяет валидность ID чата
func ValidateChatID(chatID int64) error {
	if chatID == 0 {
		return errors.New("ID чата не может быть равен нулю")
	}

	// Проверяем, что ID чата в разумных пределах для Telegram
	// Telegram ID чатов могут быть от -2^63 до 2^63-1, но ограничиваем разумными пределами
	if chatID < -999999999999999 || chatID > 999999999999999 {
		return errors.New("ID чата вне допустимого диапазона")
	}

	return nil
}

// containsDangerousChars проверяет наличие потенциально опасных символов
func containsDangerousChars(text string) bool {
	// Список потенциально опасных символов
	dangerousChars := []string{
		"<script", "</script>", "javascript:", "data:",
		"vbscript:", "onload=", "onerror=", "onclick=",
	}

	textLower := strings.ToLower(text)
	for _, char := range dangerousChars {
		if strings.Contains(textLower, char) {
			return true
		}
	}

	return false
}
