package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"telegram-quotes-bot/internal/adapters"
	"telegram-quotes-bot/internal/config"
	"telegram-quotes-bot/internal/usecases"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

// setupLogger логгер
func setupLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	return logger
}

func main() {
	// Настройка логгера
	logger := setupLogger()

	// Загрузка .env файла
	if err := godotenv.Load(); err != nil {
		logger.Warn("Файл .env не найден или не загружен")
	}

	// Создаем контекст с обработкой сигналов для graceful shutdown
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// Загрузка конфигурации
	cfg, err := config.LoadConfig(logger)
	if err != nil {
		logger.Error("Ошибка загрузки конфигурации", "error", err)
		os.Exit(1)
	}

	// Инициализация адаптеров
	quoteAPI := adapters.NewForismaticAPI()
	telegramAdapter, err := adapters.NewTelegramAdapter(cfg.BotToken, cfg.ChatID)
	if err != nil {
		logger.Error("Не удалось инициализировать TelegramAdapter", "error", err)
		os.Exit(1)
	}

	// Инициализация сервисов
	fetchQuoteService := usecases.NewFetchQuoteService(quoteAPI)
	sendQuoteService := usecases.NewSendQuoteService(telegramAdapter)

	// Планировщик Cron
	c := cron.New()
	defer c.Stop()

	// Задача отправки цитат
	_, err = c.AddFunc("0 4,8,14,18 * * *", func() {
		taskCtx := context.Background()

		// Получение цитаты на русском языке
		quote, err := fetchQuoteService.FetchQuote(taskCtx)
		if err != nil {
			logger.Error("Ошибка получения цитаты", "error", err)
			return
		}

		// Отправка цитаты
		if err := sendQuoteService.SendQuote(taskCtx, quote); err != nil {
			logger.Error("Ошибка отправки цитаты", "error", err)
		} else {
			logger.Info("Цитата успешно отправлена", "quote", quote.Text, "author", quote.Author)
		}
	})
	if err != nil {
		logger.Error("Не удалось добавить cron-задачу", "error", err)
		os.Exit(1)
	}

	// Запуск планировщика
	c.Start()
	logger.Info("Планировщик запущен. Ожидание задач.")

	// Отправка тестовой цитаты при запуске (если включено в конфигурации)
	if cfg.SendTestQuote {
		logger.Info("Отправка тестовой цитаты...")
		testCtx := context.Background()

		// Получение тестовой цитаты
		testQuote, err := fetchQuoteService.FetchQuote(testCtx)
		if err != nil {
			logger.Error("Ошибка получения тестовой цитаты", "error", err)
		} else {
			// Отправка тестовой цитаты
			if err := sendQuoteService.SendQuote(testCtx, testQuote); err != nil {
				logger.Error("Ошибка отправки тестовой цитаты", "error", err)
			} else {
				logger.Info("Тестовая цитата успешно отправлена", "quote", testQuote.Text, "author", testQuote.Author)
			}
		}
	} else {
		logger.Info("Отправка тестовой цитаты отключена в конфигурации")
	}

	// Ожидание сигнала завершения
	<-ctx.Done()
	logger.Info("Получен сигнал завершения. Останавливаем планировщик...")

	// Остановка планировщика
	stopCtx := c.Stop()
	<-stopCtx.Done()
	logger.Info("Планировщик остановлен. Программа завершена.")
}
