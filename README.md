# Telegram Quotes Bot

##### Telegram bot that automatically sends motivational quotes in Russian to the channel on a scheduled basis. The bot receives random quotes from Forismatic API and sends them to Telegram channel via Telegram Bot API.

## Features
- ✅ Receiving random motivational quotes in Russian from Forismatic API
- ✅ Sending quotes to Telegram channel on schedule (Cron)
- ✅ Test quote sent on startup (configurable)
- ✅ Graceful shutdown with signal handling
- ✅ HTTP clients with timeouts for reliability
- ✅ Input validation and security checks
- ✅ Comprehensive unit tests
- ✅ Safe logging (sensitive data masking)
- ✅ Clean Architecture pattern

## Technologies
- **Programming language**: Go 1.23
- **APIs**: Telegram Bot API, Forismatic API (Russian quotes)
- **Task Scheduler**: Cron (robfig/cron/v3)
- **Architecture**: Clean Architecture with dependency injection
- **Testing**: Go testing framework with mocks
- **Containerization**: Docker
- **Task Runner**: Taskfile

## Project Structure
```
├── cmd/                    # Application entry point
├── internal/
│   ├── adapters/          # External service adapters
│   ├── config/            # Configuration management
│   ├── entities/          # Domain models
│   ├── interfaces/        # Service contracts
│   ├── usecases/          # Business logic
│   └── validators/        # Input validation
├── Dockerfile             # Container configuration
├── Taskfile.yml          # Task automation
└── env.example           # Environment variables template
```

## Quick Start

### Prerequisites
- Go 1.23+
- Docker (optional)
- Task (optional, for task automation)

### Setup
1. Clone the repository
2. Copy `env.example` to `.env` and configure:
   ```bash
   cp env.example .env
   ```
3. Set your Telegram bot token and chat ID in `.env`

### Running
```bash
# Using Go directly
go run cmd/main.go

# Using Task
task run

# Using Docker
docker build -t telegram-quotes-bot .
docker run --env-file .env telegram-quotes-bot
```

### Development
```bash
# Run tests
task test

# Run tests with coverage
task test-coverage

# Build application
task build

# Lint code
task lint

# Clean artifacts
task clean
```

## Configuration
The bot requires the following environment variables:
- `BOT_TOKEN`: Your Telegram bot token (from @BotFather)
- `CHAT_ID`: Target chat/channel ID for sending quotes
- `SEND_TEST_QUOTE`: Send test quote on startup (true/false, default: true)

## Schedule
The bot sends quotes at: 4:00, 8:00, 14:00, and 18:00 daily (UTC).

## API Information
The bot uses [Forismatic API](http://api.forismatic.com/) to get random quotes in Russian:
- **Endpoint**: `http://api.forismatic.com/api/1.0/?method=getQuote&format=json&lang=ru`
- **Response format**: JSON with `quoteText` and `quoteAuthor` fields
- **Language**: Russian quotes directly from the API (no translation needed)

## Security Features
- Input validation for all external data
- XSS protection in quote content
- Safe logging with token masking
- HTTP timeouts to prevent hanging requests
- Graceful shutdown handling
