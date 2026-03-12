# 📧 Mail Service

Лёгкий HTTP-сервис на Go для отправки HTML-писем через SMTP. Разработан как отдельный микросервис.

---

## Возможности

- Отправка HTML-писем через SMTP
- Структурированное логирование через `slog` (текст в dev, JSON в prod)
- Авторотация лог-файлов через `lumberjack`
- Готов к запуску через Docker / docker-compose

---

## API

### `POST /send`

Отправляет письмо на указанный адрес.

**Тело запроса:**

```json
{
  "to": "user@example.com",
  "message": "<h1>Привет!</h1><p>Это тестовое письмо.</p>",
  "subject": "Тема письма"
}
```

| Поле      | Тип    | Обязательное | Описание                                         |
| --------- | ------ | :----------: | ------------------------------------------------ |
| `to`      | string |      ✅      | Email получателя                                 |
| `message` | string |      ✅      | Тело письма в формате HTML                       |
| `subject` | string |      ❌      | Тема письма. Дефолт: `Mail service notification` |

---

## Переменные окружения

Создайте файл `.env` на основе `.env.example`:

```env
# Сервер
PORT=4005
IS_PROD=false

# SMTP
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_FROM=your@gmail.com
SMTP_PASSWORD=your_app_password

# Логи
LOG_FILE=logs/mail-service.log
```

---

## Запуск

### Локально

```bash
# Установить зависимости
go mod download

# Запустить
go run cmd/main.go
```

### Сборка бинарника

```bash
go build -o mail-service ./cmd/main.go
```

### Docker

```bash
docker build -t mail-service .
docker run --env-file .env -p 4005:4005 mail-service
```
