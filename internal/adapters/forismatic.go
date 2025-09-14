package adapters

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"telegram-quotes-bot/internal/entities"
)

// ForismaticAPI реализует интерфейс QuoteAPI для получения случайных цитат на русском языке из Forismatic API.
type ForismaticAPI struct {
	client *http.Client
}

// NewForismaticAPI создаёт новый экземпляр ForismaticAPI.
func NewForismaticAPI() *ForismaticAPI {
	return &ForismaticAPI{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// GetRandomQuote получает случайную цитату на русском языке из Forismatic API.
// Возвращает структуру Quote или ошибку, если запрос или декодирование не удались.
func (f *ForismaticAPI) GetRandomQuote(ctx context.Context) (*entities.Quote, error) {
	// Создаем запрос с контекстом
	req, err := http.NewRequestWithContext(ctx, "GET", "http://api.forismatic.com/api/1.0/?method=getQuote&format=json&lang=ru", nil)
	if err != nil {
		return nil, errors.New("ошибка создания запроса")
	}

	// Выполняем GET-запрос к Forismatic API для получения случайной цитаты
	resp, err := f.client.Do(req)
	if err != nil {
		// Если произошла ошибка при выполнении запроса, возвращаем её
		return nil, errors.New("ошибка запроса к API")
	}
	defer resp.Body.Close() // Закрываем тело ответа после завершения работы

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("API вернул неожиданный статус")
	}

	// Определяем структуру для декодирования JSON-ответа
	var quoteResponse struct {
		QuoteText   string `json:"quoteText"`   // Текст цитаты
		QuoteAuthor string `json:"quoteAuthor"` // Имя автора
		SenderName  string `json:"senderName"`  // Имя отправителя (обычно пустое)
		SenderLink  string `json:"senderLink"`  // Ссылка отправителя (обычно пустая)
	}

	// Декодируем JSON-ответ от API
	if err := json.NewDecoder(resp.Body).Decode(&quoteResponse); err != nil {
		// Если произошла ошибка при декодировании JSON, возвращаем её
		return nil, errors.New("ошибка декодирования JSON")
	}

	// Проверяем, что ответ содержит текст цитаты
	if quoteResponse.QuoteText == "" {
		return nil, errors.New("получена пустая цитата")
	}

	// Если автор не указан, используем "Неизвестный автор"
	author := quoteResponse.QuoteAuthor
	if author == "" {
		author = "Неизвестный автор"
	}

	// Возвращаем цитату, преобразованную в структуру Quote
	return &entities.Quote{
		Text:   quoteResponse.QuoteText, // Текст цитаты
		Author: author,                  // Имя автора
	}, nil
}
